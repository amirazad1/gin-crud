package service

import (
	"github.com/amirazad1/gin-crud/reposirotry/mysql"
)

var BookServ *BookService

func Setup() {
	//repo := memory.NewBookRepository()
	//repo := mysql.NewBookRepository()
	repo := mysql.NewBookXRepository()
	BookServ = NewBookService(repo)
}
