package controller

import (
	"colorist/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAColor(c *gin.Context) {
	var color models.Color
	c.BindJSON(&color)
	err := models.CreateAColor(&color)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, color)
	}
}

func GetColorList(c *gin.Context) {
	colorList, err := models.GetAllColor()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, colorList)
	}
}

func DeleteAColor(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	if err := models.DeleteAColor(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
