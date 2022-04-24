package models

import "github.com/jinzhu/gorm"

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Signature string  `json:"signature"`
	Integral  int `json:"integral"`
}

// 检查是否有该用户
func CheckUser(username, password string) (error, User) {
	var user User
	err := db.Select("id").Where(User{
		Username: username,
		Password: password,
	}).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err, user
	}

	if user.ID > 0 {
		return nil, user
	}

	return err, user
}
// 注册
func RegisterUser(username, password string)  error {
	err := db.Create(&User{
		Username: username,
		Password: password,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
