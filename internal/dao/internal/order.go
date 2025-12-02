package internal

import (
    "context"

    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
)

// OrderDao is the data access object for table order.
type OrderDao struct {
    table   string
    group   string
    columns OrderColumns
}

// OrderColumns defines and stores column names for table order.
type OrderColumns struct {
    OrderId         string
    Uid             string
    MerchantId      string
    Status          string
    OrderTime       string
    DeliveryAddress string
}

// orderColumns holds the columns for table order.
var orderColumns = OrderColumns{
    OrderId:         "order_id",
    Uid:             "uid",
    MerchantId:      "merchant_id",
    Status:          "status",
    OrderTime:       "order_time",
    DeliveryAddress: "delivery_address",
}

// NewOrderDao creates and returns a new DAO object for table data access.
func NewOrderDao() *OrderDao {
    return &OrderDao{
        group:   "default",
        table:   "`order`",
        columns: orderColumns,
    }
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderDao) DB() gdb.DB {
    return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderDao) Table() string {
    return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderDao) Columns() OrderColumns {
    return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderDao) Group() string {
    return dao.group
}

// Ctx creates and returns the Model for current DAO.
func (dao *OrderDao) Ctx(ctx context.Context) *gdb.Model {
    return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *OrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
    return dao.Ctx(ctx).Transaction(ctx, f)
}