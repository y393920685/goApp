// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BuyDao is the data access object for the table buy.
type BuyDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of the current DAO.
	columns BuyColumns // columns contains all the column names of Table for convenient usage.
}

// BuyColumns defines and stores column names for the table buy.
type BuyColumns struct {
	Id               string // 采购id
	BuyNo            string // 采购单号
	BuyDate          string // 采购日期
	BuyType          string // 采购类型 0:普通采购 1:其他入库
	BuyTypeName      string // 采购类型名称
	CustomerId       string // 供货商id
	WarehouseId      string // 仓库id
	BankId           string // 银行id
	ReceivableAmount string // 应付金额
	ReceivedAmount   string // 实付金额
	ArrearsAmount    string // 欠款金额
	CreateUserId     string // 创建人id
	UpdateUserId     string // 修改人id
	AuditUserId      string // 审核人id
	AuditStatus      string // 采购状态 0:未审核 1:已审核
	Remark           string // 备注
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	AuditAt          string // 审核时间
}

// buyColumns holds the columns for the table buy.
var buyColumns = BuyColumns{
	Id:               "id",
	BuyNo:            "buy_no",
	BuyDate:          "buy_date",
	BuyType:          "buy_type",
	BuyTypeName:      "buy_type_name",
	CustomerId:       "customer_id",
	WarehouseId:      "warehouse_id",
	BankId:           "bank_id",
	ReceivableAmount: "receivable_amount",
	ReceivedAmount:   "received_amount",
	ArrearsAmount:    "arrears_amount",
	CreateUserId:     "create_user_id",
	UpdateUserId:     "update_user_id",
	AuditUserId:      "audit_user_id",
	AuditStatus:      "audit_status",
	Remark:           "remark",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	AuditAt:          "audit_at",
}

// NewBuyDao creates and returns a new DAO object for table data access.
func NewBuyDao() *BuyDao {
	return &BuyDao{
		group:   "default",
		table:   "buy",
		columns: buyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BuyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BuyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BuyDao) Columns() BuyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BuyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BuyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *BuyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
