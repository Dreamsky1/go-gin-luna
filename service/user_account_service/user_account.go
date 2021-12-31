package user_account_service

import "github.com/EDDYCJY/go-gin-example/models"

type UserAccount struct {
	Id 			int
	UserId 		int
	Budget		int
	Integral	int
}

func (userAccount *UserAccount) Exist() (bool, error) {
	return models.ExistUserAccount(userAccount.UserId)
}

func (userAccount *UserAccount) Created() error {
	param := map[string]interface{}{
		"user_id": userAccount.UserId,
		"budget": userAccount.Budget,
		"integral": userAccount.Integral,
	}

	return models.CreateUserAccount(param)
}

func (userAccount *UserAccount) UpdateIntegral() error {
	data := make(map[string]interface{})
	data["integral"] = userAccount.Integral
	return models.UpdateUserAccount(userAccount.UserId, data)
}

func (userAccount *UserAccount) UpdateBudget() error {
	data := make(map[string]interface{})
	data["budget"] = userAccount.Budget
	return models.UpdateUserAccount(userAccount.UserId, data)
}