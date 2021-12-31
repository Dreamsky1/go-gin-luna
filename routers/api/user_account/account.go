package user_account

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/service/user_account_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type accountParam struct {
	UserId 	int 	`form:"user_id" valid:"Required"`
	Budget  int  	`form:"budget"`
	Integral int    `form:"integral"`
}
// 创建
func Put(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form accountParam
	)

	// 判断参数是否合规
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userAccountService := user_account_service.UserAccount{
		UserId: form.UserId,
		Budget: form.Budget,
		Integral: form.Integral,
	}

	err := userAccountService.Created()
	if err != nil {
		appG.Response(http.StatusInternalServerError, 5555, "created_bill:fail")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
