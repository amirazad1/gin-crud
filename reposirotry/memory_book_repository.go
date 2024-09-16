package reposirotry

import (
	"github.com/amirazad1/gin-crud/models"
	"time"
)

type MemoryBookRepository struct {
	db *[]models.Book
}

func NewMemoryBookRepository() *MemoryBookRepository {
	var books []models.Book
	return &MemoryBookRepository{
		db: &books,
	}
}

func (repo *MemoryBookRepository) GetAll() (*[]models.Book, error) {
	return repo.db, nil
}

func (repo *MemoryBookRepository) GetById(id int64) (*models.Book, error) {
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

func (repo *MemoryBookRepository) GetByName(name string) (*models.Book, error) {
	found := -1
	for index, value := range *repo.db {
		if name == value.Name {
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

func (repo *MemoryBookRepository) Create(book *models.Book) error {
	book.CreatedAt = time.Now()
	book.ID = int64(len(*repo.db)) + 1
	*repo.db = append(*repo.db, *book)
	return nil
}

func (repo *MemoryBookRepository) Update(id int64, book *models.Book) error {
	for index, value := range *repo.db {
		if id == value.ID {
			(*repo.db)[index].Name = book.Name
			(*repo.db)[index].Author = book.Author
			break
		}
	}
	return nil
}

func (repo *MemoryBookRepository) Delete(id int64) error {
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
