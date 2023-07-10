package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlshorter/models"
	"urlshorter/services"
	"urlshorter/utils"
)

func AddLink(c *gin.Context) {
	var request services.AddRequest

	err := c.ShouldBindJSON(&request)

	if err != nil || !utils.IsValidURL(request.Link) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Bad request",
		})

		return
	}

	linkModel := &models.LinkModel{DB: utils.GetDBConnection()}

	slug, _ := linkModel.InsertLink(request.Link)

	c.JSON(http.StatusOK, services.AddResponse{Status: true, Slug: slug})
}
