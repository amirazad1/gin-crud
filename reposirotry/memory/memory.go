package memory

import (
	"github.com/amirazad1/gin-crud/models"
	"time"
)

type BookRepository struct {
	db *[]models.Book
}

func NewBookRepository() *BookRepository {
	var books []models.Book
	return &BookRepository{
		db: &books,
	}
}

func (repo *BookRepository) GetAll() (*[]models.Book, error) {
	return repo.db, nil
}

func (repo *BookRepository) GetById(id int64) (*models.Book, error) {
	found := -1
	for index, value := range *repo.db {
		if id == value.ID {
			found = index
			break
		}
	}

	var book models.Book
	if found != -1 {
		book = (*repo.db)[found]
	}
	return &book, nil
}

func (repo *BookRepository) GetByName(name string) (*[]models.Book, error) {
	var books []models.Book
	for _, value := range *repo.db {
		if name == value.Name {
			books = append(books, value)
			break
		}
	}
	return &books, nil
}

func (repo *BookRepository) Create(book *models.Book) error {
	book.CreatedAt = time.Now()
	book.ID = int64(len(*repo.db)) + 1
	*repo.db = append(*repo.db, *book)
	return nil
}

func (repo *BookRepository) Update(id int64, book *models.Book) error {
	for index, value := range *repo.db {
		if id == value.ID {
			(*repo.db)[index].Name = book.Name
			(*repo.db)[index].Author = book.Author
			break
		}
	}
	return nil
}

func (repo *BookRepository) Delete(id int64) error {
	found := -1
	for index, value := range *repo.db {
		if id == value.ID {
			found = index
			break
		}
	}

	*repo.db = append((*repo.db)[:found], (*repo.db)[found+1:]...)
	return nil
}
