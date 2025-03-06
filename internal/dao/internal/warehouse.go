// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WarehouseDao is the data access object for the table warehouse.
type WarehouseDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns WarehouseColumns // columns contains all the column names of Table for convenient usage.
}

// WarehouseColumns defines and stores column names for the table warehouse.
type WarehouseColumns struct {
	Id        string // 仓库id
	Name      string // 仓库名称
	Contact   string // 仓库联系人
	Phone     string // 仓库联系电话
	Status    string // 仓库状态 0:禁用 1:启用
	Active    string // 仓库默认 0:否 1:是
	Address   string // 仓库地址
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// warehouseColumns holds the columns for the table warehouse.
var warehouseColumns = WarehouseColumns{
	Id:        "id",
	Name:      "name",
	Contact:   "contact",
	Phone:     "phone",
	Status:    "status",
	Active:    "active",
	Address:   "address",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewWarehouseDao creates and returns a new DAO object for table data access.
func NewWarehouseDao() *WarehouseDao {
	return &WarehouseDao{
		group:   "default",
		table:   "warehouse",
		columns: warehouseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WarehouseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WarehouseDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WarehouseDao) Columns() WarehouseColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WarehouseDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WarehouseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *WarehouseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
