// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// IdentityClaims is the golang structure for table identity_claims.
type IdentityClaims struct {
	Type               string      `json:"type"               orm:"type"               description:""`
	Identityresourceid int64       `json:"identityresourceid" orm:"identityresourceid" description:""`
	Id                 int64       `json:"id"                 orm:"id"                 description:""`
	Creator            string      `json:"creator"            orm:"creator"            description:""`
	CreatorId          string      `json:"creatorId"          orm:"creator_id"         description:""`
	Updator            string      `json:"updator"            orm:"updator"            description:""`
	UpdatorId          string      `json:"updatorId"          orm:"updator_id"         description:""`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"         description:""`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"         description:""`
}
