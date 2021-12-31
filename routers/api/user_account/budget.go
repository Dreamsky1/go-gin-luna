package user_account

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/service/user_account_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type budGetParam struct {
	UserId  int      `form:"user_id" valid:"Required;Min(1)"`
	budget int    `form:"budget" valid:"Required"`
}

func PostBudget(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form budGetParam
	)

	// 判断参数是否合规
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userAccountService := user_account_service.UserAccount{
		UserId: form.UserId,
		Budget: form.budget,
	}

	err := userAccountService.UpdateBudget()

	if err != nil {
		appG.Response(http.StatusInternalServerError, 666, "created_bill:fail")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
