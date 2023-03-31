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
	r.GET("/download", func(c *gin.Context) {
		// c.up
	})
	r.POST("/upload", controllers.ImageUpload)

	r.GET("/firefox", controllers.Firefox)
	// for Nginx routing
	r.GET("/nginx_start", controllers.NginxStart)
	r.GET("/nginx_stop", controllers.NginxStop)
	r.GET("/nginx_restart", controllers.NginxRestart)
	r.GET("/nginx_config_set", controllers.NginxConfigSet)
	r.GET("/nginx_config_get", controllers.NginxConfigGet)
	r.GET("/nginx_config_test", controllers.NginxConfigTest)
	r.GET("/nginx_config_reload", controllers.NginxConfigReload)

	// for Redis routing
	r.GET("/redis_start", controllers.RedisStart)
	r.GET("/redis_stop", controllers.RedisStop)
	r.GET("/redis_restart", controllers.RedisRestart)
	r.GET("/redis_key_set", controllers.RedisKeySet)
	r.GET("/redis_key_get", controllers.RedisKeyGet)
	r.GET("/redis_keys_find", controllers.RedisKeysFind)
	r.GET("/redis_keys_delete", controllers.RedisKeysDelete)

	r.Run(":8084")
}
