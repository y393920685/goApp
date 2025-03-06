// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Buy is the golang structure of table buy for DAO operations like Where/Data.
type Buy struct {
	g.Meta           `orm:"table:buy, do:true"`
	Id               interface{} // 采购id
	BuyNo            interface{} // 采购单号
	BuyDate          *gtime.Time // 采购日期
	BuyType          interface{} // 采购类型 0:普通采购 1:其他入库
	BuyTypeName      interface{} // 采购类型名称
	CustomerId       interface{} // 供货商id
	WarehouseId      interface{} // 仓库id
	BankId           interface{} // 银行id
	ReceivableAmount interface{} // 应付金额
	ReceivedAmount   interface{} // 实付金额
	ArrearsAmount    interface{} // 欠款金额
	CreateUserId     interface{} // 创建人id
	UpdateUserId     interface{} // 修改人id
	AuditUserId      interface{} // 审核人id
	AuditStatus      interface{} // 采购状态 0:未审核 1:已审核
	Remark           interface{} // 备注
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	AuditAt          *gtime.Time // 审核时间
}
