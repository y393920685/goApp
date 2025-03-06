// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SaleDetailDao is the data access object for the table sale_detail.
type SaleDetailDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SaleDetailColumns // columns contains all the column names of Table for convenient usage.
}

// SaleDetailColumns defines and stores column names for the table sale_detail.
type SaleDetailColumns struct {
	Id           string // 销售明细id
	SaleNo       string // 销售单号
	ProductId    string // 产品id
	ProductNum   string // 销售数量
	ProductPrice string // 销售单价
	ProductTotal string // 销售总价
	Profit       string // 销售利润
	Remark       string // 备注
}

// saleDetailColumns holds the columns for the table sale_detail.
var saleDetailColumns = SaleDetailColumns{
	Id:           "id",
	SaleNo:       "sale_no",
	ProductId:    "product_id",
	ProductNum:   "product_num",
	ProductPrice: "product_price",
	ProductTotal: "product_total",
	Profit:       "profit",
	Remark:       "remark",
}

// NewSaleDetailDao creates and returns a new DAO object for table data access.
func NewSaleDetailDao() *SaleDetailDao {
	return &SaleDetailDao{
		group:   "default",
		table:   "sale_detail",
		columns: saleDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SaleDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SaleDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SaleDetailDao) Columns() SaleDetailColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SaleDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SaleDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SaleDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
