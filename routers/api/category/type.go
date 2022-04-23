package category

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获得所有的分类
func GetAllTypeCategories(c *gin.Context) {

	appG := app.Gin{C: c}

	maps := make(map[string]interface{}) // 就是查询参数
	data := make(map[string]interface{})

	code := e.SUCCESS

	lists, err := models.GetTypeCategories(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"], err = models.GetAllTypeCategoryTotal(maps)
	data["lists"] = lists

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CATEGORY_ONE_FAIL, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type AddTypeFrom struct {
	Name string `form:"name" valid:"Required;MaxSize(100)"`
}

// 新增分类类型
func AddType(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddTypeFrom
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	code := e.INVALID_PARAMS
	// 操作数据库之后会修改成service的模式

	err := models.AddType(form.Name)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CATEGORY_ONE_FAIL, nil)
		return
	}
	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
