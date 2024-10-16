// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClientsDao is the data access object for table clients.
type ClientsDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ClientsColumns // columns contains all the column names of Table for convenient usage.
}

// ClientsColumns defines and stores column names for table clients.
type ClientsColumns struct {
	Enabled                           string //
	ClientId                          string //
	GrantTypes                        string //
	Scopes                            string //
	RedirectUris                      string //
	ProtocolType                      string //
	RequireSecret                     string //
	ClientName                        string //
	Description                       string //
	ClientUri                         string //
	LogoUri                           string //
	RequireConsent                    string //
	AllowRememberConsent              string //
	AlwaysIncludeUserClaimsInIdToken  string //
	RequirePkce                       string //
	AllowPlainTextPkce                string //
	AllowAccessTokensViaBrowser       string //
	FrontChannelLogoutUri             string //
	FrontChannelLogoutSessionRequired string //
	BackChannelLogoutUri              string //
	BackChannelLogoutSessionRequired  string //
	AllowOfflineAccess                string //
	IdentityTokenLifetime             string //
	AccessTokenLifetime               string //
	AuthorizationCodeLifetime         string //
	ConsentLifetime                   string //
	AbsoluteRefreshTokenLifetime      string //
	SlidingRefreshTokenLifetime       string //
	RefreshTokenUsage                 string //
	UpdateAccessTokenClaimsOnRefresh  string //
	RefreshTokenExpiration            string //
	EnableLocalLogin                  string //
	AlwaysSendClientClaims            string //
	ClientClaimsPrefix                string //
	PairWiseSubjectSalt               string //
	UserSsoLifetime                   string //
	Claims                            string //
	Properties                        string //
	Id                                string //
	Creator                           string //
	CreatorId                         string //
	Updator                           string //
	UpdatorId                         string //
	CreatedAt                         string //
	UpdatedAt                         string //
}

// clientsColumns holds the columns for table clients.
var clientsColumns = ClientsColumns{
	Enabled:                           "enabled",
	ClientId:                          "client_id",
	GrantTypes:                        "grant_types",
	Scopes:                            "scopes",
	RedirectUris:                      "redirect_uris",
	ProtocolType:                      "protocol_type",
	RequireSecret:                     "require_secret",
	ClientName:                        "client_name",
	Description:                       "description",
	ClientUri:                         "client_uri",
	LogoUri:                           "logo_uri",
	RequireConsent:                    "require_consent",
	AllowRememberConsent:              "allow_remember_consent",
	AlwaysIncludeUserClaimsInIdToken:  "always_include_user_claims_in_id_token",
	RequirePkce:                       "require_pkce",
	AllowPlainTextPkce:                "allow_plain_text_pkce",
	AllowAccessTokensViaBrowser:       "allow_access_tokens_via_browser",
	FrontChannelLogoutUri:             "front_channel_logout_uri",
	FrontChannelLogoutSessionRequired: "front_channel_logout_session_required",
	BackChannelLogoutUri:              "back_channel_logout_uri",
	BackChannelLogoutSessionRequired:  "back_channel_logout_session_required",
	AllowOfflineAccess:                "allow_offline_access",
	IdentityTokenLifetime:             "identity_token_lifetime",
	AccessTokenLifetime:               "access_token_lifetime",
	AuthorizationCodeLifetime:         "authorization_code_lifetime",
	ConsentLifetime:                   "consent_lifetime",
	AbsoluteRefreshTokenLifetime:      "absolute_refresh_token_lifetime",
	SlidingRefreshTokenLifetime:       "sliding_refresh_token_lifetime",
	RefreshTokenUsage:                 "refresh_token_usage",
	UpdateAccessTokenClaimsOnRefresh:  "update_access_token_claims_on_refresh",
	RefreshTokenExpiration:            "refresh_token_expiration",
	EnableLocalLogin:                  "enable_local_login",
	AlwaysSendClientClaims:            "always_send_client_claims",
	ClientClaimsPrefix:                "client_claims_prefix",
	PairWiseSubjectSalt:               "pair_wise_subject_salt",
	UserSsoLifetime:                   "user_sso_lifetime",
	Claims:                            "claims",
	Properties:                        "properties",
	Id:                                "id",
	Creator:                           "creator",
	CreatorId:                         "creator_id",
	Updator:                           "updator",
	UpdatorId:                         "updator_id",
	CreatedAt:                         "created_at",
	UpdatedAt:                         "updated_at",
}

// NewClientsDao creates and returns a new DAO object for table data access.
func NewClientsDao() *ClientsDao {
	return &ClientsDao{
		group:   "default",
		table:   "clients",
		columns: clientsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ClientsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ClientsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ClientsDao) Columns() ClientsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ClientsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ClientsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ClientsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
