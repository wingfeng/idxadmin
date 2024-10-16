// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PersistedGrantsDao is the data access object for table persisted_grants.
type PersistedGrantsDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns PersistedGrantsColumns // columns contains all the column names of Table for convenient usage.
}

// PersistedGrantsColumns defines and stores column names for table persisted_grants.
type PersistedGrantsColumns struct {
	Principal string //
	ClientId  string //
	Scope     string //
	Creator   string //
	CreatorId string //
	Updator   string //
	UpdatorId string //
	CreatedAt string //
	UpdatedAt string //
}

// persistedGrantsColumns holds the columns for table persisted_grants.
var persistedGrantsColumns = PersistedGrantsColumns{
	Principal: "principal",
	ClientId:  "client_id",
	Scope:     "scope",
	Creator:   "creator",
	CreatorId: "creator_id",
	Updator:   "updator",
	UpdatorId: "updator_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPersistedGrantsDao creates and returns a new DAO object for table data access.
func NewPersistedGrantsDao() *PersistedGrantsDao {
	return &PersistedGrantsDao{
		group:   "default",
		table:   "persisted_grants",
		columns: persistedGrantsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PersistedGrantsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PersistedGrantsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PersistedGrantsDao) Columns() PersistedGrantsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PersistedGrantsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PersistedGrantsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PersistedGrantsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
