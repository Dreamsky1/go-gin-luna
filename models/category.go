package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Category struct {
	Id			int		`json:"id"`
	Name 		string 	`json:"name"`
	Type 		int		`json:"type"`
	Image       string	`json:"image"`
	CreatedBy	int		`json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

const CATEGORY_TYPE_EXPENSES = 0
const CATEGORY_TYPE_INCOME = 1

var CATEGORY_TYPE_COODE2TYPE = map[string]int{
	"expenses": CATEGORY_TYPE_EXPENSES,
	"income": CATEGORY_TYPE_INCOME,
}

var TYPE2CATEGORY_TYPE_CODE = map[int]string{
	CATEGORY_TYPE_EXPENSES: "expenses",
	CATEGORY_TYPE_INCOME: "income",
}

// 创建分类
func CreateCategory(categoryParam map[string]interface{}) error {
	category := Category{
		Name: categoryParam["name"].(string),
		Type: CATEGORY_TYPE_COODE2TYPE[categoryParam["type"].(string)],
		CreatedBy: categoryParam["createdBy"].(int),
		Image: categoryParam["image"].(string),
		CreatedAt: time.Now(),
	}

	if err := db.Create(&category).Error; err != nil {
		return err
	}

	return nil
}

// 判断是否存在
func ExistCategoryName(name string) (bool, error) {
	var category Category
	err := db.Select("id").Where("name = ?", name).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return false, err
	}

	if category.Id > 0 {
		return true, nil
	}
	return false, nil
}
