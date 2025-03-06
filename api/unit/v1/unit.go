package v1

import (
	"goApp/api"
	"goApp/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta `path:"unit" method:"post" tags:"Unit" summary:"创建单位"`
	Name   string `json:"name"  v:"required|regex:^\\S{1,30}$#单位名称不能为空|单位名称1-30非空格字符" dc:"单位名称"`
}
type CreateRes struct{}
type UpdateReq struct {
	g.Meta `path:"unit/{id}" method:"put" tags:"Unit" summary:"更新单位"`
	Id     int64 `json:"id" in:"path" v:"required|min:1#单位id不能为空|单位id最小{min}" dc:"单位id"`
}
type UpdateRes struct{}
type DeleteReq struct {
	g.Meta `path:"unit/{id}" method:"delete" tags:"Unit" summary:"删除单位"`
	Id     int64 `json:"id" in:"path" v:"required|min:1#单位id不能为空|单位id最小{min}" dc:"单位id"`
}
type DeleteRes struct{}
type GetOneReq struct {
	g.Meta `path:"unit/{id}" method:"get" tags:"Unit" summary:"获取单位"`
	Id     int64 `json:"id" in:"path" v:"required|min:1#单位id不能为空|单位id最小{min}" dc:"单位id"`
}
type GetOneRes struct {
	*entity.Unit
}
type GetListReq struct {
	g.Meta `path:"unit/list" method:"get" tags:"Unit" summary:"获取单位列表"`
	*api.PageListReq
}
type GetListRes struct {
	*api.PageListRes
	List []*entity.Unit `json:"list"`
}
