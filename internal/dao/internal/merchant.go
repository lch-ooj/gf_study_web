package internal

import (
    "context"

    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
)

// MerchantDao is the data access object for table merchant.
type MerchantDao struct {
    table   string
    group   string
    columns MerchantColumns
}

// MerchantColumns defines and stores column names for table merchant.
type MerchantColumns struct {
    MerchantId string
    Name       string
    Contact    string
    Address    string
}

// merchantColumns holds the columns for table merchant.
var merchantColumns = MerchantColumns{
    MerchantId: "merchant_id",
    Name:       "name",
    Contact:    "contact",
    Address:    "address",
}

// NewMerchantDao creates and returns a new DAO object for table data access.
func NewMerchantDao() *MerchantDao {
    return &MerchantDao{
        group:   "default",
        table:   "merchant",
        columns: merchantColumns,
    }
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MerchantDao) DB() gdb.DB {
    return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MerchantDao) Table() string {
    return dao.table
}

// Columns returns all column names of current dao.
func (dao *MerchantDao) Columns() MerchantColumns {
    return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MerchantDao) Group() string {
    return dao.group
}

// Ctx creates and returns the Model for current DAO.
func (dao *MerchantDao) Ctx(ctx context.Context) *gdb.Model {
    return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *MerchantDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
    return dao.Ctx(ctx).Transaction(ctx, f)
}