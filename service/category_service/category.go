package category_service

import "github.com/EDDYCJY/go-gin-example/models"

type Category struct {
	ID     int
	Name   string
	TypeId int
	State  int
	Image  string
}

func (category *Category) GetCategoryByTypeId(typeId int) ([]*models.Category, error) {
	return models.GetCategoriesByTypeId(typeId)
}
