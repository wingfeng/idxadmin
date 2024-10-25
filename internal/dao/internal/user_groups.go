// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserGroupsDao is the data access object for table user_groups.
type UserGroupsDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns UserGroupsColumns // columns contains all the column names of Table for convenient usage.
}

// UserGroupsColumns defines and stores column names for table user_groups.
type UserGroupsColumns struct {
	UserId  string //
	GroupId string //
}

// userGroupsColumns holds the columns for table user_groups.
var userGroupsColumns = UserGroupsColumns{
	UserId:  "user_id",
	GroupId: "group_id",
}

// NewUserGroupsDao creates and returns a new DAO object for table data access.
func NewUserGroupsDao() *UserGroupsDao {
	return &UserGroupsDao{
		group:   "default",
		table:   "user_groups",
		columns: userGroupsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserGroupsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserGroupsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserGroupsDao) Columns() UserGroupsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserGroupsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserGroupsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserGroupsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
