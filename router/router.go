package router

import (
	"github.com/amirazad1/gin-crud/pkg/setting"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)
	server := gin.Default()
	registerRoutes(server)
	return server
}

func registerRoutes(server *gin.Engine) {
	//baseUrl := setting.ServerSetting.BaseUrl
	root := server.Group("/")
	//api := server.Group(baseUrl)

	root.GET("ping", func(context *gin.Context) {
		response := "pong"
		_, _ = context.Writer.Write([]byte(response))
	})
}
