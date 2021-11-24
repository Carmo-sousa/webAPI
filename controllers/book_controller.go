package controllers

import (
	"net/http"
	"strconv"

	"github.com/Carmo-sousa/webAPI/database"
	"github.com/Carmo-sousa/webAPI/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "O id deve ser um inteiro.",
		})
		return
	}

	db := database.GetDataBase()
	var book models.Book
	err = db.First(&book, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "O livro não foi encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDataBase()

	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "cannot bind JSON" + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "cannot create a book" + err.Error(),
		})
		return
	}

	if book.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "O livro não pode ter canpos vazios!",
		})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func ShowAllBooks(c *gin.Context) {
	db := database.GetDataBase()

	var p []models.Book
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "cannot list books: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDataBase()
	idParam := c.Param("id")
	bookID, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "O id deve ser um inteiro!",
		})
	}

	var book models.Book
	c.ShouldBindJSON(&book)

	db.Model(models.Book{}).Where("id = ?", bookID).Updates(&book)
	c.JSON(http.StatusCreated, book)
}

func DeleteBook(c *gin.Context) {
	db := database.GetDataBase()
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "O id deve ser um inteiro!",
		})
	}

	var book models.Book
	db.Delete(&book, id)

	c.JSON(http.StatusOK, book)
}
