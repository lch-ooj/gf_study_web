package dao

import (
    "gf-demo-user/internal/dao/internal"
)

type internalCustomerDao = *internal.CustomerDao

type customerDao struct {
    internalCustomerDao
}

var (
    Customer = customerDao{
        internal.NewCustomerDao(),
    }
)