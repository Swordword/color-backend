package controller

import (
	"colorist/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建颜色
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
	colorList, err := models.GetPagerColor(c)
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

// 点赞
func StarAColor(c *gin.Context) {
	id, ok := c.Params.Get("id")
	logger := log.Default()
	logger.Printf("id:", id)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
	}
	fmt.Printf("CreateColor id: %v\n", id)
	if err := models.StarAColor(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "点赞成功"})
	}
}

func CancelStar(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
	}
	if err := models.CancelStar(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "取消失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "取消成功"})
	}
}
