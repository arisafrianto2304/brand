package models

import (
	"time"
)

type User struct {
	UserID     int       `json:"-"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	SoftDelete time.Time `json:"-"`
}

type UserResponse struct {
	Username string `json:"username" example:"arisafrianto"`
	Password string `json:"password" example:"arisafrianto12345"`
}
