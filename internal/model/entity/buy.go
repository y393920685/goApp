// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Buy is the golang structure for table buy.
type Buy struct {
	Id               uint        `json:"id"               orm:"id"                description:"采购id"`
	BuyNo            string      `json:"buyNo"            orm:"buy_no"            description:"采购单号"`
	BuyDate          *gtime.Time `json:"buyDate"          orm:"buy_date"          description:"采购日期"`
	BuyType          string      `json:"buyType"          orm:"buy_type"          description:"采购类型 0:普通采购 1:其他入库"`
	BuyTypeName      string      `json:"buyTypeName"      orm:"buy_type_name"     description:"采购类型名称"`
	CustomerId       uint        `json:"customerId"       orm:"customer_id"       description:"供货商id"`
	WarehouseId      uint        `json:"warehouseId"      orm:"warehouse_id"      description:"仓库id"`
	BankId           uint        `json:"bankId"           orm:"bank_id"           description:"银行id"`
	ReceivableAmount float64     `json:"receivableAmount" orm:"receivable_amount" description:"应付金额"`
	ReceivedAmount   float64     `json:"receivedAmount"   orm:"received_amount"   description:"实付金额"`
	ArrearsAmount    float64     `json:"arrearsAmount"    orm:"arrears_amount"    description:"欠款金额"`
	CreateUserId     uint        `json:"createUserId"     orm:"create_user_id"    description:"创建人id"`
	UpdateUserId     uint        `json:"updateUserId"     orm:"update_user_id"    description:"修改人id"`
	AuditUserId      uint        `json:"auditUserId"      orm:"audit_user_id"     description:"审核人id"`
	AuditStatus      string      `json:"auditStatus"      orm:"audit_status"      description:"采购状态 0:未审核 1:已审核"`
	Remark           string      `json:"remark"           orm:"remark"            description:"备注"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"        description:"更新时间"`
	AuditAt          *gtime.Time `json:"auditAt"          orm:"audit_at"          description:"审核时间"`
}
