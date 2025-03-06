// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BankDao is the data access object for the table bank.
type BankDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of the current DAO.
	columns BankColumns // columns contains all the column names of Table for convenient usage.
}

// BankColumns defines and stores column names for the table bank.
type BankColumns struct {
	Id             string // 银行id
	Name           string // 银行名称
	Holder         string // 开户人
	Account        string // 银行账号
	Address        string // 开户银行地址
	Status         string // 银行状态 0:禁用 1:启用
	Active         string // 银行默认 0:否 1:是
	OpeningBalance string // 期初余额
	Balance        string // 余额
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// bankColumns holds the columns for the table bank.
var bankColumns = BankColumns{
	Id:             "id",
	Name:           "name",
	Holder:         "holder",
	Account:        "account",
	Address:        "address",
	Status:         "status",
	Active:         "active",
	OpeningBalance: "opening_balance",
	Balance:        "balance",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewBankDao creates and returns a new DAO object for table data access.
func NewBankDao() *BankDao {
	return &BankDao{
		group:   "default",
		table:   "bank",
		columns: bankColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BankDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BankDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BankDao) Columns() BankColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BankDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BankDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *BankDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
