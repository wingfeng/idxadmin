package cmd

import (
	"context"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	clientReqV1 "github.com/wingfeng/idxadmin/api/client/v1"
	v1 "github.com/wingfeng/idxadmin/api/user/v1"
	"github.com/wingfeng/idxadmin/internal/consts"
	"github.com/wingfeng/idxadmin/internal/model/entity"
	"github.com/wingfeng/idxadmin/service"
)

var (
	Preset = gcmd.Command{
		Name:  "preset",
		Usage: "preset",
		Brief: "Initialize the preset client and user",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, "Begin to initialize preset data...")
			ctx = context.WithValue(ctx, consts.SUBJECT_KEY, "1838872840128958464")
			ctx = context.WithValue(ctx, consts.ACCOUNT_KEY, "admin")
			// ctx.SetCtxVar(consts.SUBJECT_KEY, "0")
			// ctx.SetCtxVar(consts.USER_NAME, "admin")
			admin := &entity.Users{
				Id:       1838872840128958464,
				UserName: "admin",

				Email:       "admin@idx.local",
				PhoneNumber: "123456789",
			}

			service.User().Save(ctx, &v1.SaveReq{
				Users: *admin,
			})
			idStr := strconv.Itoa(int(admin.Id))
			newPwd, err := service.User().ResetPwd(ctx, &v1.ResetPwdReq{Id: idStr})
			if err != nil {
				g.Log().Error(ctx, err)
				panic(err)
			}
			g.Log().Info(ctx, "admin password:", newPwd)
			adminui := &entity.Clients{
				Id:                 1838872840128958465,
				ClientName:         "IDX Admin UI",
				ClientId:           "adminui",
				Enabled:            true,
				GrantTypes:         "authorization_code",
				ClientUri:          "http://localhost:8090/",
				Scopes:             "openid profile email",
				Description:        "IDX Admin UI",
				RedirectUris:       "http://localhost:*/ui/auth/callback http://localhost:5666/auth/callback http://localhost:*/auth/callback http://localhost/ui/auth/callback",
				RequireSecret:      false,
				RequirePkce:        true,
				AllowPlainTextPkce: false,
			}
			err = service.Client().Save(ctx, &clientReqV1.SaveReq{
				Clients: *adminui,
			})
			if err != nil {
				g.Log().Error(ctx, "create admin ui client error", "error", err)
				panic(err)
			}
			req := &clientReqV1.SetCORSReq{}
			req.ClientId = adminui.Id
			req.Origins = []string{"*", "http://localhost"}
			err = service.Client().SetCORS(ctx, req)
			if err != nil {
				g.Log().Error(ctx, "create admin ui client cors error", "error", err)
				panic(err)
			}

			// service.Client().Save(ctx, &v1.SaveReq{
			// 	ClientId:     "idxadmin",
			// 	ClientSecret: "secret",
			// 	Scopes:       "openid,profile,email,phone",
			// 	Description:  "idxadmin client",
			// })
			return nil
		},
	}
)
