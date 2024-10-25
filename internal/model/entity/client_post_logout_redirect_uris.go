// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientPostLogoutRedirectUris is the golang structure for table client_post_logout_redirect_uris.
type ClientPostLogoutRedirectUris struct {
	PostLogoutRedirectUri string      `json:"postLogoutRedirectUri" orm:"post_logout_redirect_uri" description:""`
	ClientId              ID          `json:"clientId"              orm:"client_id"                description:""`
	Id                    ID          `json:"id"                    orm:"id"                       description:""`
	Creator               string      `json:"creator"               orm:"creator"                  description:""`
	CreatorId             string      `json:"creatorId"             orm:"creator_id"               description:""`
	Updator               string      `json:"updator"               orm:"updator"                  description:""`
	UpdatorId             string      `json:"updatorId"             orm:"updator_id"               description:""`
	CreatedAt             *gtime.Time `json:"createdAt"             orm:"created_at"               description:""`
	UpdatedAt             *gtime.Time `json:"updatedAt"             orm:"updated_at"               description:""`
}
