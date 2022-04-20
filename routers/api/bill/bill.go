package bill

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/unknwon/com"
)

//获取单个账单
func GetBill(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.INVALID_PARAMS

	var data interface {}
	if ! valid.HasErrors() {
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
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

//获取多个账单
func GetBills(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var categoryId int = -1
	if arg :=  c.Query("category_id"); arg != "" {
		categoryId = com.StrTo(arg).MustInt()
		maps["category_id"] = categoryId
		valid.Min(categoryId, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetBills(util.GetPage(c), setting.AppSetting.PageSize, maps)
		data["total"] = models.GetBillTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

type AddBillForm struct {
	TypeId int  `form:"type_id" valid:"Required;Min(1)"`
	CategoryId int `form:"category_id" valid:"Required;Min(1)"`
	Amount int `form:"amount" valid:"Required;Min(1)"`
	Remark  string `form:"remark" valid:"Required;MaxSize(65535)"`
}

// 新增账单
func AddBill(c *gin.Context) {
	//categoryId := com.StrTo(c.Query("category_id")).MustInt()
	//typeId := com.StrTo(c.Query("type_id")).MustInt()
	//remark := c.Query("remark")
	//amount := com.StrTo(c.Query("amount")).MustInt()
	//
	//valid := validation.Validation{}
	//valid.Min(categoryId, 1, "category_id").Message("分类ID必须大于0")
	//valid.Min(typeId, 1, "type_id").Message("类型ID必须大于0")

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

	data := make(map[string]interface{})
	data["category_id"] = form.CategoryId
	data["type_id"] = form.TypeId
	data["remark"] = form.Remark
	data["amount"] = form.Amount
	err := models.AddBill(data)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errCode, nil)
		return
	}
	code = e.SUCCESS

	// 使用valid的方法
	//if !valid.HasErrors() {
	//	data := make(map[string]interface{})
	//	data["category_id"] = categoryId
	//	data["type_id"] = typeId
	//	data["remark"] = remark
	//	data["amount"] = amount
	//	err := models.AddBill(data)
	//	if err != nil {
	//		code = e.SUCCESS
	//	}
	//} else {
	//	for _, err := range valid.Errors {
	//		log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
	//	}
	//}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}

//修改账单
func EditBill(c *gin.Context) {
}

//删除账单
func DeleteBill(c *gin.Context) {
}
