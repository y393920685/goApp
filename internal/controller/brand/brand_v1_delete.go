package brand

import (
	"context"

	v1 "goApp/api/brand/v1"
	"goApp/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	if err = service.Brand().Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
