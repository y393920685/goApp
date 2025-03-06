package hello

import (
	"context"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "goApp/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	return nil, gerror.NewCode(gcode.CodeOK)
}
