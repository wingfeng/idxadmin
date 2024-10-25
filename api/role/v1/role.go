package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

type GetReq struct {
	g.Meta `path:"/role/{Id}" method:"get" tags:"Role" summary:"获取角色信息"`
	Id     int64 `json:"id" v:"required#角色ID不能为空" dc:"角色ID"`
}
type GetRes struct {
	g.Meta `mime:"application/json" type:"object" example:"model.Role"`
	entity.Roles
}
type ListReq struct {
	g.Meta `path:"/role/list" method:"post" tags:"Role" summary:"获取角色列表"`
	model.PageReq
}
type ListRes struct {
	g.Meta `mime:"application/json" type:"object" example:"model.PageRes"`
	model.PageRes
}
type SaveReq struct {
	g.Meta `path:"/role/" method:"put" tags:"Role" summary:"新增或更新角色"`
	entity.Roles
}
type SaveRes struct {
}
type DeleteReq struct {
	g.Meta `path:"/role/{Id}" method:"delete" tags:"Role" summary:"删除角色"`
	Id     int64 `json:"id" v:"required#角色ID不能为空" dc:"角色ID"`
}
type DeleteRes struct {
}
type MembersReq struct {
	g.Meta `path:"/role/{Id}/members" method:"post" tags:"Role" summary:"获取角色成员"`
	Id     int64 `json:"id" v:"required#角色ID不能为空" dc:"角色ID"`
	model.PageReq
}
type MembersRes struct {
	g.Meta `mime:"application/json" type:"object" example:"model.PageRes"`
	model.PageRes
}
