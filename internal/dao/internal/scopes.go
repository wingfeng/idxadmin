// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ScopesDao is the data access object for table scopes.
type ScopesDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ScopesColumns // columns contains all the column names of Table for convenient usage.
}

// ScopesColumns defines and stores column names for table scopes.
type ScopesColumns struct {
	Enabled     string // enabled
	Name        string //
	DisplayName string //
	Description string //
	Claims      string //
	Properties  string //
	Id          string //
	Creator     string //
	CreatorId   string //
	Updator     string //
	UpdatorId   string //
	CreatedAt   string //
	UpdatedAt   string //
}

// scopesColumns holds the columns for table scopes.
var scopesColumns = ScopesColumns{
	Enabled:     "enabled",
	Name:        "name",
	DisplayName: "display_name",
	Description: "description",
	Claims:      "claims",
	Properties:  "properties",
	Id:          "id",
	Creator:     "creator",
	CreatorId:   "creator_id",
	Updator:     "updator",
	UpdatorId:   "updator_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewScopesDao creates and returns a new DAO object for table data access.
func NewScopesDao() *ScopesDao {
	return &ScopesDao{
		group:   "default",
		table:   "scopes",
		columns: scopesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ScopesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ScopesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ScopesDao) Columns() ScopesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ScopesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ScopesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ScopesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
