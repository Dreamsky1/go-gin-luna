package routers

import (
	"github.com/EDDYCJY/go-gin-example/routers/api/bill"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/EDDYCJY/go-gin-example/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"github.com/EDDYCJY/go-gin-example/pkg/upload"
	"github.com/EDDYCJY/go-gin-example/routers/api"
	"github.com/EDDYCJY/go-gin-example/routers/api/category"
	"github.com/EDDYCJY/go-gin-example/routers/api/user"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	// 注册
	r.PUT("api/register", user.Register)
	// 登录
	r.POST("api/login", user.Login)

	// 分类***********
	apiCategory := r.Group("/api/category")
	apiCategory.Use(jwt.JWT())
	{
		//获取标签列表
		apiCategory.GET("/categories", category.GetCategories)
		//新建标签
		apiCategory.PUT("/category", category.AddCategory)
		//更新指定标签
		apiCategory.PUT("/categories/:id", category.EditCategory)
		//删除指定标签
		apiCategory.DELETE("/categories", category.DeleteCategory)

		// 一级分类
		// 获得一级分类并且携带了二级分类
		apiCategory.GET("/type/categories", category.GetAllTypeCategories)
		// 创建一级分类
		apiCategory.PUT("/type/category", category.AddType)
	}

	apiBill := r.Group("/api/bill")
	apiBill.Use(jwt.JWT())
	{
		// 获取多个账单
		apiBill.GET("/bills", bill.GetBills)
		// 获得单个账单
		apiBill.GET("/bill", bill.GetBill)
		// 更新指定的账单
		apiBill.POST("/bill", bill.EditBill)
		// 新增账单
		apiBill.PUT("/bill", bill.AddBill)
		//删除账单
		apiBill.DELETE("/bill", bill.DeleteBill)
	}

	return r
}
