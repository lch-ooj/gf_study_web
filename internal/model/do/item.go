package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Item is the golang structure of table item for DAO operations like Where/Data.
type Item struct {
	g.Meta      `orm:"table:item, do:true"`
	Id          interface{}
	Name        interface{}
	Description interface{}
	Price       interface{}
	Stock       interface{}
	CreatedAt   *gtime.Time
	UpdatedAt   *gtime.Time
}
