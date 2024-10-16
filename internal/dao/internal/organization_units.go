// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrganizationUnitsDao is the data access object for table organization_units.
type OrganizationUnitsDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns OrganizationUnitsColumns // columns contains all the column names of Table for convenient usage.
}

// OrganizationUnitsColumns defines and stores column names for table organization_units.
type OrganizationUnitsColumns struct {
	Name        string //
	DisplayName string //
	ParentId    string //
	SortOrder   string //
	Path        string //
	Id          string //
	Creator     string //
	CreatorId   string //
	Updator     string //
	UpdatorId   string //
	CreatedAt   string //
	UpdatedAt   string //
}

// organizationUnitsColumns holds the columns for table organization_units.
var organizationUnitsColumns = OrganizationUnitsColumns{
	Name:        "name",
	DisplayName: "display_name",
	ParentId:    "parent_id",
	SortOrder:   "sort_order",
	Path:        "path",
	Id:          "id",
	Creator:     "creator",
	CreatorId:   "creator_id",
	Updator:     "updator",
	UpdatorId:   "updator_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewOrganizationUnitsDao creates and returns a new DAO object for table data access.
func NewOrganizationUnitsDao() *OrganizationUnitsDao {
	return &OrganizationUnitsDao{
		group:   "default",
		table:   "organization_units",
		columns: organizationUnitsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrganizationUnitsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrganizationUnitsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrganizationUnitsDao) Columns() OrganizationUnitsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrganizationUnitsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrganizationUnitsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrganizationUnitsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
