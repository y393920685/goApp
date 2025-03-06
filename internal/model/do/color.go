// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Color is the golang structure of table color for DAO operations like Where/Data.
type Color struct {
	g.Meta    `orm:"table:color, do:true"`
	Id        interface{} // 颜色id
	Name      interface{} // 颜色名称
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
