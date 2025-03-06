package v1

import (
	"goApp/api"
	"goApp/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta `path:"/brand" method:"post" tags:"Brand" summary:"创建品牌"`
	Name   string `json:"name"  v:"required|regex:^\\S{1,30}$#品牌名称不能为空|品牌名称1-30非空格字符" dc:"品牌名称"`
}
type CreateRes struct{}
type UpdateReq struct {
	g.Meta `path:"/brand/{id}" method:"put" tags:"Brand" summary:"更新品牌"`
	Id     int64  `json:"id" in:"path" v:"required|min:1#品牌id不能为空|品牌id最小{min}" dc:"品牌id"`
	Name   string `json:"name"  v:"required|regex:^\\S{1,30}$#品牌名称不能为空|品牌名称1-30非空格字符" dc:"品牌名称"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/brand/{id}" method:"delete" tags:"Brand" summary:"删除品牌"`
	Id     int64 `json:"id" in:"path" v:"required|min:1#品牌id不能为空|品牌id最小{min}" dc:"品牌id"`
}
type DeleteRes struct{}

type GetOneReq struct {
	g.Meta `path:"/brand/{id}" method:"get" tags:"Brand" summary:"获取品牌"`
	Id     int64 `json:"id" in:"path" v:"required|min:1#品牌id不能为空|品牌id最小{min}" dc:"品牌id"`
}
type GetOneRes struct {
	*entity.Brand
}
type GetListReq struct {
	g.Meta `path:"/brand/list" method:"get" tags:"Brand" summary:"获取品牌列表"`
	*api.PageListReq
}
type GetListRes struct {
	*api.PageListRes
	List []*entity.Brand `json:"list"`
}
