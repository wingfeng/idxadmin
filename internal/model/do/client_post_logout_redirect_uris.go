// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientPostLogoutRedirectUris is the golang structure of table client_post_logout_redirect_uris for DAO operations like Where/Data.
type ClientPostLogoutRedirectUris struct {
	g.Meta                `orm:"table:client_post_logout_redirect_uris, do:true"`
	PostLogoutRedirectUri interface{} //
	ClientId              interface{} //
	Id                    interface{} //
	Creator               interface{} //
	CreatorId             interface{} //
	Updator               interface{} //
	UpdatorId             interface{} //
	CreatedAt             *gtime.Time //
	UpdatedAt             *gtime.Time //
}
