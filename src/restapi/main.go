package main

import (
	"golang-starter/src/restapi/config"
	"golang-starter/src/restapi/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.Migrate()
	config.InitDb()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Static("assets", "./assets")
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.GET("/download",func(c *gin.Context){
		// c.up
	})
	r.POST("/upload", controllers.ImageUpload)

	r.Run(":8084")
}
