package brand

import (
	"context"

	"goApp/api"
	v1 "goApp/api/brand/v1"
	"goApp/internal/model"
	"goApp/internal/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	out, err := service.Brand().GetList(ctx, &model.BrandGetListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Search:   req.Search,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetListRes{
		PageListRes: &api.PageListRes{
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
			Total:    out.Total,
		},
		List: out.List,
	}, nil
}
