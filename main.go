package main

import (
	"fmt"

	"example.com/restapi/db"
	"example.com/restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("server is starting")
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
	fmt.Println("server is running")
}
