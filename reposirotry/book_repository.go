package reposirotry

import "github.com/amirazad1/gin-crud/models"

type BookRepository interface {
	GetAll() (*[]models.Book, error)
	GetById(int64) (*models.Book, error)
	GetByName(string) (*models.Book, error)
	Create(book *models.Book) error
	Update(book *models.Book) error
	Delete(int64) error
}
