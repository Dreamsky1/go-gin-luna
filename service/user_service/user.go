package user_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
)

type User struct {
	ID 			int
	Name 		string
	Phone 		string
	Avatar 		string
	UnionId 	string
	OpenId		string
	Code 		string
	Gender		string
	Username 	string
	Password	string
}

func (user *User) CheckUser() (bool, error) {
	return models.CheckWeixinUser(user.UnionId, user.OpenId)
}
func (user *User) CheckUserByUsername() (bool, error) {
	return models.CheckUsername(user.Username)
}

func (user *User) CreateUser() error {
	userParam := map[string]interface{}{
		"username": user.Username,
		"password": user.Password,
	}

	return models.CreateUser(userParam)
}

func (user *User) CreateUserByUnionId() error {
	userParam := map[string]interface{}{
		"name":		user.Name,
		"unionid": user.UnionId,
		"openid": user.OpenId,
		"avatar": user.Avatar,
		"phone": user.Phone,
		"code": user.Code,
		"gender": user.Gender,
		"username": user.Username,
		"password": user.Password,
	}
	// 这里在用户登录之后，就给用户去创建一个userAccount
	return models.CreateUserByUnionId(userParam)
}