package category

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)
//修改文章
//func EditArticle(c *gin.Context) {
//	valid := validation.Validation{}
//
//	id := com.StrTo(c.Param("id")).MustInt()
//	tagId := com.StrTo(c.Query("tag_id")).MustInt()
//	title := c.Query("title")
//	desc := c.Query("desc")
//	content := c.Query("content")
//	modifiedBy := c.Query("modified_by")
//
//	var state int = -1
//	if arg := c.Query("state"); arg != "" {
//		state = com.StrTo(arg).MustInt()
//		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
//	}
//
//	valid.Min(id, 1, "id").Message("ID必须大于0")
//	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
//	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
//	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
//	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
//	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
//
//	code := e.INVALID_PARAMS
//	if ! valid.HasErrors() {
//		if models.ExistArticleByID(id) {
//			if models.ExistTagByID(tagId) {
//				data := make(map[string]interface {})
//				if tagId > 0 {
//					data["tag_id"] = tagId
//				}
//				if title != "" {
//					data["title"] = title
//				}
//				if desc != "" {
//					data["desc"] = desc
//				}
//				if content != "" {
//					data["content"] = content
//				}
//
//				data["modified_by"] = modifiedBy
//
//				models.EditArticle(id, data)
//				code = e.SUCCESS
//			} else {
//				code = e.ERROR_NOT_EXIST_TAG
//			}
//		} else {
//			code = e.ERROR_NOT_EXIST_ARTICLE
//		}
//	} else {
//		for _, err := range valid.Errors {
//			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code" : code,
//		"msg" : e.GetMsg(code),
//		"data" : make(map[string]string),
//	})
//}

//获取多个分类
func GetCategories(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetCategories(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetCategoryTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

//新增分类
func AddCategory(c *gin.Context) {
	name := c.Query("name")
	image := c.Query("image")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	typeId := com.StrTo(c.DefaultQuery("type_id", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		models.AddCategory(name, state, typeId, image)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

//修改分类
func EditCategory(c *gin.Context) {
	// 暂时不做
}

//删除分类
func DeleteCategory(c *gin.Context) {
	//fmt.Print("蔬果来的formdata", ("id"))
	name := c.PostForm("name")
	fmt.Print("name名字***、", name)
	// 暂时不做
	id := com.StrTo(c.Query("id")).MustInt()

	fmt.Print("输出这个id", id)

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistCategoryByID(id) {
			models.DeleteCategory(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}