// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CustomerDao is the data access object for the table customer.
type CustomerDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns CustomerColumns // columns contains all the column names of Table for convenient usage.
}

// CustomerColumns defines and stores column names for the table customer.
type CustomerColumns struct {
	Id             string // 客户id
	Name           string // 客户名称
	Pinyin         string // 客户名称拼音
	Contact        string // 客户联系人
	Phone          string // 客户联系电话
	CompanyName    string // 客户公司名称
	Address        string // 客户地址
	TaxNumber      string // 客户纳税号
	BankName       string // 客户开户行
	BankAccount    string // 客户银行账号
	BankHolder     string // 银行开户人
	OpeningBalance string // 客户期初余额
	Status         string // 客户状态 0:禁用 1:启用
	Active         string // 客户默认 0:否 1:是
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// customerColumns holds the columns for the table customer.
var customerColumns = CustomerColumns{
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

// NewCustomerDao creates and returns a new DAO object for table data access.
func NewCustomerDao() *CustomerDao {
	return &CustomerDao{
		group:   "default",
		table:   "customer",
		columns: customerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CustomerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CustomerDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CustomerDao) Columns() CustomerColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CustomerDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CustomerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CustomerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
