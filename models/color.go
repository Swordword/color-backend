package models

import (
	"colorist/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Color struct {
	gorm.Model
	Value string `json:"value"`
	Story string `json:"story"`
	Stars uint   `json:"stars"`
}

// 创建颜色
func CreateAColor(color *Color) (err error) {
	err = config.DB.Create(&color).Error
	return
}

// 分页获取颜色
func GetPagerColor(c *gin.Context) (colorList []*Color, err error) {
	if err = config.DB.Scopes(Paginate(c)).Find(&colorList).Error; err != nil {
		return nil, err
	}
	return
}

// 获取所有颜色
func GetAllColor() (colorList []*Color, err error) {
	if err = config.DB.Find(&colorList).Error; err != nil {
		return nil, err
	}
	return
}

// 删除颜色
func DeleteAColor(id string) (err error) {
	err = config.DB.Where("id = ?", id).Delete(&Color{}).Error
	return
}

// star 颜色
func StarAColor(id string) (err error) {
	var color Color
	config.DB.First(&color, id)
	fmt.Println("stars: ", color.Stars)
	config.DB.Model(&color).Where("id = ?", id).Update("stars", color.Stars+1)
	return err
}

// 取消star
func CancelStar(id string) (err error) {
	var color Color
	config.DB.First(&color, id)
	config.DB.Model(&color).Where("id = ?", id).Update("stars", color.Stars-1)
	return err
}
