// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Bank is the golang structure of table bank for DAO operations like Where/Data.
type Bank struct {
	g.Meta         `orm:"table:bank, do:true"`
	Id             interface{} // 银行id
	Name           interface{} // 银行名称
	Holder         interface{} // 开户人
	Account        interface{} // 银行账号
	Address        interface{} // 开户银行地址
	Status         interface{} // 银行状态 0:禁用 1:启用
	Active         interface{} // 银行默认 0:否 1:是
	OpeningBalance interface{} // 期初余额
	Balance        interface{} // 余额
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
