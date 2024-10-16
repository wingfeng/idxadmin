// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta               `orm:"table:users, do:true"`
	OuId                 interface{} //
	Ou                   interface{} //
	Account              interface{} //
	DisplayName          interface{} //
	NormalizedAccount    interface{} //
	Email                interface{} //
	NormalizedEmail      interface{} //
	EmailConfirmed       interface{} //
	PasswordHash         interface{} //
	SecurityStamp        interface{} //
	PhoneNumber          interface{} //
	PhoneNumberConfirmed interface{} //
	TwoFactorEnabled     interface{} //
	IsTemporaryPassword  interface{} //
	LockoutEnd           *gtime.Time //
	LockoutEnabled       interface{} //
	AccessFailedCount    interface{} //
	Id                   interface{} //
	Creator              interface{} //
	CreatorId            interface{} //
	Updator              interface{} //
	UpdatorId            interface{} //
	CreatedAt            *gtime.Time //
	UpdatedAt            *gtime.Time //
	Claims               *gjson.Json //
}
