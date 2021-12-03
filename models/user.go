package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID 		int `json:"id"`
	UnionId  string 	`json:"union_id"`
	OpenId   string    `json:"open_id"`
	Phone    string		`json:"phone"`
	Code     string    `json:"code"`
	Name     string 	`json:"name"`
	Avatar   string 	`json:"avatar"`
	Gender   string    `json:"gender"`
	Username string 	`json:"username"`
	Password string 	`json:"password"`
	//CratedAt time.Time `json:"crated_at"`
}
// check
func CheckWeixinUser(unionid, openid string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{UnionId: unionid, OpenId: openid}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("check_err", err)
		return false, err
	}
	if user.ID > 0 {
		fmt.Println("应该是存在的")
		return true, nil
	}

	return false, nil
}


// create user
func CreateUserByUnionId(userParam map[string]interface{}) error {
	user := User {
		UnionId: userParam["unionid"].(string),
		OpenId: userParam["openid"].(string),
		Name: userParam["name"].(string),
		Phone: userParam["phone"].(string),
		Code: userParam["code"].(string),
		Avatar: userParam["avatar"].(string),
		Gender: userParam["gender"].(string),
		Username: userParam["username"].(string),
		Password: userParam["password"].(string),
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}