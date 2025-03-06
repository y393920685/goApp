// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UnitDao is the data access object for the table unit.
type UnitDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of the current DAO.
	columns UnitColumns // columns contains all the column names of Table for convenient usage.
}

// UnitColumns defines and stores column names for the table unit.
type UnitColumns struct {
	Id        string // 单位id
	Name      string // 单位名称
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// unitColumns holds the columns for the table unit.
var unitColumns = UnitColumns{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUnitDao creates and returns a new DAO object for table data access.
func NewUnitDao() *UnitDao {
	return &UnitDao{
		group:   "default",
		table:   "unit",
		columns: unitColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UnitDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UnitDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UnitDao) Columns() UnitColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UnitDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UnitDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UnitDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
