package brand

import (
	"context"

	v1 "goApp/api/brand/v1"
	"goApp/internal/model"
	"goApp/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = service.Brand().Update(ctx, req.Id, &model.BrandUpdateInput{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return
}
