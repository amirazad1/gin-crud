package mysql

import (
	"fmt"
	"github.com/amirazad1/gin-crud/models"
	"github.com/amirazad1/gin-crud/pkg/setting"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type BookXRepository struct {
	db *sqlx.DB
}

func NewBookXRepository() *BookXRepository {
	mdb, err := sqlx.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	))

	if err != nil {
		log.Fatal(err)
	}

	return &BookXRepository{
		db: mdb,
	}
}

func (repo *BookXRepository) GetAll() (*[]models.Book, error) {
	query := "SELECT id,name,author,created_at FROM books"
	var items []models.Book
	err := repo.db.Select(&items, query)
	if err != nil {
		return nil, err
	}

	return &items, nil
}

func (repo *BookXRepository) GetByID(id int64) (*models.Book, error) {
	query := "SELECT id,name,author,created_at FROM books WHERE id=?"
	var item models.Book
	err := repo.db.Get(&item, query, id)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (repo *BookXRepository) GetByName(name string) (*[]models.Book, error) {
	query := "SELECT id,name,author,created_at FROM books WHERE name=?"
	var items []models.Book
	err := repo.db.Select(&items, query, name)
	if err != nil {
		return nil, err
	}

	return &items, nil
}

func (repo *BookXRepository) Create(book *models.Book) error {
	query := "INSERT books (name,author,created_at) VALUES(:name,:author,:created_at)"
	book.CreatedAt = time.Now()
	result, err := repo.db.NamedExec(query, book)
	if err != nil {
		return err
	}

	book.ID, _ = result.LastInsertId()
	return nil
}

func (repo *BookXRepository) Update(id int64, book *models.Book) error {
	query := "UPDATE books SET name=:name,author=:author WHERE id=:id"
	book.ID = id
	result, err := repo.db.NamedExec(query, book)
	if err != nil {
		return err
	}

	book.ID, _ = result.LastInsertId()
	return nil
}

func (repo *BookXRepository) Delete(id int64) error {
	query := "DELETE FROM books WHERE id=?"
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	return err
}
