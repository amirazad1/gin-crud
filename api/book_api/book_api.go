package book_api

import (
	"github.com/amirazad1/gin-crud/models"
	"github.com/amirazad1/gin-crud/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAll(context *gin.Context) {
	books, err := service.BookServ.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    books,
	})
}

func GetByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	book, err := service.BookServ.GetById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    book,
	})
}

func Create(context *gin.Context) {
	var item *models.Book
	err := context.ShouldBind(&item)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid item",
		})
		return
	}
	err = service.BookServ.Create(item)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func Update(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	var item *models.Book
	err = context.ShouldBind(&item)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid item",
		})
		return
	}

	err = service.BookServ.Update(id, item)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func Delete(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	err = service.BookServ.Delete(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
