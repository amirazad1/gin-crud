package models

import "time"

type Book struct {
	ID        int64     `json:"id" form:"id" db:"id"` //nolint:govt
	Name      string    `json:"name" form:"name" db:"name" binding:"required" faker:"name"`
	Author    string    `json:"author" form:"author" db:"author" binding:"required" faker:"last_name"`
	CreatedAt time.Time `json:"created_at" form:"created_at" db:"created_at"`
}
