// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserTokensDao is the data access object for table user_tokens.
type UserTokensDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns UserTokensColumns // columns contains all the column names of Table for convenient usage.
}

// UserTokensColumns defines and stores column names for table user_tokens.
type UserTokensColumns struct {
	LoginProvider string //
	Name          string //
	Value         string //
	Creator       string //
	CreatorId     string //
	Updator       string //
	UpdatorId     string //
	CreatedAt     string //
	UpdatedAt     string //
}

// userTokensColumns holds the columns for table user_tokens.
var userTokensColumns = UserTokensColumns{
	LoginProvider: "login_provider",
	Name:          "name",
	Value:         "value",
	Creator:       "creator",
	CreatorId:     "creator_id",
	Updator:       "updator",
	UpdatorId:     "updator_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewUserTokensDao creates and returns a new DAO object for table data access.
func NewUserTokensDao() *UserTokensDao {
	return &UserTokensDao{
		group:   "default",
		table:   "user_tokens",
		columns: userTokensColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserTokensDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserTokensDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserTokensDao) Columns() UserTokensColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserTokensDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserTokensDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserTokensDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
