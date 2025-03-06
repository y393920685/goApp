package model

import (
	"goApp/internal/model/entity"
)

type ColorCreateInput struct {
	Name string
}
type ColorUpdateInput struct {
	Name string
}
type ColorGetListInput struct {
	PageNum  int
	PageSize int
	Search   string
}
type ColorGetListOutput struct {
	List  []*entity.Color
	Total int64
}
type ColorGetOneOutput struct {
	*entity.Color
}
