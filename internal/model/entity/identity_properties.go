// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// IdentityProperties is the golang structure for table identity_properties.
type IdentityProperties struct {
	Key                string      `json:"key"                orm:"key"                  description:""`
	Value              string      `json:"value"              orm:"value"                description:""`
	IdentityResourceId int64       `json:"identityResourceId" orm:"identity_resource_id" description:""`
	Id                 int64       `json:"id"                 orm:"id"                   description:""`
	Creator            string      `json:"creator"            orm:"creator"              description:""`
	CreatorId          string      `json:"creatorId"          orm:"creator_id"           description:""`
	Updator            string      `json:"updator"            orm:"updator"              description:""`
	UpdatorId          string      `json:"updatorId"          orm:"updator_id"           description:""`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:""`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:""`
}
