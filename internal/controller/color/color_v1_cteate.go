package color

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"goApp/api/color/v1"
)

func (c *ControllerV1) Cteate(ctx context.Context, req *v1.CteateReq) (res *v1.CteateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
