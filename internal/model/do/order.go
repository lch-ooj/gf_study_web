// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
    "github.com/gogf/gf/v2/frame/g"
)

// Order is the golang structure of table order for DAO operations like Where/Data.
type Order struct {
    g.Meta          `orm:"table:order, do:true"`
    OrderId         interface{} // 订单ID
    Uid             interface{} // 用户ID
    MerchantId      interface{} // 商户ID
    Status          interface{} // 订单状态
    OrderTime       interface{} // 下单时间
    DeliveryAddress interface{} // 配送地址
}