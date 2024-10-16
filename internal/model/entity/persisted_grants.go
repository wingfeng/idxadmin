// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PersistedGrants is the golang structure for table persisted_grants.
type PersistedGrants struct {
	Principal string      `json:"principal" orm:"principal"  description:""`
	ClientId  string      `json:"clientId"  orm:"client_id"  description:""`
	Scope     string      `json:"scope"     orm:"scope"      description:""`
	Creator   string      `json:"creator"   orm:"creator"    description:""`
	CreatorId string      `json:"creatorId" orm:"creator_id" description:""`
	Updator   string      `json:"updator"   orm:"updator"    description:""`
	UpdatorId string      `json:"updatorId" orm:"updator_id" description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
