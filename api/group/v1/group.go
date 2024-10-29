package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

type GetReq struct {
	Id     int64 `json:"id" v:"required#组ID不能为空" dc:"组ID"`
	g.Meta `path:"/group/{Id}" method:"get" tags:"Group" summary:"获取组信息"`
}

type GetRes struct {
	g.Meta `mime:"application/json" type:"object" example:"model.Group"`
	entity.Groups
}
type ListReq struct {
	g.Meta `path:"/group/list" method:"post" tags:"Group" summary:"获取组列表"`
	model.PageReq
}
type ListRes struct {
	g.Meta `mime:"application/json" type:"object" example:"model.PageRes"`
	model.PageRes
}
type SaveReq struct {
	g.Meta `path:"/group/" method:"put" tags:"Group" summary:"新增或更新组"`
	entity.Groups
}
type SaveRes struct {
}
type DeleteReq struct {
	Id     int64 `json:"id" v:"required#组ID不能为空" dc:"组ID"`
	g.Meta `path:"/group/{Id}" method:"delete" tags:"Group" summary:"删除组"`
}
type DeleteRes struct {
}
type MembersReq struct {
	Id     int64 `json:"id" v:"required#组ID不能为空" dc:"组ID"`
	g.Meta `path:"/group/{Id}/members" method:"post" tags:"Group" summary:"获取组成员"`
	model.PageReq
}
type MembersRes struct {
	g.Meta `mime:"application/json" type:"object" example:"model.PageRes"`
	model.PageRes
}
type AddMemberReq struct {
	Id      int64 `json:"id" v:"required#组ID不能为空" dc:"组ID"`
	g.Meta  `path:"/group/{Id}/members" method:"put" tags:"Group" summary:"添加组成员"`
	UserIds []int64 `json:"userIds" v:"required#用户ID不能为空" dc:"用户ID"`
}
type AddMemberRes struct {
}
type RemoveMemberReq struct {
	Id      int64 `json:"id" v:"required#组ID不能为空" dc:"组ID"`
	g.Meta  `path:"/group/{Id}/members" method:"delete" tags:"Group" summary:"删除组成员"`
	UserIds []int64 `json:"userIds" v:"required#用户ID不能为空" dc:"用户ID"`
}
type RemoveMemberRes struct {
}
