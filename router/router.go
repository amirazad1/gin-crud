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
	baseURL := setting.ServerSetting.BaseURL
	root := server.Group("/")
	api := server.Group(baseURL)

	root.GET("ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	api.GET("books", bookapi.GetAll)
	api.GET("books/:id", bookapi.GetByID)
	api.POST("books", bookapi.Create)
	api.PATCH("books/:id", bookapi.Update)
	api.DELETE("books/:id", bookapi.Delete)
}
