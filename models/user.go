package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Model

	Username  string `json:"username"`
	Password  string `json:"password"`
	Signature string `json:"signature"`
	Integral  int    `json:"integral"`
}

// 检查是否有该用户
func CheckUser(username, password string) (error, User, bool) {
	var user User
	err := db.Where(User{
		Username: username,
		Password: password,
	}).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err, user, false
	}

	if user.ID > 0 {
		return nil, user, true
	}

	return err, user, false
}

// 注册
func RegisterUser(username, password string) error {
	if err := db.Create(&User{
		Username:  username,
		Password:  password,
		Signature: "1",
		Integral:  1,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
