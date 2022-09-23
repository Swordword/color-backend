package models

import "colorist/config"

type Color struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
	Story string `json:"story"`
}

// 创建颜色
func CreateAColor(color *Color) (err error) {
	err = config.DB.Create(&color).Error
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
