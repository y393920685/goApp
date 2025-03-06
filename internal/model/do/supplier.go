// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Supplier is the golang structure of table supplier for DAO operations like Where/Data.
type Supplier struct {
	g.Meta         `orm:"table:supplier, do:true"`
	Id             interface{} // 供货商id
	Name           interface{} // 供货商名称
	Pinyin         interface{} // 供货商名称拼音
	Contact        interface{} // 供货商联系人
	Phone          interface{} // 供货商联系电话
	CompanyName    interface{} // 供货商公司名称
	Address        interface{} // 供货商地址
	TaxNumber      interface{} // 供货商纳税号
	BankName       interface{} // 供货商开户行
	BankAccount    interface{} // 供货商银行账号
	BankHolder     interface{} // 银行开户人
	OpeningBalance interface{} // 供货商期初余额
	Status         interface{} // 供货商状态 0:禁用 1:启用
	Active         interface{} // 供货商默认 0:否 1:是
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
