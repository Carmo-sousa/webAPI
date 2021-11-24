package controllers

import (
	"net/http"
	"strconv"

	"github.com/Carmo-sousa/webAPI/database"
	"github.com/Carmo-sousa/webAPI/models"
	"github.com/Carmo-sousa/webAPI/services"
	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`
	DeletedAt string `json:"deleted"`
}

func CreateUser(c *gin.Context) {
	db := database.GetDataBase()
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Não foi possivel criar o usuário: " + err.Error(),
		})
		return
	}

	if user.Name == "" || user.Password == "" || user.Email == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "O usuário não pode ter campos vazios",
		})
		return
	}

	user.Password = services.SHA265Encoder(user.Password)

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Não foi possivel criar o usuário: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": "Usuário criado.",
	})

}

func ShowAllUsers(c *gin.Context) {
	db := database.GetDataBase()
	var users []UserAPI

	err := db.Model(&models.User{}).Find(&users).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func ShowUser(c *gin.Context) {
	IDparam := c.Param("id")
	db := database.GetDataBase()

	id, err := strconv.Atoi(IDparam)

	var user UserAPI

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = db.Model(&models.User{}).First(&user, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
