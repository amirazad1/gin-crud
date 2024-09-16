package book_api

import (
	"github.com/amirazad1/gin-crud/models"
	"github.com/amirazad1/gin-crud/pkg/msg"
	"github.com/amirazad1/gin-crud/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAll(context *gin.Context) {
	books, err := service.BookServ.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.ERROR,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": msg.SUCCESS,
		"data":    books,
	})
}

func GetByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": msg.INVALID_ID,
		})
		return
	}

	book, err := service.BookServ.GetById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.ERROR,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": msg.SUCCESS,
		"data":    book,
	})
}

func Create(context *gin.Context) {
	var item *models.Book
	err := context.ShouldBind(&item)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": msg.INVALID_FORM,
		})
		return
	}
	err = service.BookServ.Create(item)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.ERROR,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": msg.SUCCESS,
	})
}

func Update(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": msg.INVALID_ID,
		})
		return
	}

	var item *models.Book
	err = context.ShouldBind(&item)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": msg.INVALID_FORM,
		})
		return
	}

	err = service.BookServ.Update(id, item)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.ERROR,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": msg.SUCCESS,
	})
}

func Delete(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": msg.INVALID_ID,
		})
		return
	}

	err = service.BookServ.Delete(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.ERROR,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": msg.SUCCESS,
	})
}
