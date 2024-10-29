package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

type GetReq struct {
	g.Meta `path:"/user/{id}" tags:"User" method:"get" summary:"get client info by id"`
	Id     int64 `v:"required"`
}
type GetRes struct {
	entity.Users
}
type DeleteReq struct {
	g.Meta `path:"/user/{Id}" tags:"User" method:"delete" summary:"delete a user by id"`
	Id     int64 `v:"required"`
}
type DeleteRes struct {
}
type ListReq struct {
	g.Meta `path:"/user/list" tags:"User" method:"post" summary:"get a page of users"`
	model.PageReq
}
type ListRes struct {
	model.PageRes
}
type SaveReq struct {
	entity.Users
	g.Meta `path:"/user/" tags:"User" method:"put" summary:"insert or update a user"`
}
type SaveRes struct {
}
type ResetPwdReq struct {
	g.Meta `path:"/user/resetpwd" tags:"User" method:"post" summary:"reset password"`
	Id     string `v:"required"`
}
type ResetPwdRes struct {
	NewPwd string
}
