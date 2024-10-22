package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/do"
)

type GetReq struct {
	g.Meta `path:"/user/{id}" tags:"user" method:"get" summary:"get client info by id"`
	Id     int64
}
type GetRes struct {
	do.Users
}
type DeleteReq struct {
	g.Meta `path:"/user/{Id}" tags:"user" method:"delete" summary:"delete a user by id"`
	Id     int64
}
type DeleteRes struct {
}
type PageReq struct {
	g.Meta `path:"/user/list" tags:"user" method:"post" summary:"get a page of users"`
	model.PageReq
}
type PageRes struct {
	model.PageRes
}
type SaveReq struct {
	do.Users
	g.Meta `path:"/user/" tags:"user" method:"put" summary:"insert or update a user"`
}
type SaveRes struct {
}
type ResetPwdReq struct {
	g.Meta `path:"/user/resetpwd" tags:"user" method:"post" summary:"reset password"`
	Id     string
}
type ResetPwdRes struct {
	NewPwd string
}
