// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Customer is the golang structure of table customer for DAO operations like Where/Data.
type Customer struct {
	g.Meta         `orm:"table:customer, do:true"`
	Id             interface{} // 客户id
	Name           interface{} // 客户名称
	Pinyin         interface{} // 客户名称拼音
	Contact        interface{} // 客户联系人
	Phone          interface{} // 客户联系电话
	CompanyName    interface{} // 客户公司名称
	Address        interface{} // 客户地址
	TaxNumber      interface{} // 客户纳税号
	BankName       interface{} // 客户开户行
	BankAccount    interface{} // 客户银行账号
	BankHolder     interface{} // 银行开户人
	OpeningBalance interface{} // 客户期初余额
	Status         interface{} // 客户状态 0:禁用 1:启用
	Active         interface{} // 客户默认 0:否 1:是
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
