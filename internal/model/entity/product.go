// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure for table product.
type Product struct {
	Id            uint        `json:"id"            orm:"id"              description:"产品id"`
	Name          string      `json:"name"          orm:"name"            description:"产品名称"`
	Pinyin        string      `json:"pinyin"        orm:"pinyin"          description:"拼音码"`
	CategoryId    uint        `json:"categoryId"    orm:"category_id"     description:"产品类别id"`
	Unitd         string      `json:"unitd"         orm:"unitd"           description:"单位"`
	Color         string      `json:"color"         orm:"color"           description:"颜色"`
	Brand         string      `json:"brand"         orm:"brand"           description:"品牌"`
	Spec          string      `json:"spec"          orm:"spec"            description:"规格"`
	Code          string      `json:"code"          orm:"code"            description:"条码"`
	BuyPrice      float64     `json:"buyPrice"      orm:"buy_price"       description:"销售价格"`
	SalePrice     float64     `json:"salePrice"     orm:"sale_price"      description:"进货价格"`
	VipPrice      float64     `json:"vipPrice"      orm:"vip_price"       description:"VIP会员价"`
	LowSalePrice  float64     `json:"lowSalePrice"  orm:"low_sale_price"  description:"最低售价"`
	HightBuyPrice float64     `json:"hightBuyPrice" orm:"hight_buy_price" description:"最高进价"`
	StockUpper    uint        `json:"stockUpper"    orm:"stock_upper"     description:"库存上限"`
	StockLower    uint        `json:"stockLower"    orm:"stock_lower"     description:"库存下限"`
	Status        string      `json:"status"        orm:"status"          description:"产品状态 0:禁用 1:启用"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`
}
