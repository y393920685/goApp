package main

import (
	_ "goApp/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goApp/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
