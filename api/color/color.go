// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package color

import (
	"context"

	"goApp/api/color/v1"
)

type IColorV1 interface {
	Cteate(ctx context.Context, req *v1.CteateReq) (res *v1.CteateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
}
