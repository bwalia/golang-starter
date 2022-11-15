package controllers

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func RedisStart(c *gin.Context) {

	cmd := exec.Command("powershell", "-c", "start nginx")
	cmd.Dir = "E:/server/nginx"
	out, err := cmd.Output()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not run command"})
		fmt.Println("could not run command: ", err)
	} else {
		fmt.Println("Output: ", string(out))
		c.JSON(http.StatusOK, gin.H{"data": "Nginx Started OK"})
	}

}

func RedisStop(c *gin.Context) {

	cmd := exec.Command("powershell", "-c", "nginx -s stop")
	cmd.Dir = "E:/server/nginx"
	out, err := cmd.Output()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not run command"})
		fmt.Println("could not run command: ", err)
	} else {
		fmt.Println("Output: ", string(out))
		c.JSON(http.StatusOK, gin.H{"data": "Nginx Stopped Successfully!"})
	}

}

func RedisRestart(c *gin.Context) {

	cmd := exec.Command("powershell", "-c", "nginx -s stop")
	cmd.Dir = "E:/server/nginx"
	out, err := cmd.Output()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not run command"})
		fmt.Println("could not run command: ", err)
	} else {
		fmt.Println("Output: ", string(out))
		c.JSON(http.StatusOK, gin.H{"data": "Nginx Restarted Successfully!"})
	}

}

func RedisKeySet(c *gin.Context) {

	cmd := exec.Command("powershell", "-c", "nginx -s stop")
	cmd.Dir = "E:/server/nginx"
	out, err := cmd.Output()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not run command"})
		fmt.Println("could not run command: ", err)
	} else {
		fmt.Println("Output: ", string(out))
		c.JSON(http.StatusOK, gin.H{"data": "Nginx config set!"})
	}

}

func RedisKeyGet(c *gin.Context) {

	cmd := exec.Command("powershell", "-c", "nginx -s stop")
	cmd.Dir = "E:/server/nginx"
	out, err := cmd.Output()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not run command"})
		fmt.Println("could not run command: ", err)
	} else {
		fmt.Println("Output: ", string(out))
		c.JSON(http.StatusOK, gin.H{"data": "Nginx config get!"})
	}

}

func RedisKeysFind(c *gin.Context) {

	cmd := exec.Command("powershell", "-c", "nginx -s stop")
	cmd.Dir = "E:/server/nginx"
	out, err := cmd.Output()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not run command"})
		fmt.Println("could not run command: ", err)
	} else {
		fmt.Println("Output: ", string(out))
		c.JSON(http.StatusOK, gin.H{"data": "Nginx configtest OK!"})
	}

}

func RedisKeysDelete(c *gin.Context) {

	cmd := exec.Command("powershell", "-c", "nginx -s stop")
	cmd.Dir = "E:/server/nginx"
	out, err := cmd.Output()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not run command"})
		fmt.Println("could not run command: ", err)
	} else {
		fmt.Println("Output: ", string(out))
		c.JSON(http.StatusOK, gin.H{"data": "Nginx config reload done.."})
	}

}
