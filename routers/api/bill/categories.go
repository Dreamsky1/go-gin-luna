package bill

import (
	"github.com/gin-gonic/gin"
)

type categoryFrom struct {
	CType string `form:"c_type"`
}

func GetCategories(c *gin.Context)  {

	//appG := app.Gin{C: c}
	//CType := ""
	//if arg := c.Query("c_type"); arg != "" {
	//	CType = c.Query("c_type")
	//}
	//
	//categoryService := category_service.Category{
	//	CType: CType,
	//}

}