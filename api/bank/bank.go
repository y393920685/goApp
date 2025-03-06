// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package bank

import (
	"context"

	"goApp/api/bank/v1"
)

type IBankV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
}
