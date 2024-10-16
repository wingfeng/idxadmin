// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PersistedGrants is the golang structure of table persisted_grants for DAO operations like Where/Data.
type PersistedGrants struct {
	g.Meta    `orm:"table:persisted_grants, do:true"`
	Principal interface{} //
	ClientId  interface{} //
	Scope     interface{} //
	Creator   interface{} //
	CreatorId interface{} //
	Updator   interface{} //
	UpdatorId interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
