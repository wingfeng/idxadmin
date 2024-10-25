// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrganizationUnits is the golang structure for table organization_units.
type OrganizationUnits struct {
	Name        string      `json:"name"        orm:"name"         description:""`
	DisplayName string      `json:"displayName" orm:"display_name" description:""`
	ParentId    ID          `json:"parentId"    orm:"parent_id"    description:""`
	SortOrder   ID          `json:"sortOrder"   orm:"sort_order"   description:""`
	Path        string      `json:"path"        orm:"path"         description:""`
	Id          ID          `json:"id"          orm:"id"           description:""`
	Creator     string      `json:"creator"     orm:"creator"      description:""`
	CreatorId   string      `json:"creatorId"   orm:"creator_id"   description:""`
	Updator     string      `json:"updator"     orm:"updator"      description:""`
	UpdatorId   string      `json:"updatorId"   orm:"updator_id"   description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""`
}
