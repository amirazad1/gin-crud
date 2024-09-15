package service

import (
	"github.com/amirazad1/gin-crud/models"
	"github.com/amirazad1/gin-crud/reposirotry"
)

type BookService struct {
	repo reposirotry.BookRepository
}

func NewBookService(repository reposirotry.BookRepository) *BookService {
	return &BookService{
		repo: repository,
	}
}

func (s *BookService) GetAll() (*[]models.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetById(id int64) (*models.Book, error) {
	return s.repo.GetById(id)
}

func (s *BookService) GetByName(name string) (*models.Book, error) {
	return s.repo.GetByName(name)
}

func (s *BookService) Create(book *models.Book) error {
	return s.repo.Create(book)
}

func (s *BookService) Update(book *models.Book) error {
	return s.repo.Update(book)
}

func (s *BookService) Delete(id int64) error {
	return s.repo.Delete(id)
}
