// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Customer is the golang structure for table customer.
type Customer struct {
	Id             uint        `json:"id"             orm:"id"              description:"客户id"`
	Name           string      `json:"name"           orm:"name"            description:"客户名称"`
	Pinyin         string      `json:"pinyin"         orm:"pinyin"          description:"客户名称拼音"`
	Contact        string      `json:"contact"        orm:"contact"         description:"客户联系人"`
	Phone          string      `json:"phone"          orm:"phone"           description:"客户联系电话"`
	CompanyName    string      `json:"companyName"    orm:"company_name"    description:"客户公司名称"`
	Address        string      `json:"address"        orm:"address"         description:"客户地址"`
	TaxNumber      string      `json:"taxNumber"      orm:"tax_number"      description:"客户纳税号"`
	BankName       string      `json:"bankName"       orm:"bank_name"       description:"客户开户行"`
	BankAccount    string      `json:"bankAccount"    orm:"bank_account"    description:"客户银行账号"`
	BankHolder     string      `json:"bankHolder"     orm:"bank_holder"     description:"银行开户人"`
	OpeningBalance float64     `json:"openingBalance" orm:"opening_balance" description:"客户期初余额"`
	Status         string      `json:"status"         orm:"status"          description:"客户状态 0:禁用 1:启用"`
	Active         string      `json:"active"         orm:"active"          description:"客户默认 0:否 1:是"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`
}
