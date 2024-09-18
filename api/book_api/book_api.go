package bookapi

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
			"message": msg.Error + err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": msg.Success,
		"data":    books,
	})
}

func GetByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": msg.InvalidID,
		})
		return
	}

	book, err := service.BookServ.GetByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.Error + err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": msg.Success,
		"data":    book,
	})
}

func Create(context *gin.Context) {
	var item *models.Book
	err := context.ShouldBind(&item)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": msg.InvalidForm,
		})
		return
	}
	err = service.BookServ.Create(item)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.Error + err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": msg.Success,
		"data":    item,
	})
}

func Update(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": msg.InvalidID,
		})
		return
	}

	var item *models.Book
	err = context.ShouldBind(&item)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": msg.InvalidForm,
		})
		return
	}

	err = service.BookServ.Update(id, item)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.Error + err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": msg.Success,
		"data":    item,
	})
}

func Delete(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": msg.InvalidID,
		})
		return
	}

	err = service.BookServ.Delete(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.Error + err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": msg.Success,
	})
}
