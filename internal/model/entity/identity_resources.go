// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// IdentityResources is the golang structure for table identity_resources.
type IdentityResources struct {
	Enabled                 bool        `json:"enabled"                 orm:"enabled"                    description:""`
	Name                    string      `json:"name"                    orm:"name"                       description:""`
	DisplayName             string      `json:"displayName"             orm:"display_name"               description:""`
	Description             string      `json:"description"             orm:"description"                description:""`
	Required                bool        `json:"required"                orm:"required"                   description:""`
	Emphasize               bool        `json:"emphasize"               orm:"emphasize"                  description:""`
	ShowInDiscoveryDocument bool        `json:"showInDiscoveryDocument" orm:"show_in_discovery_document" description:""`
	Id                      int64       `json:"id"                      orm:"id"                         description:""`
	Creator                 string      `json:"creator"                 orm:"creator"                    description:""`
	CreatorId               string      `json:"creatorId"               orm:"creator_id"                 description:""`
	Updator                 string      `json:"updator"                 orm:"updator"                    description:""`
	UpdatorId               string      `json:"updatorId"               orm:"updator_id"                 description:""`
	CreatedAt               *gtime.Time `json:"createdAt"               orm:"created_at"                 description:""`
	UpdatedAt               *gtime.Time `json:"updatedAt"               orm:"updated_at"                 description:""`
}
