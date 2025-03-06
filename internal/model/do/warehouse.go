// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Warehouse is the golang structure of table warehouse for DAO operations like Where/Data.
type Warehouse struct {
	g.Meta    `orm:"table:warehouse, do:true"`
	Id        interface{} // 仓库id
	Name      interface{} // 仓库名称
	Contact   interface{} // 仓库联系人
	Phone     interface{} // 仓库联系电话
	Status    interface{} // 仓库状态 0:禁用 1:启用
	Active    interface{} // 仓库默认 0:否 1:是
	Address   interface{} // 仓库地址
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
