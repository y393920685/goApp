package model

import (
	"goApp/internal/model/entity"
)

type UnitCreateInput struct {
	Name string
}
type UnitUpdateInput struct {
	Name string
}
type UnitGetListInput struct {
	PageNum  int
	PageSize int
	Search   string
}
type UnitGetListOutput struct {
	List  []*entity.Unit
	Total int64
}
type UnitGetOneOutput struct {
	*entity.Unit
}
