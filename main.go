package main

import (
	
	"log"
	"crudemp/api"

	"github.com/gin-gonic/gin"
	
)

func main() {
	r := gin.Default()

	// Swagger documentation setup
	r.GET("/swagger/*any", gin.WrapH(api.GetSwaggerHandler()))

	// Define routes
	api.SetupRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
