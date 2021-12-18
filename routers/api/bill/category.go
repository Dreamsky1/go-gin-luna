package bill

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/service/category_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type categoryForm struct {
	Name    string  `form:"name" valid:"Required"`
	CType   string `form:"c_type" valid:"Required"`
	Image    string  `form:"image" valid:"Required"`
	CreatedBy  int `form:"created_by"`
}

func CreatedCategory(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form categoryForm
	)

	// 判断参数是否合规
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	categoryService := category_service.Category{
		Name: form.Name,
		Image: form.Image,
		CType: form.CType,
		CreatedBy: form.CreatedBy,
	}
	exists, err := categoryService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, 4444, "category_name:exist")
		return
	}

	if exists {
		appG.Response(http.StatusInternalServerError, 4444, "category_name:exist")
		return
	}

	err = categoryService.Created()
	if err != nil {
		appG.Response(http.StatusInternalServerError, 4444, "created_category:fail")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
