// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrganizationUnits is the golang structure of table organization_units for DAO operations like Where/Data.
type OrganizationUnits struct {
	g.Meta      `orm:"table:organization_units, do:true"`
	Name        interface{} //
	DisplayName interface{} //
	ParentId    interface{} //
	SortOrder   interface{} //
	Path        interface{} //
	Id          interface{} //
	Creator     interface{} //
	CreatorId   interface{} //
	Updator     interface{} //
	UpdatorId   interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
