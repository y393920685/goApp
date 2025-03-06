// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BuyDetailDao is the data access object for the table buy_detail.
type BuyDetailDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns BuyDetailColumns // columns contains all the column names of Table for convenient usage.
}

// BuyDetailColumns defines and stores column names for the table buy_detail.
type BuyDetailColumns struct {
	Id           string // 采购明细id
	BuyNo        string // 采购单号
	ProductId    string // 产品id
	ProductNum   string // 采购数量
	ProductPrice string // 采购单价
	ProductTotal string // 采购总价
	Remark       string // 备注
}

// buyDetailColumns holds the columns for the table buy_detail.
var buyDetailColumns = BuyDetailColumns{
	Id:           "id",
	BuyNo:        "buy_no",
	ProductId:    "product_id",
	ProductNum:   "product_num",
	ProductPrice: "product_price",
	ProductTotal: "product_total",
	Remark:       "remark",
}

// NewBuyDetailDao creates and returns a new DAO object for table data access.
func NewBuyDetailDao() *BuyDetailDao {
	return &BuyDetailDao{
		group:   "default",
		table:   "buy_detail",
		columns: buyDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BuyDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BuyDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BuyDetailDao) Columns() BuyDetailColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BuyDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BuyDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *BuyDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
