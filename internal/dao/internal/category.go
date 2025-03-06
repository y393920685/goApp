// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CategoryDao is the data access object for the table category.
type CategoryDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns CategoryColumns // columns contains all the column names of Table for convenient usage.
}

// CategoryColumns defines and stores column names for the table category.
type CategoryColumns struct {
	Id        string // 产品类别id
	Name      string // 产品类别名称
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// categoryColumns holds the columns for the table category.
var categoryColumns = CategoryColumns{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCategoryDao creates and returns a new DAO object for table data access.
func NewCategoryDao() *CategoryDao {
	return &CategoryDao{
		group:   "default",
		table:   "category",
		columns: categoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CategoryDao) Columns() CategoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
