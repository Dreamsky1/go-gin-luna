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
	"github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	// 注册
	r.PUT("/register", user.Register)
	// 登录
	r.POST("/login", user.Login)

	// 分类***********
	apiCategory := r.Group("/api/category")
	//apiCategory.Use(jwt.JWT())z暂时不用
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
	//apiBill.Use(jwt.JWT())
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

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
