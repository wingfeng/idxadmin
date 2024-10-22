package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/wingfeng/idxadmin/internal/controller/client"
	"github.com/wingfeng/idxadmin/internal/controller/orgunit"
	"github.com/wingfeng/idxadmin/internal/controller/user"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api/v1/oauth2", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(

					client.NewV1(),
				)
			})
			s.Group("/api/v1/system", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					user.NewV1(),
					orgunit.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
