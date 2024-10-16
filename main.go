package main

import (
	_ "github.com/wingfeng/idxadmin/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/wingfeng/idxadmin/internal/logic"

	"github.com/wingfeng/idxadmin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
