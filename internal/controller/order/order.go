// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package order

import (
    "context"

    orderv1 "gf-demo-user/api/order/v1"
)

type IOrderV1 interface {
    OrderList(ctx context.Context, req *orderv1.OrderListReq) (res *orderv1.OrderListRes, err error)
    OrderCreate(ctx context.Context, req *orderv1.OrderCreateReq) (res *orderv1.OrderCreateRes, err error)
    OrderUpdate(ctx context.Context, req *orderv1.OrderUpdateReq) (res *orderv1.OrderUpdateRes, err error)
    OrderDelete(ctx context.Context, req *orderv1.OrderDeleteReq) (res *orderv1.OrderDeleteRes, err error)
    OrderGet(ctx context.Context, req *orderv1.OrderGetReq) (res *orderv1.OrderGetRes, err error)
}