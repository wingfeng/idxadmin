// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLogins is the golang structure for table user_logins.
type UserLogins struct {
	LoginProvider       string      `json:"loginProvider"       orm:"login_provider"        description:""`
	ProviderKey         string      `json:"providerKey"         orm:"provider_key"          description:""`
	ProviderDisplayName string      `json:"providerDisplayName" orm:"provider_display_name" description:""`
	UserId              int64       `json:"userId"              orm:"user_id"               description:""`
	Creator             string      `json:"creator"             orm:"creator"               description:""`
	CreatorId           string      `json:"creatorId"           orm:"creator_id"            description:""`
	Updator             string      `json:"updator"             orm:"updator"               description:""`
	UpdatorId           string      `json:"updatorId"           orm:"updator_id"            description:""`
	CreatedAt           *gtime.Time `json:"createdAt"           orm:"created_at"            description:""`
	UpdatedAt           *gtime.Time `json:"updatedAt"           orm:"updated_at"            description:""`
}
