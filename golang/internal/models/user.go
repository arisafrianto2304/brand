package models

import (
	"time"
)

type User struct {
	UserID     int       `json:"-"`
	Username   string    `json:"username" example:"johndoe"` // username of the user
	Password   string    `json:"password" example:"s3cr3t"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	SoftDelete time.Time `json:"-"`
}
