// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Supplier is the golang structure for table supplier.
type Supplier struct {
	Id             uint        `json:"id"             orm:"id"              description:"供货商id"`
	Name           string      `json:"name"           orm:"name"            description:"供货商名称"`
	Pinyin         string      `json:"pinyin"         orm:"pinyin"          description:"供货商名称拼音"`
	Contact        string      `json:"contact"        orm:"contact"         description:"供货商联系人"`
	Phone          string      `json:"phone"          orm:"phone"           description:"供货商联系电话"`
	CompanyName    string      `json:"companyName"    orm:"company_name"    description:"供货商公司名称"`
	Address        string      `json:"address"        orm:"address"         description:"供货商地址"`
	TaxNumber      string      `json:"taxNumber"      orm:"tax_number"      description:"供货商纳税号"`
	BankName       string      `json:"bankName"       orm:"bank_name"       description:"供货商开户行"`
	BankAccount    string      `json:"bankAccount"    orm:"bank_account"    description:"供货商银行账号"`
	BankHolder     string      `json:"bankHolder"     orm:"bank_holder"     description:"银行开户人"`
	OpeningBalance float64     `json:"openingBalance" orm:"opening_balance" description:"供货商期初余额"`
	Status         string      `json:"status"         orm:"status"          description:"供货商状态 0:禁用 1:启用"`
	Active         string      `json:"active"         orm:"active"          description:"供货商默认 0:否 1:是"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`
}
