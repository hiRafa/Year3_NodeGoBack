package main

import (
	"api.com/routes"
	"api.com/sqldb"
	"github.com/gin-gonic/gin"
)

func main() {
	sqldb.InitDB()

	//server is a pointer because of gin package
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
