package models

import (
	"github.com/EDDYCJY/go-gin-example/service/user_account_service"
	"github.com/jinzhu/gorm"
	"time"
)

type UserAccount struct {
	Id  		int 	`json:"id"`
	UserId   	int  	`json:"user_id"`
	Budget		int		`json:"budget"`
	Integral 	int		`json:"integral"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt	time.Time `json:"update_at"`
}

// 创建账户
func CreateUserAccount(param map[string]interface{}) error {
	userAccount := UserAccount{
		UserId: param["user_id"].(int),
		Budget: param["budget"].(int),
		Integral: param["integral"].(int),
		CreatedAt: time.Now(),
		UpdateAt: time.Now(),
	}

	if err := db.Create(&userAccount).Error; err != nil {
		return err
	}

	return nil
}

// 判断是否已经存在
func ExistUserAccount(userId int) (bool, error) {
	var userAccount UserAccount
	err := db.Select("id").Where("user_id = ?", userId).First(&userAccount).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return false, err
	}

	if userAccount.Id > 0 {
		return true, nil
	}
	return false, nil
}

func UpdateUserAccount(userId int, data interface{}) error {
	if err := db.Model(&UserAccount{}).Where("user_id = ?", user_account_service.UserAccount{}).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
