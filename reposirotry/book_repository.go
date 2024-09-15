package reposirotry

import "github.com/amirazad1/gin-crud/models"

type BookRepository interface {
	GetAll() (*[]models.Book, error)
	GetById(int64) (*models.Book, error)
	GetByName(string) (*models.Book, error)
	Create(*models.Book) error
	Update(*models.Book) error
	Delete(int64) error
}
