package main

import (
	"fmt"
	"github.com/amirazad1/gin-crud/pkg/setting"
	"github.com/amirazad1/gin-crud/router"
	"github.com/amirazad1/gin-crud/service"
	"log"
)

func init() {
	setting.Setup("config/app.ini")
	service.Setup()
}

func main() {
	server := router.Setup()
	fmt.Printf("Server starting on port %v ", setting.ServerSetting.HttpPort)
	err := server.Run(fmt.Sprintf(":%d", setting.ServerSetting.HttpPort))
	if err != nil {
		log.Fatal(err)
	}

}
