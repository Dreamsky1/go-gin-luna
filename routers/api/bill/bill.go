package bill

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"

	"github.com/unknwon/com"
)

//获取单个账单
func GetBill(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.INVALID_PARAMS

	var data interface{}
	if !valid.HasErrors() {
		if models.ExistBillByID(id) {
			data = models.GetBill(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//获取多个账单
func GetBills(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var categoryId int = -1
	if arg := c.Query("category_id"); arg != "" {
		categoryId = com.StrTo(arg).MustInt()
		maps["category_id"] = categoryId
	}

	if time1 := c.Query("time1"); time1 != "" {
		parseTime, _ := time.Parse("2006-01-02 15:04:05", time1)
		fmt.Print("测试一下**", parseTime.Unix())
	}
	if time2 := c.Query("time2"); time2 != "" {
		parseTime, _ := time.Parse("2006-01-02 15:04:05", time2)
		fmt.Print("测试一下**", parseTime)
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		bills := models.GetBills(util.GetPage(c), setting.AppSetting.PageSize, maps)
		data["lists"] = bills
		data["total"] = models.GetBillTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type AddBillForm struct {
	TypeId         int    `form:"type_id" valid:"Required;Min(1)"`
	CategoryId     int    `form:"category_id" valid:"Required;Min(1)"`
	AccountingDate string `form:"accounting_date" valid:"Required;MaxSize(65535)"`
	Amount         int    `form:"amount" valid:"Required;Min(1)"`
	Remark         string `form:"remark" valid:"Required;MaxSize(65535)"`
}

// 新增账单
func AddBill(c *gin.Context) {
	// 使用form去使用
	var (
		appG = app.Gin{C: c}
		form AddBillForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	code := e.INVALID_PARAMS
	parseTime, errs := time.Parse("2006-01-02 15:04:05", form.AccountingDate)
	if errs != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, "出粗哦了")
		return
	}

	data := make(map[string]interface{})
	data["category_id"] = form.CategoryId
	data["type_id"] = form.TypeId
	data["remark"] = form.Remark
	data["amount"] = form.Amount
	data["accounting_date"] = parseTime.Unix()
	err := models.AddBill(data)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errCode, nil)
		return
	}
	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

type EditBillForm struct {
	ID             int    `form:"id" valid:"Required;Min(1)"`
	TypeId         int    `form:"type_id" valid:"Required;Min(1)"`
	CategoryId     int    `form:"category_id" valid:"Required;Min(1)"`
	AccountingDate int    `form:"accounting_date" valid:"Required;Min(1)"`
	Amount         int    `form:"amount" valid:"Required;Min(1)"`
	Remark         string `form:"remark" valid:"Required;MaxSize(65535)"`
}

//修改账单
func EditBill(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form EditBillForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	exists := models.ExistBillByID(form.ID)

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	err := models.EditBill(form.ID, map[string]interface{}{
		"type_id":         form.TypeId,
		"amount":          form.Amount,
		"category_id":     form.CategoryId,
		"remark":          form.Remark,
		"accounting_date": form.AccountingDate,
	})
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

//删除账单
func DeleteBill(c *gin.Context) {
}
