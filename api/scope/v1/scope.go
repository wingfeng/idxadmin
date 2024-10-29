package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

type GetReq struct {
	g.Meta `path:"/scope/{id}" tags:"Scope" method:"get" summary:"get scope info by id"`
	Id     int64
}
type GetRes struct {
	entity.Scopes
}
type SaveReq struct {
	g.Meta `path:"/scope/" tags:"Scope" method:"put" summary:"insert or update a scope"`
	entity.Scopes
}
type SaveRes struct {
}
type ListReq struct {
	g.Meta `path:"/scope/list" tags:"Scope" method:"post" summary:"get a page of scopes"`
	model.PageReq
}
type ListRes struct {
	model.PageRes
}
type DeleteReq struct {
	Id     int64
	g.Meta `path:"/scope/{Id}" tags:"Scope" method:"delete" summary:"delete a scope by id"`
}
type DeleteRes struct {
}
