package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ItemDao is the data access object for table item.
type ItemDao struct {
	table   string
	group   string
	columns ItemColumns
}

// ItemColumns defines and stores column names for table item.
type ItemColumns struct {
	Id          string
	Name        string
	Description string
	Price       string
	Stock       string
	CreatedAt   string
	UpdatedAt   string
}

// itemColumns holds the columns for table item.
var itemColumns = ItemColumns{
	Id:          "id",
	Name:        "name",
	Description: "description",
	Price:       "price",
	Stock:       "stock",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewItemDao creates and returns a new DAO object for table data access.
func NewItemDao() *ItemDao {
	return &ItemDao{
		group:   "default",
		table:   "item",
		columns: itemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ItemDao) Columns() ItemColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO.
func (dao *ItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *ItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
