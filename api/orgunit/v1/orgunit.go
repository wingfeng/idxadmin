package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

type TreeReq struct {
	g.Meta `path:"/ou/tree" tags:"organization unit" method:"get" summary:"get Organization tree"`
}
type TreeRes struct {
	g.Meta `mime:"json/application"`
	Items  []*entity.OrganizationUnitsEx `json:"items"`
}
type GetReq struct {
	g.Meta `path:"/ou/{id}" tags:"organization unit" method:"get" summary:"get Organization info by id"`
	Id     int64 `v:"required"`
}
type GetRes struct {
	entity.OrganizationUnits
}
type SaveReq struct {
	g.Meta `path:"/ou/" tags:"organization unit" method:"put" summary:"insert or update a Organization"`
	entity.OrganizationUnits
}
type SaveRes struct {
}
type DeleteReq struct {
	Id     int64 `v:"required"`
	g.Meta `path:"/ou/{Id}" tags:"organization unit" method:"delete" summary:"delete a Organization by id"`
}
type DeleteRes struct {
}
type ListReq struct {
	g.Meta `path:"/ou/list" tags:"organization unit" method:"post" summary:"get a page of Organizations"`
	model.PageReq
}
type ListRes struct {
	model.PageRes
}
