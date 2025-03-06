// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Category is the golang structure of table category for DAO operations like Where/Data.
type Category struct {
	g.Meta    `orm:"table:category, do:true"`
	Id        interface{} // 产品类别id
	Name      interface{} // 产品类别名称
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
