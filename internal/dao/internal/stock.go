// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StockDao is the data access object for the table stock.
type StockDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns StockColumns // columns contains all the column names of Table for convenient usage.
}

// StockColumns defines and stores column names for the table stock.
type StockColumns struct {
	Id          string // 库存id
	ProductId   string // 产品id
	WarehouseId string // 仓库id
	StockNum    string // 库存数量
	StockPrice  string // 库存单价
	StockTotal  string // 库存总价
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// stockColumns holds the columns for the table stock.
var stockColumns = StockColumns{
	Id:          "id",
	ProductId:   "product_id",
	WarehouseId: "warehouse_id",
	StockNum:    "stock_num",
	StockPrice:  "stock_price",
	StockTotal:  "stock_total",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewStockDao creates and returns a new DAO object for table data access.
func NewStockDao() *StockDao {
	return &StockDao{
		group:   "default",
		table:   "stock",
		columns: stockColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StockDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StockDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StockDao) Columns() StockColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StockDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StockDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *StockDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
