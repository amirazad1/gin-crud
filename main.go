package main

import (
	"fmt"
	"github.com/amirazad1/gin-crud/router"
	"log"
)

func init() {

}

func main() {
	server := router.Setup()
	fmt.Printf("Server starting on port %v ", 8080)
	err := server.Run(fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatal(err)
	}

}
