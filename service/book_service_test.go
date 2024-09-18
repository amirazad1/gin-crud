package service

import (
	"database/sql"
	"github.com/amirazad1/gin-crud/models"
	"github.com/amirazad1/gin-crud/pkg/setting"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMain(m *testing.M) {
	setting.Setup("../config/app.ini")
	Setup()
	m.Run()
}

func createFakeBook(t *testing.T) models.Book {
	name := faker.LastName()
	author := faker.FirstName()
	book := &models.Book{
		Name:   name,
		Author: author,
	}

	err := BookServ.Create(book)
	require.NoError(t, err)
	require.NotEmpty(t, book.ID)
	require.NotZero(t, book.ID)

	return *book
}

func TestBookService_Create(t *testing.T) {
	createFakeBook(t)
}

func TestBookService_GetByID(t *testing.T) {
	book := createFakeBook(t)

	foundBook, err := BookServ.GetByID(book.ID)
	require.NoError(t, err)
	require.NotEmpty(t, foundBook)

	require.Equal(t, book.ID, foundBook.ID)
	require.Equal(t, book.Name, foundBook.Name)
	require.Equal(t, book.Author, foundBook.Author)

}

func TestBookService_Update(t *testing.T) {
	book := createFakeBook(t)

	newName := faker.LastName()
	newAuthor := faker.FirstName()
	newBook := &models.Book{
		Name:   newName,
		Author: newAuthor,
	}
	err := BookServ.Update(book.ID, newBook)
	require.NoError(t, err)
	require.NotEmpty(t, newBook)

	require.Equal(t, book.ID, newBook.ID)
	require.Equal(t, newName, newBook.Name)
	require.Equal(t, newAuthor, newBook.Author)

}

func TestBookService_Delete(t *testing.T) {
	book := createFakeBook(t)

	err := BookServ.Delete(book.ID)
	require.NoError(t, err)

	foundBook, err := BookServ.GetByID(book.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, foundBook)
}

func TestBookService_GetAll(t *testing.T) {
	oldBooks, err := BookServ.GetAll()
	require.NoError(t, err)

	for i := 0; i < 10; i++ {
		createFakeBook(t)
	}

	books, err := BookServ.GetAll()
	require.NoError(t, err)
	// require.GreaterOrEqual(t, len(*books), 10)
	// require.Len(t, *books, len(*oldBooks)+10)
	require.Equal(t, len(*books), len(*oldBooks)+10)
	for _, book := range *books {
		require.NotEmpty(t, book)
	}

}
