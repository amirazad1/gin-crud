package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()
	registerRoutes(server)
	return server
}

func registerRoutes(server *gin.Engine) {
	//baseUrl := "api/v1/"
	root := server.Group("/")
	//api := server.Group(baseUrl)

	root.GET("ping", func(context *gin.Context) {
		response := "pong"
		_, err := context.Writer.Write([]byte(response))
		log.Fatal(err)
	})
}
