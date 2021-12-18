package models

import (
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/jinzhu/gorm"
	"time"
)

type Bill struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	CategoryId int   `json:"category_id"`
	Amount     int   `json:"amount"`
	Image     string  `json:"image"`
	Remark    string  `json:"remark"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdateAt   time.Time  `json:"update_at"`
}
// 判断id是否存在
func ExistBillById(id int) (bool, error) {
	var bill Bill
	err := db.Select("id").Where("id = ?", id).First(&bill).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if bill.Id > 0 {
		return true, nil
	}

	return false, err
}
// 创建账单
func CreateBill(billParam map[string]interface{}) error {
	bill := Bill{
		UserId: billParam["user_id"].(int),
		CategoryId: billParam["category_id"].(int),
		Amount: billParam["amount"].(int),
		Image: billParam["image"].(string),
		Remark: billParam["remark"].(string),
		CreatedAt: util.ParseTime(billParam["created_at"].(string)),
		UpdateAt: time.Now(),
	}


	if err := db.Create(&bill).Error; err != nil {
		return err
	}

	return nil
}

// 更新
func UpdateBill(id int, data interface{}) error {
	err := db.Model(&Bill{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}

// 删除
func DeleteBill(id int) error {
	if err := db.Where("id = ?", id).Delete(&Bill{}).Error; err != nil {
		return err
	}

	return nil
}