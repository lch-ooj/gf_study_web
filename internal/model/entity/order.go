// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import "github.com/gogf/gf/v2/os/gtime"

// Order is the golang structure for table order.
type Order struct {
    OrderId         int         `json:"orderId"         orm:"order_id"         description:"订单ID"`
    Uid             int         `json:"uid"             orm:"uid"              description:"用户ID"`
    MerchantId      int         `json:"merchantId"      orm:"merchant_id"      description:"商户ID"`
    Status          string      `json:"status"          orm:"status"           description:"订单状态"`
    OrderTime       *gtime.Time `json:"orderTime"       orm:"order_time"       description:"下单时间"`
    DeliveryAddress string      `json:"deliveryAddress" orm:"delivery_address" description:"配送地址"`
}