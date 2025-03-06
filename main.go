package main

import (
	_ "goApp/internal/packed"

	_ "goApp/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"goApp/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
