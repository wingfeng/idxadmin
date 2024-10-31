package cmd

import (
	"context"
	"path"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/wingfeng/idxadmin/internal/conf"
	"github.com/wingfeng/idxadmin/internal/controller/client"
	"github.com/wingfeng/idxadmin/internal/controller/common"
	ggroup "github.com/wingfeng/idxadmin/internal/controller/group"
	"github.com/wingfeng/idxadmin/internal/controller/orgunit"
	"github.com/wingfeng/idxadmin/internal/controller/role"
	"github.com/wingfeng/idxadmin/internal/controller/scope"
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
						scope.NewV1(),
					)
				})
				group.Group("/system", func(group *ghttp.RouterGroup) {
					group.Bind(
						user.NewV1(),
						orgunit.NewV1(),
						common.NewV1(),
						role.NewV1(),
						ggroup.NewV1(),
					)
				})
			})
			options := conf.GetOptions(ctx)
			if options.EnabledUI {
				s.AddStaticPath("/ui", options.UIPath)

				s.AddStaticPath("/ui", options.UIPath)
				s.BindHandler("GET:/ui/*", func(r *ghttp.Request) {
					ext := path.Ext(r.URL.Path)
					if strings.EqualFold(ext, "") {
						r.Response.ServeFile(path.Join(options.UIPath, "index.html"))
					}

				})
			}

			s.Run()
			return nil
		},
	}
)
