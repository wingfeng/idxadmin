// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientCorsOrigins is the golang structure of table client_cors_origins for DAO operations like Where/Data.
type ClientCorsOrigins struct {
	g.Meta    `orm:"table:client_cors_origins, do:true"`
	Origin    interface{} //
	ClientId  interface{} //
	Id        interface{} //
	Creator   interface{} //
	CreatorId interface{} //
	Updator   interface{} //
	UpdatorId interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
