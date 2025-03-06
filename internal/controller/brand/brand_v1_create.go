package brand

import (
	"context"

	v1 "goApp/api/brand/v1"
	"goApp/internal/model"
	"goApp/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.Brand().Create(ctx, &model.BrandCreateInput{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return
}
