package main

import (
	"flag"

	_ "github.com/wingfeng/idxadmin/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/wingfeng/idxadmin/internal/logic"

	"github.com/wingfeng/idxadmin/internal/cmd"
)

func main() {
	initPreset := flag.Bool("init", false, "Initialize the preset client and user")
	flag.Parse()
	if *initPreset {
		cmd.Preset.Run(gctx.GetInitCtx())
		return
	}
	cmd.Main.Run(gctx.GetInitCtx())
}
