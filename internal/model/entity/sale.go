// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Sale is the golang structure for table sale.
type Sale struct {
	Id               uint        `json:"id"               orm:"id"                description:"销售id"`
	SaleNo           string      `json:"saleNo"           orm:"sale_no"           description:"销售单号"`
	SaleDate         *gtime.Time `json:"saleDate"         orm:"sale_date"         description:"销售日期"`
	SaleType         string      `json:"saleType"         orm:"sale_type"         description:"销售类型 0:普通销售 1:其他出库"`
	SaleTypeName     string      `json:"saleTypeName"     orm:"sale_type_name"    description:"销售类型名称"`
	CustomerId       uint        `json:"customerId"       orm:"customer_id"       description:"客户id"`
	WarehouseId      uint        `json:"warehouseId"      orm:"warehouse_id"      description:"仓库id"`
	BankId           uint        `json:"bankId"           orm:"bank_id"           description:"银行id"`
	ReceivableAmount float64     `json:"receivableAmount" orm:"receivable_amount" description:"应收金额"`
	ReceivedAmount   float64     `json:"receivedAmount"   orm:"received_amount"   description:"实收金额"`
	SaleProfit       float64     `json:"saleProfit"       orm:"sale_profit"       description:"销售利润"`
	ArrearsAmount    float64     `json:"arrearsAmount"    orm:"arrears_amount"    description:"欠款金额"`
	CreateUserId     uint        `json:"createUserId"     orm:"create_user_id"    description:"创建人id"`
	UpdateUserId     uint        `json:"updateUserId"     orm:"update_user_id"    description:"修改人id"`
	AuditUserId      uint        `json:"auditUserId"      orm:"audit_user_id"     description:"审核人id"`
	AuditStatus      string      `json:"auditStatus"      orm:"audit_status"      description:"销售状态 0:未审核 1:已审核"`
	Remark           string      `json:"remark"           orm:"remark"            description:"备注"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"        description:"更新时间"`
	AuditAt          *gtime.Time `json:"auditAt"          orm:"audit_at"          description:"审核时间"`
}
