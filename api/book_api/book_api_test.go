package bookapi_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/amirazad1/gin-crud/models"
	"github.com/amirazad1/gin-crud/pkg/setting"
	"github.com/amirazad1/gin-crud/router"
	"github.com/amirazad1/gin-crud/service"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	baseURL = "http://127.0.0.1:8080/api/v1/"
	server  *gin.Engine
)

type response struct {
	Message string      `json:"message"`
	Data    models.Book `json:"data"`
}

type responseAll struct {
	Message string        `json:"message"`
	Data    []models.Book `json:"data"`
}

func TestMain(m *testing.M) {
	setting.Setup("../../config/app.ini")
	service.Setup()
	server = router.Setup()
	baseURL = setting.ServerSetting.TestBaseURL
	m.Run()
}

func createFakePostData() (models.Book, *multipart.Writer, *bytes.Buffer) {
	b := models.Book{}
	_ = faker.FakeData(&b)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("name", b.Name)
	_ = writer.WriteField("author", b.Author)
	_ = writer.Close()

	return b, writer, payload
}

func createEmptyPostData() (*multipart.Writer, *bytes.Buffer) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("name", "")
	_ = writer.WriteField("author", "")
	_ = writer.Close()

	return writer, payload
}

func createNullPostData() (*multipart.Writer, *bytes.Buffer) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.Close()

	return writer, payload
}

func createBook(t *testing.T) models.Book {
	book, writer, payload := createFakePostData()
	req, _ := http.NewRequest("POST", baseURL+"books", payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)
	var result response
	err := json.Unmarshal(rr.Body.Bytes(), &result)

	assert.NoError(t, err)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, book.Name, result.Data.Name)
	assert.Equal(t, book.Author, result.Data.Author)

	book.ID = result.Data.ID
	return book
}

func TestCreateBook(t *testing.T) {
	createBook(t)
}

func TestCreateBookWithEmptyForm(t *testing.T) {
	writer, payload := createEmptyPostData()
	req, _ := http.NewRequest("POST", baseURL+"books", payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestCreateBookWithNullForm(t *testing.T) {
	writer, payload := createNullPostData()
	req, _ := http.NewRequest("POST", baseURL+"books", payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusInternalServerError)
}

func TestGetBookByID(t *testing.T) {
	book := createBook(t)
	req, _ := http.NewRequest("GET", baseURL+"books/"+fmt.Sprintf("%d", book.ID), nil)
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)
	var result response
	err := json.Unmarshal(rr.Body.Bytes(), &result)

	assert.NoError(t, err)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.ObjectsAreEqual(book, result.Data)
}

func TestGetBookByStringID(t *testing.T) {
	req, _ := http.NewRequest("GET", baseURL+"books/"+faker.UUIDHyphenated(), nil)
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestUpdateBook(t *testing.T) {
	book := createBook(t)

	newBook, writer, payload := createFakePostData()

	req, _ := http.NewRequest("PATCH", baseURL+"books/"+fmt.Sprintf("%d", book.ID), payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)
	var result response
	err := json.Unmarshal(rr.Body.Bytes(), &result)

	assert.NoError(t, err)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.ObjectsAreEqual(newBook, result.Data)
	assert.NotEqualValues(t, book, result.Data)
}

func TestDeleteBook(t *testing.T) {
	book := createBook(t)

	req, _ := http.NewRequest("DELETE", baseURL+"books/"+fmt.Sprintf("%d", book.ID), nil)
	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)

	req, _ = http.NewRequest("GET", baseURL+"books/"+fmt.Sprintf("%d", book.ID), nil)
	rr = httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Error(t, sql.ErrNoRows)
}

func TestGetAllBook(t *testing.T) {
	req, _ := http.NewRequest("GET", baseURL+"books", nil)
	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	var resultOld responseAll
	err := json.Unmarshal(rr.Body.Bytes(), &resultOld)
	assert.NoError(t, err)
	assert.Equal(t, rr.Code, http.StatusOK)

	for i := 0; i < 10; i++ {
		createBook(t)
	}

	req, _ = http.NewRequest("GET", baseURL+"books", nil)
	rr = httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	var result responseAll
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, rr.Code, http.StatusOK)

	assert.Equal(t, len(result.Data), len(resultOld.Data)+10)
	for _, book := range result.Data {
		assert.NotEmpty(t, book)
	}
}
