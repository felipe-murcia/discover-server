package main

import (
	"net/http"

	"discoverco.co/server/configs"
	"discoverco.co/server/routes"
	"github.com/gin-gonic/gin"
	//cors "github.com/itsjamie/gin-cors"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	// Apply the middleware to the router (works with groups too)
	/*
		r.Use(cors.Middleware(cors.Config{
			Origins:         "*",
			Methods:         "GET, PUT, POST, DELETE",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			Credentials:     false,
			ValidateHeaders: false,
		}))

	*/
	routes.PersonRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run()
}
