// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	OuId                 ID          `json:"ouId"                 orm:"ou_id"                  description:""`
	Ou                   string      `json:"ou"                   orm:"ou"                     description:""`
	UserName             string      `json:"userName"             orm:"user_name"              description:""`
	DisplayName          string      `json:"displayName"          orm:"display_name"           description:""`
	Email                string      `json:"email"                orm:"email"                  description:""`
	EmailConfirmed       bool        `json:"emailConfirmed"       orm:"email_confirmed"        description:""`
	PasswordHash         string      `json:"-"         orm:"password_hash"          description:""`
	PhoneNumber          string      `json:"phoneNumber"          orm:"phone_number"           description:""`
	PhoneNumberConfirmed bool        `json:"phoneNumberConfirmed" orm:"phone_number_confirmed" description:""`
	TwoFactorEnabled     bool        `json:"twoFactorEnabled"     orm:"two_factor_enabled"     description:""`
	IsTemporaryPassword  bool        `json:"isTemporaryPassword"  orm:"is_temporary_password"  description:""`
	LockoutEnd           *gtime.Time `json:"lockoutEnd"           orm:"lockout_end"            description:""`
	LockoutEnabled       bool        `json:"lockoutEnabled"       orm:"lockout_enabled"        description:""`
	AccessFailedCount    ID          `json:"accessFailedCount"    orm:"access_failed_count"    description:""`
	Claims               *gjson.Json `json:"claims"               orm:"claims"                 description:""`
	Id                   ID          `json:"id"                   orm:"id"                     description:""`
	Creator              string      `json:"creator"              orm:"creator"                description:""`
	CreatorId            string      `json:"creatorId"            orm:"creator_id"             description:""`
	Updator              string      `json:"updator"              orm:"updator"                description:""`
	UpdatorId            string      `json:"updatorId"            orm:"updator_id"             description:""`
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"             description:""`
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"             description:""`
}
