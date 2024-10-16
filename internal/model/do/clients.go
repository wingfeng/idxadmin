// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Clients is the golang structure of table clients for DAO operations like Where/Data.
type Clients struct {
	g.Meta                            `orm:"table:clients, do:true"`
	Enabled                           interface{} //
	ClientId                          interface{} //
	GrantTypes                        interface{} //
	Scopes                            interface{} //
	RedirectUris                      interface{} //
	ProtocolType                      interface{} //
	RequireSecret                     interface{} //
	ClientName                        interface{} //
	Description                       interface{} //
	ClientUri                         interface{} //
	LogoUri                           interface{} //
	RequireConsent                    interface{} //
	AllowRememberConsent              interface{} //
	AlwaysIncludeUserClaimsInIdToken  interface{} //
	RequirePkce                       interface{} //
	AllowPlainTextPkce                interface{} //
	AllowAccessTokensViaBrowser       interface{} //
	FrontChannelLogoutUri             interface{} //
	FrontChannelLogoutSessionRequired interface{} //
	BackChannelLogoutUri              interface{} //
	BackChannelLogoutSessionRequired  interface{} //
	AllowOfflineAccess                interface{} //
	IdentityTokenLifetime             interface{} //
	AccessTokenLifetime               interface{} //
	AuthorizationCodeLifetime         interface{} //
	ConsentLifetime                   interface{} //
	AbsoluteRefreshTokenLifetime      interface{} //
	SlidingRefreshTokenLifetime       interface{} //
	RefreshTokenUsage                 interface{} //
	UpdateAccessTokenClaimsOnRefresh  interface{} //
	RefreshTokenExpiration            interface{} //
	EnableLocalLogin                  interface{} //
	AlwaysSendClientClaims            interface{} //
	ClientClaimsPrefix                interface{} //
	PairWiseSubjectSalt               interface{} //
	UserSsoLifetime                   interface{} //
	Claims                            *gjson.Json //
	Properties                        *gjson.Json //
	Id                                interface{} //
	Creator                           interface{} //
	CreatorId                         interface{} //
	Updator                           interface{} //
	UpdatorId                         interface{} //
	CreatedAt                         *gtime.Time //
	UpdatedAt                         *gtime.Time //
}
