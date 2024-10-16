// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserTokens is the golang structure for table user_tokens.
type UserTokens struct {
	LoginProvider string      `json:"loginProvider" orm:"login_provider" description:""`
	Name          string      `json:"name"          orm:"name"           description:""`
	Value         string      `json:"value"         orm:"value"          description:""`
	Creator       string      `json:"creator"       orm:"creator"        description:""`
	CreatorId     string      `json:"creatorId"     orm:"creator_id"     description:""`
	Updator       string      `json:"updator"       orm:"updator"        description:""`
	UpdatorId     string      `json:"updatorId"     orm:"updator_id"     description:""`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:""`
}
