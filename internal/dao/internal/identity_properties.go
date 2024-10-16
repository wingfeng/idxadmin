// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// IdentityPropertiesDao is the data access object for table identity_properties.
type IdentityPropertiesDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns IdentityPropertiesColumns // columns contains all the column names of Table for convenient usage.
}

// IdentityPropertiesColumns defines and stores column names for table identity_properties.
type IdentityPropertiesColumns struct {
	Key                string //
	Value              string //
	IdentityResourceId string //
	Id                 string //
	Creator            string //
	CreatorId          string //
	Updator            string //
	UpdatorId          string //
	CreatedAt          string //
	UpdatedAt          string //
}

// identityPropertiesColumns holds the columns for table identity_properties.
var identityPropertiesColumns = IdentityPropertiesColumns{
	Key:                "key",
	Value:              "value",
	IdentityResourceId: "identity_resource_id",
	Id:                 "id",
	Creator:            "creator",
	CreatorId:          "creator_id",
	Updator:            "updator",
	UpdatorId:          "updator_id",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewIdentityPropertiesDao creates and returns a new DAO object for table data access.
func NewIdentityPropertiesDao() *IdentityPropertiesDao {
	return &IdentityPropertiesDao{
		group:   "default",
		table:   "identity_properties",
		columns: identityPropertiesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *IdentityPropertiesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *IdentityPropertiesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *IdentityPropertiesDao) Columns() IdentityPropertiesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *IdentityPropertiesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *IdentityPropertiesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *IdentityPropertiesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
