package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Bill struct {
	Model

	TypeId       int      `json:"type_id"`
	CategoryId   int      `json:"category_id"`
	CategoryName string   `json:"category_name"`
	Amount       int      `json:"amount"`
	Remark       string   `json:"remark"`
	Category     Category `json:"category"`
}

// 判断是否存在
func ExistBillByID(id int) bool {
	var bill Bill
	db.Select("id").Where("id = ?", id).First(&bill)

	if bill.ID > 0 {
		return true
	}

	return false
}

// 获得所有数量
func GetBillTotal(maps interface{}) (count int) {
	db.Model(&Bill{}).Where(maps).Count(&count)

	return
}

// 获得所有的账单 跟 分类有关的, 这个maps是查询条件
func GetBills(pageNum int, pageSize int, maps interface{}) (bills []Bill) {
	db.Preload("Category").Where(maps).Offset(pageNum).Limit(pageSize).Find(&bills)

	return
}

// 获得单个的账单
func GetBill(id int) (bill Bill) {
	db.Where("id = ?", id).First(&bill)
	db.Model(&bill).Related(&bill.Category)

	return
}

// 更新账单
func EditBill(id int, data interface{}) error {
	if err := db.Model(&Bill{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// 添加账单
func AddBill(data map[string]interface{}) error {
	if err := db.Create(&Bill{
		CategoryId: data["category_id"].(int),
		TypeId:     data["type_id"].(int),
		Amount:     data["amount"].(int),
		Remark:     data["remark"].(string),
	}).Error; err != nil {
		return err
	}

	return nil
}

// 删除订单
func DeleteBill(id int) bool {
	db.Where("id = ?", id).Delete(Bill{})

	return true
}

func (bill *Bill) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (bill *Bill) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
