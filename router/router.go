package router

import (
	"github.com/amirazad1/gin-crud/api/book_api"
	"github.com/amirazad1/gin-crud/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)
	server := gin.Default()
	registerRoutes(server)
	return server
}

func registerRoutes(server *gin.Engine) {
	baseUrl := setting.ServerSetting.BaseUrl
	root := server.Group("/")
	api := server.Group(baseUrl)

	root.GET("ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	api.GET("books", book_api.GetAll)
	api.GET("books/:id", book_api.GetByID)
	api.POST("books", book_api.Create)
	api.PATCH("books/:id", book_api.Update)
	api.DELETE("books/:id", book_api.Delete)
}
