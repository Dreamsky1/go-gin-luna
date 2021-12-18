package bill_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"time"
)

type Bill struct {
	Id     		int
	UserId 		int
	CategoryId  int
	Amount      int
	Remark      string
	Image       string
	CreatedAt   string
}

func (bill *Bill) Created() error {
	billParam := map[string]interface{} {
		"user_id":  bill.UserId,
		"category_id": bill.CategoryId,
		"amount": bill.Amount,
		"remark": bill.Remark,
		"image": bill.Image,
		"created_at": bill.CreatedAt,
	}

	return models.CreateBill(billParam)
}

func (bill *Bill) Update() error {
	data := make(map[string]interface{})
	data["amount"] = bill.Amount
	data["remark"] = bill.Remark
	data["image"] = bill.Image
	data["created_at"] = util.ParseTime(bill.CreatedAt)
	data["update_at"] = time.Now()
	return models.UpdateBill(bill.Id, data)
}

func (bill *Bill) Delete() error {
	return models.DeleteBill(bill.Id)
}

func (bill *Bill) ExistById() (bool, error)  {
	return models.ExistBillById(bill.Id)
}