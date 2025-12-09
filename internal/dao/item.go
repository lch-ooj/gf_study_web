package dao

import "gf-demo-user/internal/dao/internal"

type internalItemDao = *internal.ItemDao

type itemDao struct {
	internalItemDao
}

var Item = itemDao{internal.NewItemDao()}
