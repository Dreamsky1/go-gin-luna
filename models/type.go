package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TypeCategory struct {
	Model

	Name             string      `json:"name"`
	SecondCategories []*Category `json:"secondCategories"`
}

func (typeCategory *TypeCategory) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// 获得列表
func GetTypeCategories(pageNum int, pageSize int, maps interface{}) ([]*TypeCategory, error) {
	var typeCategories []*TypeCategory
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&typeCategories).Error

	// 填充category
	//var typeId2Category = make(map[int]*Category)
	for _, typeCategory := range typeCategories {
		//typeId2Category[typeCategory.ID] =
		var categories []*Category
		db.Where("type_id = ?", typeCategory.ID).Find(&categories)
		typeCategory.SecondCategories = categories
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return typeCategories, nil
}

// 获得的是数量
func GetAllTypeCategoryTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&TypeCategory{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func AddType(name string) error {
	if err := db.Create(&TypeCategory{
		Name: name,
	}).Error; err != nil {
		return err
	}

	return nil
}
