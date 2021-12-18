package bill

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/service/bill_service"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type billForm struct {
	UserId   int  `form:"user_id" valid:"Required"`
	CategoryId int `form:"category_id" valid:"Required"`
	Amount    int  `form:"amount" valid:"Required"`
	CreatedAt string  `form:"created_at" valid:"Required"`
	Remark    string  `form:"remark"`
	Image    string  `form:"image"`
}

// 创建账单
func Put(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form billForm
	)

	// 判断参数是否合规
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	billService := bill_service.Bill{
		UserId: form.UserId,
		CategoryId: form.CategoryId,
		Amount: form.Amount,
		Image: form.Image,
		Remark: form.Remark,
		CreatedAt: form.CreatedAt,
	}

	err := billService.Created()
	if err != nil {
		appG.Response(http.StatusInternalServerError, 5555, "created_bill:fail")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type PostBillForm struct {
	Id       int   `form:"id" valid:"Required;Min(1)"`
	UserId   int  `form:"user_id" valid:"Required"`
	CategoryId int `form:"category_id" valid:"Required"`
	Amount    int  `form:"amount" valid:"Required"`
	CreatedAt string  `form:"created_at" valid:"Required"`
	Remark    string  `form:"remark"`
	Image    string  `form:"image"`
}

// 修改账单
func Post(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form = PostBillForm{Id: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	billService := bill_service.Bill{
		Id: form.Id,
		UserId: form.UserId,
		CategoryId: form.CategoryId,
		Amount: form.Amount,
		Image: form.Image,
		Remark: form.Remark,
		CreatedAt: form.CreatedAt,
	}

	exists, err := billService.ExistById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, 11111, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, 11111, "bill_id_is_not_exist")
		return
	}

	err = billService.Update()
	if err != nil {
		appG.Response(http.StatusInternalServerError, 1111, "update_bill:fail")
		return
	}
}

// 删除账单
func Delete(c *gin.Context)  {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	billService := bill_service.Bill{Id: id}

	exists, err := billService.ExistById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, 11111, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, 11111, "bill_id_is_not_exist")
		return
	}

	if err := billService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, "delete_bill_fail")
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}