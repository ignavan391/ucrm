package models

import "time"

type User struct {
	Id        string    `json:"id"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	AvatarUrl *string   `json:"avatar_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
