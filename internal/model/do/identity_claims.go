// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// IdentityClaims is the golang structure of table identity_claims for DAO operations like Where/Data.
type IdentityClaims struct {
	g.Meta             `orm:"table:identity_claims, do:true"`
	Type               interface{} //
	Identityresourceid interface{} //
	Id                 interface{} //
	Creator            interface{} //
	CreatorId          interface{} //
	Updator            interface{} //
	UpdatorId          interface{} //
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
}
