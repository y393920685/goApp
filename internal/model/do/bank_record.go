// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BankRecord is the golang structure of table bank_record for DAO operations like Where/Data.
type BankRecord struct {
	g.Meta    `orm:"table:bank_record, do:true"`
	Id        interface{} // 银行记录id
	BankId    interface{} // 银行id
	Type      interface{} // 记录类型 CG:采购入库 XS:销售出库
	TypeName  interface{} // 记录类型名称
	BillNo    interface{} // 单据号
	Amount    interface{} // 交易金额
	Balance   interface{} // 交易后余额
	UserId    interface{} // 创建人id
	CreatedAt *gtime.Time // 创建时间
}
