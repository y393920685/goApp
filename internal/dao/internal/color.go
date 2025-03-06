// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ColorDao is the data access object for the table color.
type ColorDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns ColorColumns // columns contains all the column names of Table for convenient usage.
}

// ColorColumns defines and stores column names for the table color.
type ColorColumns struct {
	Id        string // 颜色id
	Name      string // 颜色名称
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// colorColumns holds the columns for the table color.
var colorColumns = ColorColumns{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewColorDao creates and returns a new DAO object for table data access.
func NewColorDao() *ColorDao {
	return &ColorDao{
		group:   "default",
		table:   "color",
		columns: colorColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ColorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ColorDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ColorDao) Columns() ColorColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ColorDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ColorDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ColorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
