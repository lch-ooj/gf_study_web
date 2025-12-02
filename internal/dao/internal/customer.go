package internal

import (
    "context"

    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
)

// CustomerDao is the data access object for table user (外卖系统用户).
type CustomerDao struct {
    table   string
    group   string
    columns CustomerColumns
}

// CustomerColumns defines and stores column names for table user.
type CustomerColumns struct {
    Uid      string
    Name     string
    Password string
    Phone    string
    Address  string
}

// customerColumns holds the columns for table user.
var customerColumns = CustomerColumns{
    Uid:      "uid",
    Name:     "name",
    Password: "password",
    Phone:    "phone",
    Address:  "address",
}

// NewCustomerDao creates and returns a new DAO object for table data access.
func NewCustomerDao() *CustomerDao {
    return &CustomerDao{
        group:   "default",
        table:   "user",
        columns: customerColumns,
    }
}

// DB retrieves and returns the underlying raw database management object.
func (dao *CustomerDao) DB() gdb.DB {
    return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CustomerDao) Table() string {
    return dao.table
}

// Columns returns all column names of current dao.
func (dao *CustomerDao) Columns() CustomerColumns {
    return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CustomerDao) Group() string {
    return dao.group
}

// Ctx creates and returns the Model for current DAO.
func (dao *CustomerDao) Ctx(ctx context.Context) *gdb.Model {
    return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *CustomerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
    return dao.Ctx(ctx).Transaction(ctx, f)
}