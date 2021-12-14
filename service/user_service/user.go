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
	return models.CreateUserByUnionId(userParam)
}