// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Sale is the golang structure of table sale for DAO operations like Where/Data.
type Sale struct {
	g.Meta           `orm:"table:sale, do:true"`
	Id               interface{} // 销售id
	SaleNo           interface{} // 销售单号
	SaleDate         *gtime.Time // 销售日期
	SaleType         interface{} // 销售类型 0:普通销售 1:其他出库
	SaleTypeName     interface{} // 销售类型名称
	CustomerId       interface{} // 客户id
	WarehouseId      interface{} // 仓库id
	BankId           interface{} // 银行id
	ReceivableAmount interface{} // 应收金额
	ReceivedAmount   interface{} // 实收金额
	SaleProfit       interface{} // 销售利润
	ArrearsAmount    interface{} // 欠款金额
	CreateUserId     interface{} // 创建人id
	UpdateUserId     interface{} // 修改人id
	AuditUserId      interface{} // 审核人id
	AuditStatus      interface{} // 销售状态 0:未审核 1:已审核
	Remark           interface{} // 备注
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	AuditAt          *gtime.Time // 审核时间
}
