// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClientCorsOriginsDao is the data access object for table client_cors_origins.
type ClientCorsOriginsDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns ClientCorsOriginsColumns // columns contains all the column names of Table for convenient usage.
}

// ClientCorsOriginsColumns defines and stores column names for table client_cors_origins.
type ClientCorsOriginsColumns struct {
	Origin    string //
	ClientId  string //
	Id        string //
	Creator   string //
	CreatorId string //
	Updator   string //
	UpdatorId string //
	CreatedAt string //
	UpdatedAt string //
}

// clientCorsOriginsColumns holds the columns for table client_cors_origins.
var clientCorsOriginsColumns = ClientCorsOriginsColumns{
	Origin:    "origin",
	ClientId:  "client_id",
	Id:        "id",
	Creator:   "creator",
	CreatorId: "creator_id",
	Updator:   "updator",
	UpdatorId: "updator_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewClientCorsOriginsDao creates and returns a new DAO object for table data access.
func NewClientCorsOriginsDao() *ClientCorsOriginsDao {
	return &ClientCorsOriginsDao{
		group:   "default",
		table:   "client_cors_origins",
		columns: clientCorsOriginsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ClientCorsOriginsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ClientCorsOriginsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ClientCorsOriginsDao) Columns() ClientCorsOriginsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ClientCorsOriginsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ClientCorsOriginsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ClientCorsOriginsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
