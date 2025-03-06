// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BrandDao is the data access object for the table brand.
type BrandDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns BrandColumns // columns contains all the column names of Table for convenient usage.
}

// BrandColumns defines and stores column names for the table brand.
type BrandColumns struct {
	Id        string // 品牌id
	Name      string // 品牌名称
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// brandColumns holds the columns for the table brand.
var brandColumns = BrandColumns{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewBrandDao creates and returns a new DAO object for table data access.
func NewBrandDao() *BrandDao {
	return &BrandDao{
		group:   "default",
		table:   "brand",
		columns: brandColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BrandDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BrandDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BrandDao) Columns() BrandColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BrandDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BrandDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *BrandDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
