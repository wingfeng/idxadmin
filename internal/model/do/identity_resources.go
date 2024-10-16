// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// IdentityResources is the golang structure of table identity_resources for DAO operations like Where/Data.
type IdentityResources struct {
	g.Meta                  `orm:"table:identity_resources, do:true"`
	Enabled                 interface{} //
	Name                    interface{} //
	DisplayName             interface{} //
	Description             interface{} //
	Required                interface{} //
	Emphasize               interface{} //
	ShowInDiscoveryDocument interface{} //
	Id                      interface{} //
	Creator                 interface{} //
	CreatorId               interface{} //
	Updator                 interface{} //
	UpdatorId               interface{} //
	CreatedAt               *gtime.Time //
	UpdatedAt               *gtime.Time //
}
