// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure of table product for DAO operations like Where/Data.
type Product struct {
	g.Meta        `orm:"table:product, do:true"`
	Id            interface{} // 产品id
	Name          interface{} // 产品名称
	Pinyin        interface{} // 拼音码
	CategoryId    interface{} // 产品类别id
	Unitd         interface{} // 单位
	Color         interface{} // 颜色
	Brand         interface{} // 品牌
	Spec          interface{} // 规格
	Code          interface{} // 条码
	BuyPrice      interface{} // 销售价格
	SalePrice     interface{} // 进货价格
	VipPrice      interface{} // VIP会员价
	LowSalePrice  interface{} // 最低售价
	HightBuyPrice interface{} // 最高进价
	StockUpper    interface{} // 库存上限
	StockLower    interface{} // 库存下限
	Status        interface{} // 产品状态 0:禁用 1:启用
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
