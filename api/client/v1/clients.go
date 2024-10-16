package v1

import (
	"time"

	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/do"
	"github.com/wingfeng/idxadmin/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetReq struct {
	g.Meta `path:"/client/{Id}.info" tags:"Client" method:"get" summary:"get client info by id"`
	Id     int
}
type GetRes struct {
	g.Meta `mime:"json/application"`
	do.Clients
}

type PageReq struct {
	g.Meta `path:"/client/list" tags:"Client" method:"post" summary:"get a page of clients"`
	model.PageReq
}
type PageRes struct {
	g.Meta `mime:"json/application"`
	model.PageRes
}
type SaveReq struct {
	g.Meta `path:"/client/" tags:"Client" method:"put" summary:"insert or update a client"`
	do.Clients
}
type SaveRes struct {
}
type DeleteReq struct {
	g.Meta `path:"/client/{Id}" tags:"Client" method:"delete" summary:"delete a client by id"`
	Id     string
}
type DeleteRes struct {
}
type GenerateSecretReq struct {
	g.Meta     `path:"/client/generatesecret" tags:"Client" method:"post" summary:"generate a client secret for client"`
	ClientId   int64     `json:"clientId" form:"clientId"`
	Name       string    `json:"name" form:"name"`
	Expiration time.Time `json:"expiration" form:"expiration"`
}
type GenerateSecretRes struct {
	Secret string
}
type DeleteSecretReq struct {
	g.Meta `path:"/client/deletesecret/{Id}" tags:"Client" method:"delete" summary:"delete a client secret by id"`
	Id     string
}
type DeleteSecretRes struct {
}
type GetClientSecretsReq struct {
	g.Meta   `path:"/client/getsecret" tags:"Client" method:"get" summary:"get client secrets by client id"`
	ClientId string
}
type GetClientSecretsRes struct {
	g.Meta  `mime:"json/application"`
	Secrets []*entity.ClientSecrets
}
