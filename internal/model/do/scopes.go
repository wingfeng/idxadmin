// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Scopes is the golang structure of table scopes for DAO operations like Where/Data.
type Scopes struct {
	g.Meta      `orm:"table:scopes, do:true"`
	Enabled     interface{} // enabled
	Name        interface{} //
	DisplayName interface{} //
	Description interface{} //
	Claims      *gjson.Json //
	Properties  *gjson.Json //
	Id          interface{} //
	Creator     interface{} //
	CreatorId   interface{} //
	Updator     interface{} //
	UpdatorId   interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
