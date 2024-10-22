package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/do"
)

type GetReq struct {
	g.Meta `path:"/scope/{id}" tags:"scope" method:"get" summary:"get scope info by id"`
	Id     int64
}
type GetRes struct {
	do.Scopes
}
type SaveReq struct {
	g.Meta `path:"/scope/" tags:"scope" method:"put" summary:"insert or update a scope"`
	do.Scopes
}
type SaveRes struct {
}
type PageReq struct {
	g.Meta `path:"/scope/list" tags:"scope" method:"post" summary:"get a page of scopes"`
	model.PageReq
}
type PageRes struct {
	model.PageRes
}
type DeleteReq struct {
	Id     int64
	g.Meta `path:"/scope/{Id}" tags:"scope" method:"delete" summary:"delete a scope by id"`
}
type DeleteRes struct {
}
