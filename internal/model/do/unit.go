// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Unit is the golang structure of table unit for DAO operations like Where/Data.
type Unit struct {
	g.Meta    `orm:"table:unit, do:true"`
	Id        interface{} // 单位id
	Name      interface{} // 单位名称
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
