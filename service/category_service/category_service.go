package category_service

import "github.com/EDDYCJY/go-gin-example/models"

type Category struct {
	Id 			int
	Name 		string
	CreatedBy   int
	Image       string
	CType       string
	State      int

	PageNum  int
	PageSize int
}

// 是否存在相同的名字
func (category *Category) ExistByName() (bool, error) {
	return models.ExistCategoryName(category.Name)
}

func (category *Category) Created() error {
	categoryParam := map[string]interface{}{
		"name":  category.Name,
		"type":  category.CType,
		"createdBy": category.CreatedBy,
		"image": category.Image,
	}
	return models.CreateCategory(categoryParam)
}

func (category *Category) GetCategoriesByType() {

}

//func (category *Category) GetAll() ([]models.Category, error) {
//	var (
//		//categoies,
//	)
//}