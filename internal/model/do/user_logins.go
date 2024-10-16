// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLogins is the golang structure of table user_logins for DAO operations like Where/Data.
type UserLogins struct {
	g.Meta              `orm:"table:user_logins, do:true"`
	LoginProvider       interface{} //
	ProviderKey         interface{} //
	ProviderDisplayName interface{} //
	UserId              interface{} //
	Creator             interface{} //
	CreatorId           interface{} //
	Updator             interface{} //
	UpdatorId           interface{} //
	CreatedAt           *gtime.Time //
	UpdatedAt           *gtime.Time //
}
