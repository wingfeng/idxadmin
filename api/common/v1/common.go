package v1

import "github.com/gogf/gf/v2/frame/g"

type NewIdReq struct {
	g.Meta `path:"/common/newid" tags:"common" method:"get" summary:"Get a new id"`
}
type NewIdRes struct {
	Id string `json:"id"`
}
