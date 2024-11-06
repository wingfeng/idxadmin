package conf

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type Options struct {
	JWKsUrl   string
	EnabledUI bool
	UIPath    string
	RedisUrl  string
	RedisDB   int
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
	e, err = g.Cfg().Get(ctx, "cache.redis_db")
	if err != nil {
		op.RedisDB = 0
	} else {
		op.RedisDB = e.Int()
	}
	e, err = g.Cfg().Get(ctx, "cache.redis_url")
	if err != nil {
		op.RedisUrl = ""
	} else {
		op.RedisUrl = e.String()
	}
	return op
}
