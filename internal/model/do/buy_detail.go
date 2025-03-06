// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BuyDetail is the golang structure of table buy_detail for DAO operations like Where/Data.
type BuyDetail struct {
	g.Meta       `orm:"table:buy_detail, do:true"`
	Id           interface{} // 采购明细id
	BuyNo        interface{} // 采购单号
	ProductId    interface{} // 产品id
	ProductNum   interface{} // 采购数量
	ProductPrice interface{} // 采购单价
	ProductTotal interface{} // 采购总价
	Remark       interface{} // 备注
}
