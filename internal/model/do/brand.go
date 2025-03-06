// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Brand is the golang structure of table brand for DAO operations like Where/Data.
type Brand struct {
	g.Meta    `orm:"table:brand, do:true"`
	Id        interface{} // 品牌id
	Name      interface{} // 品牌名称
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
