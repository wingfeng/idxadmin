// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserTokens is the golang structure of table user_tokens for DAO operations like Where/Data.
type UserTokens struct {
	g.Meta        `orm:"table:user_tokens, do:true"`
	LoginProvider interface{} //
	Name          interface{} //
	Value         interface{} //
	Creator       interface{} //
	CreatorId     interface{} //
	Updator       interface{} //
	UpdatorId     interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
