package models

import "time"

type Book struct {
	ID        int64     `json:"id" form:"id"`
	Name      string    `json:"name" form:"name" valid:"required"`
	Author    string    `json:"author" form:"author" valid:"required"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}
