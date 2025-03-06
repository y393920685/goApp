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
	IBrand interface {
		Create(ctx context.Context, in *model.BrandCreateInput) (err error)
		Update(ctx context.Context, id int64, in *model.BrandUpdateInput) (err error)
		Delete(ctx context.Context, id int64) (err error)
		GetOne(ctx context.Context, id int64) (out *model.BrandGetOneOutput, err error)
		GetList(ctx context.Context, in *model.BrandGetListInput) (out *model.BrandGetListOutput, err error)
	}
)

var (
	localBrand IBrand
)

func Brand() IBrand {
	if localBrand == nil {
		panic("implement not found for interface IBrand, forgot register?")
	}
	return localBrand
}

func RegisterBrand(i IBrand) {
	localBrand = i
}
