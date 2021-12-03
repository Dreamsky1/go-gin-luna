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

// 定义接受数据的结构体
type UserParam struct {
	Name string `form:"name" valid:"Required;"`
	Avatar string `form:"avatar" valid:"Required;"`
	Phone  string `form:"phone" valid:"Required"`
	Gender string `form:"gender" valid:"Required"`
	Code   string `form:"code" valid:"Required"`
	UnionId string `from:"unionid"`
	OpenId string `from:"openid"`
}


func LoginMobileUser(ctx *gin.Context) {
	var (
		appG = app.Gin{C: ctx}
		form UserParam
	)
	unionId := ctx.Query("unionid")
	openid := ctx.Query("openid")
	fmt.Println("书法和借口个", unionId, openid)

	httpCode, errCode := app.BindAndValid(ctx, &form)
	fmt.Println("httpCode", httpCode, "errCode", errCode)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, "解析参数失败")
		return
	}
	// 1.这里要做个解析code， 获取openid和unid的操作，先省略

	// 2.根据uid和opid去查看看，有没有
	userService := user_service.User{
		Name: form.Name,
		Avatar: form.Avatar,
		Phone: form.Phone,
		Code: form.Code,
		Gender: form.Gender,
		UnionId: unionId,
		OpenId: openid,
	}
	isExist, err := userService.CheckUser()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, "检查user是否存在失败")
		return
	}

	if !isExist {
		err := userService.CreateUserByUnionId()
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, "创建用户失败")
			return
		}
	}

	token, err := util.GenerateToken(form.UnionId, form.OpenId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, "生成用户token失败")
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})

}
