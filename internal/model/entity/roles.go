// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// Roles is the golang structure for table roles.
type Roles struct {
	Name           string      `json:"name"           orm:"name"            description:""`
	NormalizedName string      `json:"normalizedName" orm:"normalized_name" description:""`
	Claims         *gjson.Json `json:"claims"         orm:"claims"          description:""`
	Id             int64       `json:"id"             orm:"id"              description:""`
	Creator        string      `json:"creator"        orm:"creator"         description:""`
	CreatorId      string      `json:"creatorId"      orm:"creator_id"      description:""`
	Updator        string      `json:"updator"        orm:"updator"         description:""`
	UpdatorId      string      `json:"updatorId"      orm:"updator_id"      description:""`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""`
}
