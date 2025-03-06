// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goApp/internal/model"
)

type (
	IBank interface {
		Create(ctx context.Context, in *model.BankCreateInput) (err error)
		Update(ctx context.Context, id int64, in *model.BankUpdateInput) (err error)
		Delete(ctx context.Context, id int64) (err error)
		GetList(ctx context.Context, in *model.BankGetListInput) (out *model.BankGetListOutput, err error)
		GetOne(ctx context.Context, id int64) (out *model.BankGetOneOutput, err error)
	}
)

var (
	localBank IBank
)

func Bank() IBank {
	if localBank == nil {
		panic("implement not found for interface IBank, forgot register?")
	}
	return localBank
}

func RegisterBank(i IBank) {
	localBank = i
}
