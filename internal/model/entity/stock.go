// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Stock is the golang structure for table stock.
type Stock struct {
	Id          uint        `json:"id"          orm:"id"           description:"库存id"`
	ProductId   uint        `json:"productId"   orm:"product_id"   description:"产品id"`
	WarehouseId uint        `json:"warehouseId" orm:"warehouse_id" description:"仓库id"`
	StockNum    int         `json:"stockNum"    orm:"stock_num"    description:"库存数量"`
	StockPrice  float64     `json:"stockPrice"  orm:"stock_price"  description:"库存单价"`
	StockTotal  float64     `json:"stockTotal"  orm:"stock_total"  description:"库存总价"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
}
