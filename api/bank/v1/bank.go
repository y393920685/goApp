package v1

import (
	"goApp/api"
	"goApp/internal/consts"
	"goApp/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type BaseType struct {
	Name    string        `json:"name"  v:"required|regex:^\\S{1,30}$#银行名称不能为空|银行名称1-30非空格字符" dc:"银行名称"`
	Holder  string        `json:"holder"  v:"regex:^\\S{1,30}$#开户人1-30非空格字符" dc:"开户人"`
	Account string        `json:"account"  v:"regex:^\\S{1,80}$#银行账号1-80非空格字符" dc:"银行账号"`
	Address string        `json:"address"  v:"regex:^\\S{1,80}$#开户银行地址1-80非空格字符" dc:"开户银行地址"`
	Status  consts.Status `json:"status" d:"1" v:"enums#请选择正确的状态" dc:"银行状态 0:禁用 1:启用"`
	Active  consts.Active `json:"active" d:"0" v:"enums#请选择正确的状态" dc:"银行默认 0:否 1:是"`
}
type CreateReq struct {
	g.Meta `path:"/bank" method:"post" tags:"Bank" summary:"新建银行"`
	BaseType
}
type CreateRes struct{}
type Update struct {
	g.Meta `path:"/bank/{id}" method:"put" tags:"Bank" summary:"更新银行"`
	Id     int64 `json:"id" in:"path" v:"required|min:1#银行id不能为空|银行id最小{min}" dc:"银行id"`
	BaseType
}
type UpdateRes struct{}
type DeleteReq struct {
	g.Meta `path:"/bank/{id}" method:"delete" tags:"Bank" summary:"删除银行"`
	Id     int64 `json:"id" in:"path" v:"required|min:1#银行id不能为空|银行id最小{min}" dc:"银行id"`
}
type DeleteRes struct{}
type GetOneReq struct {
	g.Meta `path:"/bank/{id}" method:"get" tags:"Bank" summary:"获取银行"`
	Id     int64 `json:"id" in:"path" v:"required|min:1#银行id不能为空|银行id最小{min}" dc:"银行id"`
}
type GetOneRes struct {
	*entity.Bank
}
type GetListReq struct {
	g.Meta `path:"/bank/list" method:"get" tags:"Bank" summary:"获取银行列表"`
	*api.PageListReq
}
type GetListRes struct {
	*api.PageListRes
	List []*entity.Bank `json:"list"`
}
