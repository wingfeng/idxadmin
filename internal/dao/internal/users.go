// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDao is the data access object for table users.
type UsersDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns UsersColumns // columns contains all the column names of Table for convenient usage.
}

// UsersColumns defines and stores column names for table users.
type UsersColumns struct {
	OuId                 string //
	Ou                   string //
	UserName             string //
	DisplayName          string //
	Email                string //
	EmailConfirmed       string //
	PasswordHash         string //
	PhoneNumber          string //
	PhoneNumberConfirmed string //
	TwoFactorEnabled     string //
	IsTemporaryPassword  string //
	LockoutEnd           string //
	LockoutEnabled       string //
	AccessFailedCount    string //
	Claims               string //
	Id                   string //
	Creator              string //
	CreatorId            string //
	Updator              string //
	UpdatorId            string //
	CreatedAt            string //
	UpdatedAt            string //
}

// usersColumns holds the columns for table users.
var usersColumns = UsersColumns{
	OuId:                 "ou_id",
	Ou:                   "ou",
	UserName:             "user_name",
	DisplayName:          "display_name",
	Email:                "email",
	EmailConfirmed:       "email_confirmed",
	PasswordHash:         "password_hash",
	PhoneNumber:          "phone_number",
	PhoneNumberConfirmed: "phone_number_confirmed",
	TwoFactorEnabled:     "two_factor_enabled",
	IsTemporaryPassword:  "is_temporary_password",
	LockoutEnd:           "lockout_end",
	LockoutEnabled:       "lockout_enabled",
	AccessFailedCount:    "access_failed_count",
	Claims:               "claims",
	Id:                   "id",
	Creator:              "creator",
	CreatorId:            "creator_id",
	Updator:              "updator",
	UpdatorId:            "updator_id",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao() *UsersDao {
	return &UsersDao{
		group:   "default",
		table:   "users",
		columns: usersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UsersDao) Columns() UsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
