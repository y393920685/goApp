// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BankRecordDao is the data access object for the table bank_record.
type BankRecordDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns BankRecordColumns // columns contains all the column names of Table for convenient usage.
}

// BankRecordColumns defines and stores column names for the table bank_record.
type BankRecordColumns struct {
	Id        string // 银行记录id
	BankId    string // 银行id
	Type      string // 记录类型 CG:采购入库 XS:销售出库
	TypeName  string // 记录类型名称
	BillNo    string // 单据号
	Amount    string // 交易金额
	Balance   string // 交易后余额
	UserId    string // 创建人id
	CreatedAt string // 创建时间
}

// bankRecordColumns holds the columns for the table bank_record.
var bankRecordColumns = BankRecordColumns{
	Id:        "id",
	BankId:    "bank_id",
	Type:      "type",
	TypeName:  "type_name",
	BillNo:    "bill_no",
	Amount:    "amount",
	Balance:   "balance",
	UserId:    "user_id",
	CreatedAt: "created_at",
}

// NewBankRecordDao creates and returns a new DAO object for table data access.
func NewBankRecordDao() *BankRecordDao {
	return &BankRecordDao{
		group:   "default",
		table:   "bank_record",
		columns: bankRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BankRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BankRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BankRecordDao) Columns() BankRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BankRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BankRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *BankRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
