package v1

import (
	"goApp/api"
	"goApp/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta `path:"/color" method:"post" tags:"Color" summary:"创建颜色"`
	Name   string `json:"name"  v:"required|regex:^\\S{1,30}$#颜色名称不能为空|颜色名称1-30非空格字符" dc:"颜色名称"`
}
type CreateRes struct{}
type UpdateReq struct {
	g.Meta `path:"/color/{id}" method:"put" tags:"Color" summary:"更新颜色"`
	Id     int64  `json:"id" in:"path" v:"required|min:1#颜色id不能为空|颜色id最小{min}" dc:"颜色id"`
	Name   string `json:"name"  v:"required|regex:^\\S{1,30}$#颜色名称不能为空|颜色名称1-30非空格字符" dc:"颜色名称"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/color/{id}" method:"delete" tags:"Color" summary:"删除颜色"`
	Id     int64 `json:"id" in:"path" v:"required|min:1#颜色id不能为空|颜色id最小{min}" dc:"颜色id"`
}
type DeleteRes struct{}

type GetOneReq struct {
	g.Meta `path:"/color/{id}" method:"get" tags:"Color" summary:"获取颜色"`
	Id     int64 `json:"id" in:"path" v:"required|min:1#颜色id不能为空|颜色id最小{min}" dc:"颜色id"`
}
type GetOneRes struct {
	*entity.Color
}
type GetListReq struct {
	g.Meta `path:"/color/list" method:"get" tags:"Color" summary:"获取颜色列表"`
	*api.PageListReq
}
type GetListRes struct {
	*api.PageListRes
	List []*entity.Color `json:"list"`
}
