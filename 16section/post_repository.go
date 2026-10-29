package main

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"net/url"
	"strings"
	"time"

	"github.com/dromara/carbon/v2"
)

var (
	ErrDuplicatePostTitle = errors.New("post with specified title already exists")
	ErrDuplicateVote      = errors.New("user has already voted for this post")
	ErrPostNotFound       = errors.New("post not found")
)

type Post struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	URL          string    `json:"url"`
	UserID       int       `json:"user_id"`
	UserName     string    `json:"user_name"`
	CreatedAt    time.Time `json:"created_at"`
	CommentCount int       `json:"comment_count"`
	VoteCount    int       `json:"vote_count"`
	TotalRecords int       `json:"total_records"`
}

type Comment struct {
	ID        int       `json:"id"`
	Comment   string    `json:"comment"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

type Filter struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	OrderBy  string `json:"order_by"`
	Query    string `json:"query"`
}

func (f *Filter) Validate() error {
	if f.PageSize <= 0 || f.PageSize >= 1000 {
		return errors.New("Invalid 'Page' range: 1 to 1000")
	}
	return nil
}

type Metadata struct {
	CurrentPage  int `json:"current_page"`
	PageSize     int `json:"page_size"`
	FirstPage    int `json:"first_page"`
	NextPage     int `json:"next_page"`
	PrevPage     int `json:"prev_page"`
	LastPage     int `json:"last_page"`
	TotalRecords int `json:"total_records"`
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	lastPage := int(math.Ceil(float64(totalRecords) / float64(pageSize)))
	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     lastPage,
		NextPage:     int(math.Min(float64(page+1), float64(lastPage))),
		PrevPage:     int(math.Max(float64(page-1), 1)),
		TotalRecords: totalRecords,
	}
}

type PostRepository interface {
	CreatePost(title, url string, userID int) (int, error)
	AddComment(comment string, userID, postID int) (int, error)
	AddVote(userID, postID int) error
	GetAll(filter Filter) ([]Post, Metadata, error)
	GetById(id int) (*Post, error)
	GetComments(postId int) ([]Comment, error)
}

type SqlPostRepository struct {
	db *sql.DB
}

// NewSqlPostRepository creates a new PostRepository based on SqlPostRepository struct
func NewSqlPostRepository(db *sql.DB) PostRepository {
	return &SqlPostRepository{db: db}
}

func (r *SqlPostRepository) CreatePost(title, url string, userID int) (int, error) {
	stmt := "INSERT INTO posts (title, url, user_id) VALUES (?, ?, ?)"

	result, err := r.db.Exec(stmt, title, url, userID)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: post.title") {
			return 0, ErrDuplicatePostTitle
		}
		return 0, err
	}

	postId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(postId), nil
}

func (r *SqlPostRepository) AddComment(comment string, userID, postID int) (int, error) {
	stmt := "INSERT INTO comments (comment, user_id, post_id) VALUES (?, ?, ?)"

	result, err := r.db.Exec(stmt, comment, userID, postID)
	if err != nil {
		return 0, err
	}

	commentID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(commentID), nil
}

func (r *SqlPostRepository) AddVote(userID, postID int) error {
	stmt, err := r.db.Prepare("INSERT INTO votes (user_id, post_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, postID)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: votes.user_id, votes.post_id") {
			return ErrDuplicateVote

		}
		return err
	}

	return nil
}

func (r *SqlPostRepository) GetAll(filter Filter) ([]Post, Metadata, error) {
	if err := filter.Validate(); err != nil {
		return nil, Metadata{}, err
	}

	baseQuery := `
	SELECT p.id, p.title, p.url, p.user_id, u.name AS user_name, p.created_at,
	(SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS comment_count,
	(SELECT COUNT(*) FROM votes v WHERE v.post_id = p.id) AS vote_count,
	COUNT(*) OVER() AS total_records
	FROM posts p
	INNER JOIN users u ON p.user_id = u.id
	`

	var args []interface{}

	if filter.Query != "" {
		baseQuery += "WHERE lower(p.title) LIKE ?"
		args = append(args, "%"+strings.ToLower(filter.Query)+"%")
	}

	if filter.OrderBy == "popular" {
		baseQuery += " ORDER BY vote_count, p.created_at DESC"
	} else {
		baseQuery += " ORDER BY p.created_at DESC"
	}

	limit := filter.PageSize
	offset := (filter.Page - 1) * filter.PageSize
	baseQuery += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	preparedStmt, err := r.db.Prepare(baseQuery)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer preparedStmt.Close()

	rows, err := preparedStmt.Query(args...)
	totalRecords := 0

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.URL,
			&post.UserID,
			&post.UserName,
			&post.CreatedAt,
			&post.CommentCount,
			&post.VoteCount,
			&totalRecords,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		post.TotalRecords = totalRecords
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	if len(posts) == 0 {
		return []Post{}, Metadata{}, nil
	}

	return posts, calculateMetadata(totalRecords, filter.Page, filter.PageSize), nil
}

func (r *SqlPostRepository) GetById(id int) (*Post, error) {
	query := `
	SELECT p.id, p.title, p.url, p.user_id, u.name AS user_name, p.created_at,
	(SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS comment_count,
	(SELECT COUNT(*) FROM votes v WHERE v.post_id = p.id) AS vote_count
	FROM posts p
	INNER JOIN users u ON p.user_id = u.id
	WHERE p.id = ?
	`

	row := r.db.QueryRow(query, id)
	var post Post
	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.URL,
		&post.UserID,
		&post.UserName,
		&post.CreatedAt,
		&post.CommentCount,
		&post.VoteCount,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPostNotFound
		}
	}

	return &post, nil
}

func (r *SqlPostRepository) GetComments(postId int) ([]Comment, error) {
	query := `
	SELECT c.id, c.comment, c.user_id, c.post_id, u.name AS user_name, c.created_at
	FROM comments c
	INNER JOIN users u ON c.user_id = u.id
	WHERE c.post_id = ?
	ORDER BY c.created_at DESC
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(postId)
	if err != nil {
		return nil, err
	}

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Comment,
			&comment.UserID,
			&comment.PostID,
			&comment.UserName,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return []Comment{}, nil
	}

	return comments, nil
}

func (p *Post) GetVoteCountsHuman() string {
	if p.VoteCount > 1 {
		return fmt.Sprintf("%d votes", p.VoteCount)
	}
	return fmt.Sprintf("%d vote", p.VoteCount)
}

func (p *Post) GetCommentCountsHuman() string {
	if p.CommentCount > 1 {
		return fmt.Sprintf("%d comments", p.CommentCount)
	}
	return fmt.Sprintf("%d comment", p.CommentCount)
}

func (p *Post) GetCreatedAtHuman() string {
	return carbon.NewCarbon(p.CreatedAt).DiffForHumans()
}

func (p *Post) Host() string {
	ur, err := url.Parse(p.URL)
	if err != nil {
		return "<invalid-host>"
	}
	return ur.Hostname()
}
