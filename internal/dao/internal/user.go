package internal

import (
    "context"

    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table user.
type UserDao struct {
    table   string
    group   string
    columns UserColumns
}

// UserColumns defines and stores column names for table user.
type UserColumns struct {
    Id        string
    Passport  string
    Password  string
    Nickname  string
    CreatedAt string
    UpdatedAt string
}

// userColumns holds the columns for table user.
var userColumns = UserColumns{
    Id:        "id",
    Passport:  "passport",
    Password:  "password",
    Nickname:  "nickname",
    CreatedAt: "created_at",
    UpdatedAt: "updated_at",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
    return &UserDao{
        group:   "default",
        table:   "user",
        columns: userColumns,
    }
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
    return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
    return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
    return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
    return dao.group
}

// Ctx creates and returns the Model for current DAO.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
    return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
    return dao.Ctx(ctx).Transaction(ctx, f)
}