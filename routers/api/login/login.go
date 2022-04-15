package login

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserForm struct {
	Username      string    `form:"username" valid:"Required;MaxSize(100)"`
	Password      string    `form:"password" valid:"Required;MaxSize(100)"`
}

func Login(ctx *gin.Context)  {
	var (
		appG = app.Gin{C: ctx}
		form UserForm
	)

	httpCode, errCode := app.BindAndValid(ctx, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	username := form.Username
	password  := form.Password


	userService := user_service.User{
		Username: form.Username,
		Password: form.Password,
	}
	isExist, err := userService.CheckUserByUsername()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, "检查user是否存在失败")
		return
	}

	if !isExist {
		err := userService.CreateUser()
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, "创建用户失败")
			return
		}
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, "生成用户token失败")
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
		"username": username,
	})
}
