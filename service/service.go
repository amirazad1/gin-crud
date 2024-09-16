package service

import "github.com/amirazad1/gin-crud/reposirotry"

var BookServ *BookService

func Setup() {
	repo := reposirotry.NewMemoryBookRepository()
	BookServ = NewBookService(repo)
}
