package controllers

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func NginxStart(c *gin.Context) {

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

func NginxStop(c *gin.Context) {

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

func NginxRestart(c *gin.Context) {

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

func NginxConfigSet(c *gin.Context) {

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

func NginxConfigGet(c *gin.Context) {

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

func NginxConfigTest(c *gin.Context) {

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

func NginxConfigReload(c *gin.Context) {

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

func Firefox(c *gin.Context) {

	cmd := exec.Command("powershell", "start firefo")

	err := cmd.Run()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// log.Fatal(err)
	}
}
