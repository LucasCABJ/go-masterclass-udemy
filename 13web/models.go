package main

import "time"

// User represents a user entity in database.
type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	Profile        Profile   `json:"profile"`
}

// Profile represents a profile entity in database.
type Profile struct {
	UserID    int       `json:"user_id"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}