// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// Scopes is the golang structure for table scopes.
type Scopes struct {
	Enabled     bool        `json:"enabled"     orm:"enabled"     description:""`
	Name        string      `json:"name"        orm:"name"        description:""`
	Description string      `json:"description" orm:"description" description:""`
	Properties  *gjson.Json `json:"properties"  orm:"properties"  description:""`
	Id          ID          `json:"id"          orm:"id"          description:""`
	Creator     string      `json:"creator"     orm:"creator"     description:""`
	CreatorId   string      `json:"creatorId"   orm:"creator_id"  description:""`
	Updator     string      `json:"updator"     orm:"updator"     description:""`
	UpdatorId   string      `json:"updatorId"   orm:"updator_id"  description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
}
