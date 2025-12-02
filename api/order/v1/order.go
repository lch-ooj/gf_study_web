package v1

import "github.com/gogf/gf/v2/frame/g"

// ========== 订单列表 ==========
type OrderListReq struct {
    g.Meta `path:"/order/list" method:"get" tags:"订单管理" summary:"获取订单列表"`
}
type OrderListRes struct {
    List interface{} `json:"list" dc:"订单列表"`
}

// ========== 创建订单 ==========
type OrderCreateReq struct {
    g.Meta          `path:"/order/create" method:"post" tags:"订单管理" summary:"创建订单"`
    Uid             int    `json:"uid" v:"required#用户ID不能为空" dc:"用户ID"`
    MerchantId      int    `json:"merchantId" v:"required#商户ID不能为空" dc:"商户ID"`
    DeliveryAddress string `json:"deliveryAddress" v:"required#配送地址不能为空" dc:"配送地址"`
}
type OrderCreateRes struct {
    OrderId int64 `json:"orderId" dc:"订单ID"`
}

// ========== 更新订单状态 ==========
type OrderUpdateReq struct {
    g.Meta  `path:"/order/update" method:"put" tags:"订单管理" summary:"更新订单状态"`
    OrderId int    `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
    Status  string `json:"status" v:"required|in:pending,shipped,delivered,cancelled#状态不能为空|状态值必须是pending/shipped/delivered/cancelled" dc:"订单状态"`
}
type OrderUpdateRes struct {
    Success bool `json:"success" dc:"是否成功"`
}

// ========== 删除订单 ==========
type OrderDeleteReq struct {
    g.Meta  `path:"/order/delete" method:"delete" tags:"订单管理" summary:"删除订单"`
    OrderId int `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
}
type OrderDeleteRes struct {
    Success bool `json:"success" dc:"是否成功"`
}

// ========== 获取单个订单 ==========
type OrderGetReq struct {
    g.Meta  `path:"/order/get" method:"get" tags:"订单管理" summary:"获取订单详情"`
    OrderId int `json:"orderId" in:"query" v:"required#订单ID不能为空" dc:"订单ID"`
}
type OrderGetRes struct {
    OrderId         int    `json:"orderId" dc:"订单ID"`
    Uid             int    `json:"uid" dc:"用户ID"`
    MerchantId      int    `json:"merchantId" dc:"商户ID"`
    Status          string `json:"status" dc:"订单状态"`
    OrderTime       string `json:"orderTime" dc:"下单时间"`
    DeliveryAddress string `json:"deliveryAddress" dc:"配送地址"`
}