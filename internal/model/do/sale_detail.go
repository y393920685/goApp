// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SaleDetail is the golang structure of table sale_detail for DAO operations like Where/Data.
type SaleDetail struct {
	g.Meta       `orm:"table:sale_detail, do:true"`
	Id           interface{} // 销售明细id
	SaleNo       interface{} // 销售单号
	ProductId    interface{} // 产品id
	ProductNum   interface{} // 销售数量
	ProductPrice interface{} // 销售单价
	ProductTotal interface{} // 销售总价
	Profit       interface{} // 销售利润
	Remark       interface{} // 备注
}
