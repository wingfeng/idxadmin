// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Groups is the golang structure of table groups for DAO operations like Where/Data.
type Groups struct {
	g.Meta      `orm:"table:groups, do:true"`
	Id          interface{} //
	Creator     interface{} //
	CreatorId   interface{} //
	Updator     interface{} //
	UpdatorId   interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	Name        interface{} //
	Description interface{} //
}
