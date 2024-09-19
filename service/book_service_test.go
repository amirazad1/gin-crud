package service

import (
	"database/sql"
	"github.com/amirazad1/gin-crud/models"
	"github.com/amirazad1/gin-crud/pkg/setting"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	setting.Setup("../config/app.ini")
	Setup()
	m.Run()
}

func CreateFakeBook(t *testing.T) models.Book {
	// name := faker.Name()
	// author := faker.LastName()
	// book := &models.Book{
	//	 Name:   name,
	//	 Author: author,
	// }
	book := &models.Book{}
	err := faker.FakeData(&book)
	if err != nil {
		return models.Book{}
	}

	err = BookServ.Create(book)
	assert.NoError(t, err)
	assert.NotEmpty(t, book.ID)
	assert.NotZero(t, book.ID)

	return *book
}

func TestBookService_Create(t *testing.T) {
	CreateFakeBook(t)
}

func TestBookService_GetByID(t *testing.T) {
	book := CreateFakeBook(t)

	foundBook, err := BookServ.GetByID(book.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, foundBook)

	assert.Equal(t, book.ID, foundBook.ID)
	assert.Equal(t, book.Name, foundBook.Name)
	assert.Equal(t, book.Author, foundBook.Author)
}

func TestBookService_Update(t *testing.T) {
	book := CreateFakeBook(t)

	newName := faker.Name()
	newAuthor := faker.LastName()
	newBook := &models.Book{
		Name:   newName,
		Author: newAuthor,
	}
	err := BookServ.Update(book.ID, newBook)
	assert.NoError(t, err)
	assert.NotEmpty(t, newBook)

	assert.Equal(t, book.ID, newBook.ID)
	assert.Equal(t, newName, newBook.Name)
	assert.Equal(t, newAuthor, newBook.Author)

}

func TestBookService_Delete(t *testing.T) {
	book := CreateFakeBook(t)

	err := BookServ.Delete(book.ID)
	assert.NoError(t, err)

	foundBook, err := BookServ.GetByID(book.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, foundBook)
}

func TestBookService_GetAll(t *testing.T) {
	oldBooks, err := BookServ.GetAll()
	assert.NoError(t, err)

	for i := 0; i < 10; i++ {
		CreateFakeBook(t)
	}
	books, err := BookServ.GetAll()
	assert.NoError(t, err)
	// assert.GreaterOrEqual(t, len(*books), 10)
	// assert.Len(t, *books, len(*oldBooks)+10)
	assert.Equal(t, len(*books), len(*oldBooks)+10)
	for _, book := range *books {
		assert.NotEmpty(t, book)
	}
}
