// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BankRecord is the golang structure for table bank_record.
type BankRecord struct {
	Id        uint        `json:"id"        orm:"id"         description:"银行记录id"`
	BankId    uint        `json:"bankId"    orm:"bank_id"    description:"银行id"`
	Type      string      `json:"type"      orm:"type"       description:"记录类型 CG:采购入库 XS:销售出库"`
	TypeName  string      `json:"typeName"  orm:"type_name"  description:"记录类型名称"`
	BillNo    string      `json:"billNo"    orm:"bill_no"    description:"单据号"`
	Amount    float64     `json:"amount"    orm:"amount"     description:"交易金额"`
	Balance   float64     `json:"balance"   orm:"balance"    description:"交易后余额"`
	UserId    uint        `json:"userId"    orm:"user_id"    description:"创建人id"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
}
