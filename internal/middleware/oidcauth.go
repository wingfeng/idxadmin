package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wingfeng/idxadmin/internal/conf"
	"github.com/wingfeng/idxadmin/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	Key  keyfunc.Keyfunc
	once sync.Once
)

func getKey(ctx context.Context) keyfunc.Keyfunc {

	once.Do(func() {
		url := conf.GetOptions(ctx).JWKsUrl

		k, err := keyfunc.NewDefaultCtx(ctx, []string{url}) // Context is used to end the refresh goroutine.
		if err != nil {
			log.Fatalf("Failed to create a keyfunc.Keyfunc from the server's URL.\nError: %s", err)
		}
		Key = k
	})

	return Key
}

func OidcAuth(c *ghttp.Request) {
	authHeader := c.GetHeader(consts.AUTHORIZATION_HEADER)
	token, err := verifyHeader(authHeader)
	getKey(c.GetCtx())
	if err != nil {

		g.Log().Error(c.Context(), "解析token失败", err)
		c.Response.Status = 401
		c.Response.WriteJson(g.Map{"Code": 401, "Msg": "解析token失败", "Data": err.Error()})
		c.Exit()
		//	c.AbortWithStatusJSON(401, base.SysResult{Code: 401, Msg: "解析token失败", Data: err.Error()})

		return
	}

	claims, err := extractClaims(token)
	if err != nil {
		g.Log().Error(c.Context(), "解析claims失败", err)
		c.Response.Status = 401

		c.Response.WriteJson(g.Map{"code": 401, "messsage": "解析claims失败", "data": err.Error()})
		c.Exit()
		//	c.AbortWithStatusJSON(401, base.SysResult{Code: 401, Msg: "解析claims失败", Data: err.Error()})

	}
	g.Log().Debug(c.Context(), "claims", claims)
	sub, err := claims.GetSubject()
	if err != nil {
		g.Log().Error(c.Context(), "claims.GetSubject()", err)
		c.Response.Status = 401

		c.Response.WriteJson(g.Map{"code": 401, "messsage": "解析claims失败", "data": err.Error()})
		c.Exit()
	}
	c.SetCtxVar(consts.SUBJECT_KEY, sub)
	c.SetCtxVar(consts.ACCOUNT_KEY, claims[consts.ACCOUNT_KEY])
	c.SetCtxVar(consts.EMAIL_KEY, claims[consts.EMAIL_KEY])
	c.SetCtxVar(consts.ROLES_KEY, claims[consts.ROLES_KEY])
	c.Middleware.Next()
}
func verifyHeader(header string) (string, error) {
	if strings.EqualFold(header, "") {
		return "", fmt.Errorf("http header %s为空", consts.AUTHORIZATION_HEADER)
	}
	s := strings.Split(header, " ")
	if len(s) != 2 {
		return "", fmt.Errorf("验证Header失败! Header:%s", header)
	} else if !strings.EqualFold(s[0], consts.BEARER) {
		return "", errors.New("验证方式错误,请使用Bearer验证")
	}
	return s[1], nil

}
func extractClaims(tokenStr string) (jwt.MapClaims, error) {

	// 基于公钥验证Token合法性
	token, err := jwt.Parse(tokenStr, Key.Keyfunc)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, nil

	} else {
		return nil, fmt.Errorf("invalid JWT token")
	}

}
