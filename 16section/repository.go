package main

import (
	"context"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	InvalidCredentialsErr = errors.New("invalid credentials")
)

type UserRepository interface {
	CreateUser(name, email, password, avatar string) (int, error)
	GetUserByEmail(email string) (*User, error)
	GetUsers() ([]User, error)
	Authenticate(email, plainPassword string) (int, error)
}

type SqlUserRepository struct {
	db *sql.DB
}

func NewSqlUserRepository(db *sql.DB) UserRepository {
	return &SqlUserRepository{
		db: db,
	}
}

func (r *SqlUserRepository) CreateUser(name, email, plainPassword, avatar string) (int, error) {
	ctx := context.Background()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, email, hashPass)
	if err != nil {
		return 0, err
	}

	UserID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	profileStm, err := tx.PrepareContext(ctx, `INSERT INTO profile(user_id, avatar) VALUES (?, ?)`)
	if err != nil {
		return 0, err
	}
	defer profileStm.Close()

	_, err = profileStm.Exec(UserID, avatar)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, nil
	}

	return int(UserID), nil
}

func (r *SqlUserRepository) Authenticate(email, plainPassword string) (int, error) {
	user, err := r.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, InvalidCredentialsErr
		}
		return 0, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(plainPassword)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, InvalidCredentialsErr
		}
		return 0, err
	}
	return user.ID, nil
}

func (r *SqlUserRepository) GetUserByEmail(email string) (*User, error) {
	stmt := `SELECT id, name, email, hashed_password, u.created_at, p.user_id, p.avatar, p.created_at FROM users u INNER JOIN profile p ON u.id = p.user_id WHERE email = ?`
	row := r.db.QueryRow(stmt, email)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.Profile.UserID, &user.Profile.Avatar, &user.Profile.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *SqlUserRepository) GetUsers() ([]User, error) {
	var users []User
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users`
	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
