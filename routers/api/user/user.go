package user

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userFrom struct {
	Username string `form:"username" valid:"Required; MaxSize(50)"`
	Password string `form:"password" valid:"Required; MaxSize(50)"`
}

func Register(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form userFrom
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	err, _, exist := models.CheckUser(form.Username, form.Password)
	// 存在用户
	if exist {
		appG.Response(http.StatusUnauthorized, e.ERROR_EXIST_USER, nil)
		return
	}

	// 注册
	err = models.RegisterUser(form.Username, form.Password)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_CREATE_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

type loginForm struct {
	Username string `form:"username" valid:"Required; MaxSize(50)"`
	Password string `form:"password" valid:"Required; MaxSize(50)"`
}

func Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form loginForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	data := make(map[string]interface{})
	err, user, exist := models.CheckUser(form.Username, form.Password)

	// 不存在该用户
	if !exist {
		fmt.Print("申购金额", exist)
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": e.ERROP_NOT_EXIST_USER,
			"msg":  e.GetMsg(e.ERROP_NOT_EXIST_USER),
			"data": make(map[string]interface{}),
		})
		//appG.Response(http.StatusUnauthorized, e.ERROP_NOT_EXIST_USER, nil)
		return
	}

	if err != nil {
		appG.Response(httpCode, errCode, nil)
		return
	}

	token, err := util.GenerateToken(form.Username, form.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	data["user"] = user
	data["token"] = token

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
