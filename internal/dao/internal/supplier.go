// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SupplierDao is the data access object for the table supplier.
type SupplierDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns SupplierColumns // columns contains all the column names of Table for convenient usage.
}

// SupplierColumns defines and stores column names for the table supplier.
type SupplierColumns struct {
	Id             string // 供货商id
	Name           string // 供货商名称
	Pinyin         string // 供货商名称拼音
	Contact        string // 供货商联系人
	Phone          string // 供货商联系电话
	CompanyName    string // 供货商公司名称
	Address        string // 供货商地址
	TaxNumber      string // 供货商纳税号
	BankName       string // 供货商开户行
	BankAccount    string // 供货商银行账号
	BankHolder     string // 银行开户人
	OpeningBalance string // 供货商期初余额
	Status         string // 供货商状态 0:禁用 1:启用
	Active         string // 供货商默认 0:否 1:是
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// supplierColumns holds the columns for the table supplier.
var supplierColumns = SupplierColumns{
	Id:             "id",
	Name:           "name",
	Pinyin:         "pinyin",
	Contact:        "contact",
	Phone:          "phone",
	CompanyName:    "company_name",
	Address:        "address",
	TaxNumber:      "tax_number",
	BankName:       "bank_name",
	BankAccount:    "bank_account",
	BankHolder:     "bank_holder",
	OpeningBalance: "opening_balance",
	Status:         "status",
	Active:         "active",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewSupplierDao creates and returns a new DAO object for table data access.
func NewSupplierDao() *SupplierDao {
	return &SupplierDao{
		group:   "default",
		table:   "supplier",
		columns: supplierColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SupplierDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SupplierDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SupplierDao) Columns() SupplierColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SupplierDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SupplierDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SupplierDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
