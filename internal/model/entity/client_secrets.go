// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientSecrets is the golang structure for table client_secrets.
type ClientSecrets struct {
	Name       string      `json:"name"       orm:"name"       description:""`
	Value      string      `json:"value"      orm:"value"      description:""`
	Expiration *gtime.Time `json:"expiration" orm:"expiration" description:""`
	ClientId   int64       `json:"clientId"   orm:"client_id"  description:""`
	Id         int64       `json:"id"         orm:"id"         description:""`
	Creator    string      `json:"creator"    orm:"creator"    description:""`
	CreatorId  string      `json:"creatorId"  orm:"creator_id" description:""`
	Updator    string      `json:"updator"    orm:"updator"    description:""`
	UpdatorId  string      `json:"updatorId"  orm:"updator_id" description:""`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at" description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at" description:""`
}
