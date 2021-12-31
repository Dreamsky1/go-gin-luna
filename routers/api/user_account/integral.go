package user_account

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/service/user_account_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type integralParam struct {
	UserId  int      `form:"user_id" valid:"Required;Min(1)"`
	Integral int    `form:"integral" valid:"Required"`
}
func PostIntegral(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form integralParam
	)

	// 判断参数是否合规
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userAccountService := user_account_service.UserAccount{
		UserId: form.UserId,
		Integral: form.Integral,
	}

	err := userAccountService.UpdateIntegral()

	if err != nil {
		appG.Response(http.StatusInternalServerError, 666, "created_bill:fail")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}