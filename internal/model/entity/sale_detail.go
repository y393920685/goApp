// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SaleDetail is the golang structure for table sale_detail.
type SaleDetail struct {
	Id           uint    `json:"id"           orm:"id"            description:"销售明细id"`
	SaleNo       string  `json:"saleNo"       orm:"sale_no"       description:"销售单号"`
	ProductId    uint    `json:"productId"    orm:"product_id"    description:"产品id"`
	ProductNum   uint    `json:"productNum"   orm:"product_num"   description:"销售数量"`
	ProductPrice float64 `json:"productPrice" orm:"product_price" description:"销售单价"`
	ProductTotal float64 `json:"productTotal" orm:"product_total" description:"销售总价"`
	Profit       float64 `json:"profit"       orm:"profit"        description:"销售利润"`
	Remark       string  `json:"remark"       orm:"remark"        description:"备注"`
}
