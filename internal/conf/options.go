package conf

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type Options struct {
	JWKsUrl   string
	EnabledUI bool
	UIPath    string
}

// jwt:
//
//	jwkurl: "http://localhost:9097/idx/.well-known/jwks"
func GetOptions(ctx context.Context) *Options {
	op := &Options{}
	e, err := g.Cfg().Get(ctx, "ui.enabled")
	if err != nil {
		op.EnabledUI = false
	} else {
		op.EnabledUI = e.Bool()
	}
	e, err = g.Cfg().Get(ctx, "ui.path")
	if err != nil {
		op.UIPath = ""
	} else {
		op.UIPath = e.String()
	}
	e, err = g.Cfg().Get(ctx, "jwt.jwkurl")
	if err != nil {
		op.JWKsUrl = ""
	} else {
		op.JWKsUrl = e.String()
	}
	return op
}
