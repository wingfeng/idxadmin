// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// IdentityProperties is the golang structure of table identity_properties for DAO operations like Where/Data.
type IdentityProperties struct {
	g.Meta             `orm:"table:identity_properties, do:true"`
	Key                interface{} //
	Value              interface{} //
	IdentityResourceId interface{} //
	Id                 interface{} //
	Creator            interface{} //
	CreatorId          interface{} //
	Updator            interface{} //
	UpdatorId          interface{} //
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
}
