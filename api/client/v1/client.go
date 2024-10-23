package v1

import (
	"time"

	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetReq struct {
	g.Meta `path:"/client/{Id}" tags:"Client" method:"get" summary:"get client info by id"`
	Id     int `v:"required"`
}
type GetRes struct {
	g.Meta `mime:"json/application"`
	entity.Clients
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
	entity.Clients
}
type SaveRes struct {
}
type DeleteReq struct {
	g.Meta `path:"/client/{Id}" tags:"Client" method:"delete" summary:"delete a client by id"`
	Id     string `v:"required"`
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
	g.Meta `path:"/client/deletesecret" tags:"Client" method:"delete" summary:"delete a client secret by id"`
	Id     string `v:"required"`
}
type DeleteSecretRes struct {
}
type GetClientSecretsReq struct {
	g.Meta   `path:"/client/getsecret" tags:"Client" method:"get" summary:"get client secrets by client id"`
	ClientId string `v:"required"`
}
type GetClientSecretsRes struct {
	g.Meta  `mime:"json/application"`
	Secrets []*entity.ClientSecrets
}
