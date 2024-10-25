// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// Clients is the golang structure for table clients.
type Clients struct {
	Enabled                           bool        `json:"enabled"                           orm:"enabled"                                description:""`
	ClientId                          string      `json:"clientId"                          orm:"client_id"                              description:""`
	GrantTypes                        string      `json:"grantTypes"                        orm:"grant_types"                            description:""`
	Scopes                            string      `json:"scopes"                            orm:"scopes"                                 description:""`
	RedirectUris                      string      `json:"redirectUris"                      orm:"redirect_uris"                          description:""`
	ProtocolType                      string      `json:"protocolType"                      orm:"protocol_type"                          description:""`
	RequireSecret                     bool        `json:"requireSecret"                     orm:"require_secret"                         description:""`
	ClientName                        string      `json:"clientName"                        orm:"client_name"                            description:""`
	Description                       string      `json:"description"                       orm:"description"                            description:""`
	ClientUri                         string      `json:"clientUri"                         orm:"client_uri"                             description:""`
	LogoUri                           string      `json:"logoUri"                           orm:"logo_uri"                               description:""`
	RequireConsent                    bool        `json:"requireConsent"                    orm:"require_consent"                        description:""`
	AllowRememberConsent              bool        `json:"allowRememberConsent"              orm:"allow_remember_consent"                 description:""`
	AlwaysIncludeUserClaimsInIdToken  bool        `json:"alwaysIncludeUserClaimsInIdToken"  orm:"always_include_user_claims_in_id_token" description:""`
	RequirePkce                       bool        `json:"requirePkce"                       orm:"require_pkce"                           description:""`
	AllowPlainTextPkce                bool        `json:"allowPlainTextPkce"                orm:"allow_plain_text_pkce"                  description:""`
	AllowAccessTokensViaBrowser       bool        `json:"allowAccessTokensViaBrowser"       orm:"allow_access_tokens_via_browser"        description:""`
	FrontChannelLogoutUri             string      `json:"frontChannelLogoutUri"             orm:"front_channel_logout_uri"               description:""`
	FrontChannelLogoutSessionRequired bool        `json:"frontChannelLogoutSessionRequired" orm:"front_channel_logout_session_required"  description:""`
	BackChannelLogoutUri              string      `json:"backChannelLogoutUri"              orm:"back_channel_logout_uri"                description:""`
	BackChannelLogoutSessionRequired  bool        `json:"backChannelLogoutSessionRequired"  orm:"back_channel_logout_session_required"   description:""`
	AllowOfflineAccess                bool        `json:"allowOfflineAccess"                orm:"allow_offline_access"                   description:""`
	IdentityTokenLifetime             ID          `json:"identityTokenLifetime"             orm:"identity_token_lifetime"                description:""`
	AccessTokenLifetime               ID          `json:"accessTokenLifetime"               orm:"access_token_lifetime"                  description:""`
	AuthorizationCodeLifetime         ID          `json:"authorizationCodeLifetime"         orm:"authorization_code_lifetime"            description:""`
	ConsentLifetime                   ID          `json:"consentLifetime"                   orm:"consent_lifetime"                       description:""`
	AbsoluteRefreshTokenLifetime      ID          `json:"absoluteRefreshTokenLifetime"      orm:"absolute_refresh_token_lifetime"        description:""`
	SlidingRefreshTokenLifetime       ID          `json:"slidingRefreshTokenLifetime"       orm:"sliding_refresh_token_lifetime"         description:""`
	RefreshTokenUsage                 ID          `json:"refreshTokenUsage"                 orm:"refresh_token_usage"                    description:""`
	UpdateAccessTokenClaimsOnRefresh  bool        `json:"updateAccessTokenClaimsOnRefresh"  orm:"update_access_token_claims_on_refresh"  description:""`
	RefreshTokenExpiration            ID          `json:"refreshTokenExpiration"            orm:"refresh_token_expiration"               description:""`
	EnableLocalLogin                  bool        `json:"enableLocalLogin"                  orm:"enable_local_login"                     description:""`
	AlwaysSendClientClaims            bool        `json:"alwaysSendClientClaims"            orm:"always_send_client_claims"              description:""`
	ClientClaimsPrefix                string      `json:"clientClaimsPrefix"                orm:"client_claims_prefix"                   description:""`
	PairWiseSubjectSalt               string      `json:"pairWiseSubjectSalt"               orm:"pair_wise_subject_salt"                 description:""`
	UserSsoLifetime                   ID          `json:"userSsoLifetime"                   orm:"user_sso_lifetime"                      description:""`
	Claims                            *gjson.Json `json:"claims"                            orm:"claims"                                 description:""`
	Properties                        *gjson.Json `json:"properties"                        orm:"properties"                             description:""`
	Id                                ID          `json:"id"                                orm:"id"                                     description:""`
	Creator                           string      `json:"creator"                           orm:"creator"                                description:""`
	CreatorId                         string      `json:"creatorId"                         orm:"creator_id"                             description:""`
	Updator                           string      `json:"updator"                           orm:"updator"                                description:""`
	UpdatorId                         string      `json:"updatorId"                         orm:"updator_id"                             description:""`
	CreatedAt                         *gtime.Time `json:"createdAt"                         orm:"created_at"                             description:""`
	UpdatedAt                         *gtime.Time `json:"updatedAt"                         orm:"updated_at"                             description:""`
}
