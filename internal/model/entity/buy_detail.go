// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BuyDetail is the golang structure for table buy_detail.
type BuyDetail struct {
	Id           uint    `json:"id"           orm:"id"            description:"采购明细id"`
	BuyNo        string  `json:"buyNo"        orm:"buy_no"        description:"采购单号"`
	ProductId    uint    `json:"productId"    orm:"product_id"    description:"产品id"`
	ProductNum   uint    `json:"productNum"   orm:"product_num"   description:"采购数量"`
	ProductPrice float64 `json:"productPrice" orm:"product_price" description:"采购单价"`
	ProductTotal float64 `json:"productTotal" orm:"product_total" description:"采购总价"`
	Remark       string  `json:"remark"       orm:"remark"        description:"备注"`
}
