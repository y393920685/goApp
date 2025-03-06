package model

import (
	"goApp/internal/consts"
	"goApp/internal/model/entity"

	"github.com/shopspring/decimal"
)

type BankRecordIncrementInput struct {
	BankId   int64                     // 银行id
	Type     consts.BankRecordType     // 记录类型 CG:采购入库 XS:销售出库
	TypeName consts.BankRecordTypeText // 记录类型名称
	BillNo   string                    // 单据号
	Amount   decimal.Decimal           // 交易金额
	Balance  decimal.Decimal           // 交易后余额
	UserId   int64                     // 创建人id
}
type BankRecordDecrementInput struct {
	BankId   int64                     // 银行id
	Type     consts.BankRecordType     // 记录类型 CG:采购入库 XS:销售出库
	TypeName consts.BankRecordTypeText // 记录类型名称
	BillNo   string                    // 单据号
	Amount   decimal.Decimal           // 交易金额
	Balance  decimal.Decimal           // 交易后余额
	UserId   int64                     // 创建人id
}
type BankRecordGetListInput struct {
	PageNum  int
	PageSize int
	Search   string
}
type BankRecordGetListOutput struct {
	List  []*entity.BankRecord
	Total int64
}
