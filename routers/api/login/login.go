package login

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context)  {
	fmt.Print("输出进来了这个")
	var appG = app.Gin{C: ctx}
	username := ctx.Query("username")
	password  := ctx.Query("password")

	userService := user_service.User{
		Username: username,
		Password: password,
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
	})
}
