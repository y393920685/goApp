// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Bank is the golang structure for table bank.
type Bank struct {
	Id             uint        `json:"id"             orm:"id"              description:"银行id"`
	Name           string      `json:"name"           orm:"name"            description:"银行名称"`
	Holder         string      `json:"holder"         orm:"holder"          description:"开户人"`
	Account        string      `json:"account"        orm:"account"         description:"银行账号"`
	Address        string      `json:"address"        orm:"address"         description:"开户银行地址"`
	Status         string      `json:"status"         orm:"status"          description:"银行状态 0:禁用 1:启用"`
	Active         string      `json:"active"         orm:"active"          description:"银行默认 0:否 1:是"`
	OpeningBalance float64     `json:"openingBalance" orm:"opening_balance" description:"期初余额"`
	Balance        float64     `json:"balance"        orm:"balance"         description:"余额"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`
}
