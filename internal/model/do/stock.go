// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Stock is the golang structure of table stock for DAO operations like Where/Data.
type Stock struct {
	g.Meta      `orm:"table:stock, do:true"`
	Id          interface{} // 库存id
	ProductId   interface{} // 产品id
	WarehouseId interface{} // 仓库id
	StockNum    interface{} // 库存数量
	StockPrice  interface{} // 库存单价
	StockTotal  interface{} // 库存总价
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
