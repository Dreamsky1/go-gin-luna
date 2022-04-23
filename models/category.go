package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Category struct {
	Model

	Name   string `json:"name"`
	TypeId int    `json:"type_id"`
	State  int    `json:"state"`
	Image  string `json:"image"`
}

func (category *Category) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// 判断是否存在这个分类
func ExistCategoryByID(id int) bool {
	var category Category
	db.Select("id").Where("id = ?", id).First(&category)
	fmt.Print("输出这个category", category)
	if category.ID > 0 {
		return true
	}

	return false
}

// 删除
func DeleteCategory(id int) bool {
	if err := db.Where("id = ?", id).Delete(&Category{}).Error; err != nil {
		return false
	}
	return true
}

// 获得所有的分类
func GetCategories(pageNum int, pageSize int, maps interface{}) (categories []Category) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&categories)
	return
}

// 获得categories通过typeid
func GetCategoriesByTypeId(typeId int) ([]*Category, error) {
	var categories []*Category
	err := db.Where("type_id = ?", typeId).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil

}

func GetCategoryTotal(maps interface{}) (count int) {
	db.Model(&Category{}).Where(maps).Count(&count)
	return
}

func ExistCategoryByName(name string) bool {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return true
	}

	return false
}

func AddCategory(name string, state int, typeId int, image string) error {
	if err := db.Create(&Category{
		Name:   name,
		State:  state,
		TypeId: typeId,
		Image:  image,
	}).Error; err != nil {
		return err
	}
	return nil
}
