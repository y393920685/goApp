package model

import (
	"goApp/internal/model/entity"
)

type BrandCreateInput struct {
	Name string
}
type BrandUpdateInput struct {
	Name string
}
type BrandGetListInput struct {
	PageNum  int
	PageSize int
	Search   string
}
type BrandGetListOutput struct {
	List  []*entity.Brand
	Total int
}
type BrandGetOneOutput struct {
	*entity.Brand
}
