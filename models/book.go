package models

import "time"

type Book struct {
	ID        int64     `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Author    string    `json:"author" form:"author"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}
