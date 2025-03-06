package brand

import (
	"context"

	v1 "goApp/api/brand/v1"
	"goApp/internal/service"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	out, err := service.Brand().GetOne(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{Brand: out.Brand}, nil
}
