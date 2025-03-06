package model

import (
	"goApp/internal/model/entity"

	"github.com/shopspring/decimal"
)

type BankBaseType struct {
	Name    string
	Holder  string
	Account string
	Address string
	Status  string
	Active  string
}

type BankCreateInput struct {
	BankBaseType
	OpeningBalance decimal.Decimal
	Balance        decimal.Decimal
}
type BankUpdateInput struct {
	BankBaseType
}
type BankGetListInput struct {
	PageNum  int
	PageSize int
	Search   string
}
type BankGetListOutput struct {
	List  []*entity.Bank
	Total int64
}
type BankGetOneOutput struct {
	*entity.Bank
}
