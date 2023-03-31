package controllers

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"golang-starter/src/restapi/config"
	"golang-starter/src/restapi/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// type Image struct {
// 	Front *multipart.FileHeader `form:"front" binding:"required"`
// }

type user struct {
	Id     int                   `uri:"id"`
	Name   string                `form:"name"`
	Email  string                `form:"email"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	config.InitDb().Find(&books)
	// models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author}
	config.InitDb().Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Find a book
func FindBook(c *gin.Context) { // Get model if exist
	var book models.Book

	if err := config.InitDb().Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := config.InitDb().Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.InitDb().Model(book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := config.InitDb().Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.InitDb().Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// func ImageUpload(c *gin.Context) {
// 	// var book models.Book
// 	// if err := config.InitDb().Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 	// 	return
// 	// }
// 	var userObj user
// 	// if err := c.ShouldBind(&userObj); err != nil {
// 	// 	c.String(http.StatusBadRequest, "bad request")
// 	// }

// 	// if err := c.ShouldBindUri(&userObj); err != nil {
// 	// 	c.String(http.StatusBadRequest, "bad request")
// 	// }
// 	err := c.SaveUploadedFile(userObj.Avatar, "assets/"+userObj.Avatar.Filename)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "unknown error")
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": "ok"})

// }

func ImageUpload(c *gin.Context) {
	// single file
	file, err := c.FormFile("file")
	log.Println(file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Upload Failed!"})
		return
	}
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, "./src/assets/images/"+newFileName)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"Upload": fmt.Sprintf("'%s' uploaded!", file.Filename)})
	// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
