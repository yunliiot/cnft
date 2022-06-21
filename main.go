package main

import (
	_ "cnft/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"cnft/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
