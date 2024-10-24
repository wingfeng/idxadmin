package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/wingfeng/idxadmin/internal/controller/client"
	"github.com/wingfeng/idxadmin/internal/controller/common"
	"github.com/wingfeng/idxadmin/internal/controller/orgunit"
	"github.com/wingfeng/idxadmin/internal/controller/user"
	"github.com/wingfeng/idxadmin/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(middleware.OidcAuth)
				group.Group("/oauth2", func(group *ghttp.RouterGroup) {
					group.Bind(
						client.NewV1(),
					)
				})
				group.Group("/system", func(group *ghttp.RouterGroup) {
					group.Bind(
						user.NewV1(),
						orgunit.NewV1(),
						common.NewV1(),
					)
				})
			})
			s.SetIndexFolder(true)
			s.AddStaticPath("/ui", "E:\\gowork\\adminui\\apps\\web-antd\\dist")
			// s.BindHandler("GET:/ui/*", func(r *ghttp.Request){
			// 	r.Response.Writef()
			// })
			s.Run()
			return nil
		},
	}
)
