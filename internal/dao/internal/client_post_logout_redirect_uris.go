// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClientPostLogoutRedirectUrisDao is the data access object for table client_post_logout_redirect_uris.
type ClientPostLogoutRedirectUrisDao struct {
	table   string                              // table is the underlying table name of the DAO.
	group   string                              // group is the database configuration group name of current DAO.
	columns ClientPostLogoutRedirectUrisColumns // columns contains all the column names of Table for convenient usage.
}

// ClientPostLogoutRedirectUrisColumns defines and stores column names for table client_post_logout_redirect_uris.
type ClientPostLogoutRedirectUrisColumns struct {
	PostLogoutRedirectUri string //
	ClientId              string //
	Id                    string //
	Creator               string //
	CreatorId             string //
	Updator               string //
	UpdatorId             string //
	CreatedAt             string //
	UpdatedAt             string //
}

// clientPostLogoutRedirectUrisColumns holds the columns for table client_post_logout_redirect_uris.
var clientPostLogoutRedirectUrisColumns = ClientPostLogoutRedirectUrisColumns{
	PostLogoutRedirectUri: "post_logout_redirect_uri",
	ClientId:              "client_id",
	Id:                    "id",
	Creator:               "creator",
	CreatorId:             "creator_id",
	Updator:               "updator",
	UpdatorId:             "updator_id",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
}

// NewClientPostLogoutRedirectUrisDao creates and returns a new DAO object for table data access.
func NewClientPostLogoutRedirectUrisDao() *ClientPostLogoutRedirectUrisDao {
	return &ClientPostLogoutRedirectUrisDao{
		group:   "default",
		table:   "client_post_logout_redirect_uris",
		columns: clientPostLogoutRedirectUrisColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ClientPostLogoutRedirectUrisDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ClientPostLogoutRedirectUrisDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ClientPostLogoutRedirectUrisDao) Columns() ClientPostLogoutRedirectUrisColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ClientPostLogoutRedirectUrisDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ClientPostLogoutRedirectUrisDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ClientPostLogoutRedirectUrisDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
