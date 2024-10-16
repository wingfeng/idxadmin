// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientSecrets is the golang structure of table client_secrets for DAO operations like Where/Data.
type ClientSecrets struct {
	g.Meta     `orm:"table:client_secrets, do:true"`
	Name       interface{} //
	Value      interface{} //
	Expiration *gtime.Time //
	ClientId   interface{} //
	Id         interface{} //
	Creator    interface{} //
	CreatorId  interface{} //
	Updator    interface{} //
	UpdatorId  interface{} //
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
