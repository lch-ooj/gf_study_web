package order

import (
    "context"
    "time"

    v1 "gf-demo-user/api/order/v1"
    "gf-demo-user/internal/dao"
    "gf-demo-user/internal/model/entity"
)

// ControllerV1 订单控制器
type ControllerV1 struct{}

// NewV1 创建控制器实例
func NewV1() *ControllerV1 {
    return &ControllerV1{}
}

// OrderList 获取订单列表
func (c *ControllerV1) OrderList(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error) {
    var orders []entity.Order
    err = dao.Order.Ctx(ctx).Scan(&orders)
    if err != nil {
        return nil, err
    }
    return &v1.OrderListRes{List: orders}, nil
}

// OrderCreate 创建订单
func (c *ControllerV1) OrderCreate(ctx context.Context, req *v1.OrderCreateReq) (res *v1.OrderCreateRes, err error) {
    result, err := dao.Order.Ctx(ctx).Data(map[string]interface{}{
        "uid":              req.Uid,
        "merchant_id":      req.MerchantId,
        "status":           "pending",
        "order_time":       time.Now(),
        "delivery_address": req.DeliveryAddress,
    }).Insert()
    if err != nil {
        return nil, err
    }
    id, _ := result.LastInsertId()
    return &v1.OrderCreateRes{OrderId: id}, nil
}

// OrderUpdate 更新订单状态
func (c *ControllerV1) OrderUpdate(ctx context.Context, req *v1.OrderUpdateReq) (res *v1.OrderUpdateRes, err error) {
    _, err = dao.Order.Ctx(ctx).Where("order_id", req.OrderId).Data(map[string]interface{}{
        "status": req.Status,
    }).Update()
    if err != nil {
        return nil, err
    }
    return &v1.OrderUpdateRes{Success: true}, nil
}

// OrderDelete 删除订单
func (c *ControllerV1) OrderDelete(ctx context.Context, req *v1.OrderDeleteReq) (res *v1.OrderDeleteRes, err error) {
    _, err = dao.Order.Ctx(ctx).Where("order_id", req.OrderId).Delete()
    if err != nil {
        return nil, err
    }
    return &v1.OrderDeleteRes{Success: true}, nil
}

// OrderGet 获取订单详情
func (c *ControllerV1) OrderGet(ctx context.Context, req *v1.OrderGetReq) (res *v1.OrderGetRes, err error) {
    var order entity.Order
    err = dao.Order.Ctx(ctx).Where("order_id", req.OrderId).Scan(&order)
    if err != nil {
        return nil, err
    }

    orderTime := ""
    if order.OrderTime != nil {
        orderTime = order.OrderTime.String()
    }

    return &v1.OrderGetRes{
        OrderId:         order.OrderId,
        Uid:             order.Uid,
        MerchantId:      order.MerchantId,
        Status:          order.Status,
        OrderTime:       orderTime,
        DeliveryAddress: order.DeliveryAddress,
    }, nil
}