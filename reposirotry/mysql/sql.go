package mysql

import (
	"database/sql"
	"fmt"
	"github.com/amirazad1/gin-crud/models"
	"github.com/amirazad1/gin-crud/pkg/setting"
	"log"
	"time"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository() *BookRepository {
	mdb, err := sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	))

	if err != nil {
		log.Fatal(err)
	}

	return &BookRepository{
		db: mdb,
	}
}

func (repo *BookRepository) GetAll() (*[]models.Book, error) {
	query := "SELECT id,name,author,created_at FROM books"
	rows, err := repo.db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var items []models.Book
	for rows.Next() {
		var item models.Book
		err = rows.Scan(&item.ID, &item.Name, &item.Author, &item.CreatedAt)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	return &items, nil
}

func (repo *BookRepository) GetByID(id int64) (*models.Book, error) {
	query := "SELECT id,name,author,created_at FROM books WHERE id=?"
	row := repo.db.QueryRow(query, id)

	var item models.Book
	err := row.Scan(&item.ID, &item.Name, &item.Author, &item.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (repo *BookRepository) GetByName(name string) (*[]models.Book, error) {
	query := "SELECT id,name,author,created_at FROM books WHERE name=?"
	rows, err := repo.db.Query(query, name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var items []models.Book
	for rows.Next() {
		var item models.Book
		err = rows.Scan(&item.ID, &item.Name, &item.Author, &item.CreatedAt)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	return &items, nil
}

func (repo *BookRepository) Create(book *models.Book) error {
	query := "INSERT books (name,author,created_at) VALUES(?,?,?)"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(book.Name, book.Author, time.Now())
	if err != nil {
		return err
	}

	book.ID, _ = result.LastInsertId()
	return nil
}

func (repo *BookRepository) Update(id int64, book *models.Book) error {
	query := "UPDATE books SET name=?,author=? WHERE id=?"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(book.Name, book.Author, id)
	return err
}

func (repo *BookRepository) Delete(id int64) error {
	query := "DELETE FROM books WHERE id=?"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
