package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlshorter/services"
)

func Auth(c *gin.Context) {
	var request services.RegisterRequest
	err := c.ShouldBindJSON(&request)

	if err != nil || request.Name == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Bad request",
		})

		return
	}

	authService := services.AuthService{}
	token := authService.Auth(request)

	c.JSON(http.StatusOK, services.AuthResponse{Status: true, Token: token})
}
