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
	OuId                 int64       `json:"ouId"                 orm:"ou_id"                  description:""`
	Ou                   string      `json:"ou"                   orm:"ou"                     description:""`
	Account              string      `json:"account"              orm:"account"                description:""`
	DisplayName          string      `json:"displayName"          orm:"display_name"           description:""`
	NormalizedAccount    string      `json:"normalizedAccount"    orm:"normalized_account"     description:""`
	Email                string      `json:"email"                orm:"email"                  description:""`
	NormalizedEmail      string      `json:"normalizedEmail"      orm:"normalized_email"       description:""`
	EmailConfirmed       bool        `json:"emailConfirmed"       orm:"email_confirmed"        description:""`
	PasswordHash         string      `json:"passwordHash"         orm:"password_hash"          description:""`
	SecurityStamp        string      `json:"securityStamp"        orm:"security_stamp"         description:""`
	PhoneNumber          string      `json:"phoneNumber"          orm:"phone_number"           description:""`
	PhoneNumberConfirmed bool        `json:"phoneNumberConfirmed" orm:"phone_number_confirmed" description:""`
	TwoFactorEnabled     bool        `json:"twoFactorEnabled"     orm:"two_factor_enabled"     description:""`
	IsTemporaryPassword  bool        `json:"isTemporaryPassword"  orm:"is_temporary_password"  description:""`
	LockoutEnd           *gtime.Time `json:"lockoutEnd"           orm:"lockout_end"            description:""`
	LockoutEnabled       bool        `json:"lockoutEnabled"       orm:"lockout_enabled"        description:""`
	AccessFailedCount    int64       `json:"accessFailedCount"    orm:"access_failed_count"    description:""`
	Id                   int64       `json:"id"                   orm:"id"                     description:""`
	Creator              string      `json:"creator"              orm:"creator"                description:""`
	CreatorId            string      `json:"creatorId"            orm:"creator_id"             description:""`
	Updator              string      `json:"updator"              orm:"updator"                description:""`
	UpdatorId            string      `json:"updatorId"            orm:"updator_id"             description:""`
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"             description:""`
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"             description:""`
	Claims               *gjson.Json `json:"claims"               orm:"claims"                 description:""`
}
