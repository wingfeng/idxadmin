// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GroupsDao is the data access object for table groups.
type GroupsDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns GroupsColumns // columns contains all the column names of Table for convenient usage.
}

// GroupsColumns defines and stores column names for table groups.
type GroupsColumns struct {
	Id          string //
	Creator     string //
	CreatorId   string //
	Updator     string //
	UpdatorId   string //
	CreatedAt   string //
	UpdatedAt   string //
	Name        string //
	Description string //
}

// groupsColumns holds the columns for table groups.
var groupsColumns = GroupsColumns{
	Id:          "id",
	Creator:     "creator",
	CreatorId:   "creator_id",
	Updator:     "updator",
	UpdatorId:   "updator_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Name:        "name",
	Description: "description",
}

// NewGroupsDao creates and returns a new DAO object for table data access.
func NewGroupsDao() *GroupsDao {
	return &GroupsDao{
		group:   "default",
		table:   "groups",
		columns: groupsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GroupsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GroupsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GroupsDao) Columns() GroupsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GroupsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GroupsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GroupsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
