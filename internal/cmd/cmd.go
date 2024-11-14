package cmd

import (
	"context"
	"path"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
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
			api := s.GetOpenApi()
			oauth := goai.SecurityRequirement{
				"idxauth": {"idxauth"},
			}
			api.Security = &goai.SecurityRequirements{
				oauth,
			}

			api.Components.SecuritySchemes = goai.SecuritySchemes{
				"idxauth": goai.SecuritySchemeRef{
					Value: &goai.SecurityScheme{Type: "oauth",
						OpenIdConnectUrl: "http://localhost:9097/idx/",
						Flows: &goai.OAuthFlows{AuthorizationCode: &goai.OAuthFlow{
							AuthorizationURL: "/oauth2/authorize",
							TokenURL:         "/idx/oauth2/token",
							Scopes: map[string]string{
								"admin": "admin",
							},
						}}},
				},
			}
			//	api.Components.Schemas

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
				s.BindHandler("GET:/", func(r *ghttp.Request) {
					r.Response.RedirectTo("/ui")
				})
			}

			s.Run()
			return nil
		},
	}
)
