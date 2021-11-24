package controllers

import (
	"net/http"

	"github.com/Carmo-sousa/webAPI/database"
	"github.com/Carmo-sousa/webAPI/models"
	"github.com/Carmo-sousa/webAPI/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDataBase()

	var login models.Login
	err := c.ShouldBindJSON(&login)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cannot find user: " + err.Error(),
		})
		return
	}

	var user models.User

	err = db.Where("email = ?", login.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	if user.Password != services.SHA265Encoder(login.Password) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid credentials.",
		})
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
