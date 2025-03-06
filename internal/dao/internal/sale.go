// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SaleDao is the data access object for the table sale.
type SaleDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of the current DAO.
	columns SaleColumns // columns contains all the column names of Table for convenient usage.
}

// SaleColumns defines and stores column names for the table sale.
type SaleColumns struct {
	Id               string // 销售id
	SaleNo           string // 销售单号
	SaleDate         string // 销售日期
	SaleType         string // 销售类型 0:普通销售 1:其他出库
	SaleTypeName     string // 销售类型名称
	CustomerId       string // 客户id
	WarehouseId      string // 仓库id
	BankId           string // 银行id
	ReceivableAmount string // 应收金额
	ReceivedAmount   string // 实收金额
	SaleProfit       string // 销售利润
	ArrearsAmount    string // 欠款金额
	CreateUserId     string // 创建人id
	UpdateUserId     string // 修改人id
	AuditUserId      string // 审核人id
	AuditStatus      string // 销售状态 0:未审核 1:已审核
	Remark           string // 备注
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	AuditAt          string // 审核时间
}

// saleColumns holds the columns for the table sale.
var saleColumns = SaleColumns{
	Id:               "id",
	SaleNo:           "sale_no",
	SaleDate:         "sale_date",
	SaleType:         "sale_type",
	SaleTypeName:     "sale_type_name",
	CustomerId:       "customer_id",
	WarehouseId:      "warehouse_id",
	BankId:           "bank_id",
	ReceivableAmount: "receivable_amount",
	ReceivedAmount:   "received_amount",
	SaleProfit:       "sale_profit",
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

// NewSaleDao creates and returns a new DAO object for table data access.
func NewSaleDao() *SaleDao {
	return &SaleDao{
		group:   "default",
		table:   "sale",
		columns: saleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SaleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SaleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SaleDao) Columns() SaleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SaleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SaleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SaleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
