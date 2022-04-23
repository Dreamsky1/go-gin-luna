package category

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

//获取多个分类
func GetCategories(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{}) // 查询参数
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	//maps["name"] = "李白"
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetCategories(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetCategoryTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type AddCategoryForm struct {
	Name   string `form:"name" valid:"Required;MaxSize(100)"`
	Image  string `form:"image" valid:"Required;MaxSize(100)"`
	TypeId int    `form:"type_id" valid:"Required;Range(0,100)"`
	State  int    `form:"state" valid:"Range(0,1)"`
}

//新增分类
func AddCategory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddCategoryForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	code := e.INVALID_PARAMS

	err := models.AddCategory(form.Name, form.State, form.TypeId, form.Image)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}
	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//修改分类
func EditCategory(c *gin.Context) {
	// 暂时不做
}

//删除分类
func DeleteCategory(c *gin.Context) {
	//fmt.Print("蔬果来的formdata", ("id"))
	name := c.PostForm("name")
	fmt.Print("name名字***、", name)
	// 暂时不做
	id := com.StrTo(c.Query("id")).MustInt()

	fmt.Print("输出这个id", id)

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistCategoryByID(id) {
			models.DeleteCategory(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
