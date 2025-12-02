package merchant

import (
    "context"

    v1 "gf-demo-user/api/merchant/v1"
    "gf-demo-user/internal/dao"
    "gf-demo-user/internal/model/entity"
)

// ControllerV1 商户控制器
type ControllerV1 struct{}

// NewV1 创建控制器实例
func NewV1() *ControllerV1 {
    return &ControllerV1{}
}

// MerchantList 获取商户列表
func (c *ControllerV1) MerchantList(ctx context.Context, req *v1.MerchantListReq) (res *v1.MerchantListRes, err error) {
    var merchants []entity.Merchant
    err = dao.Merchant.Ctx(ctx).Scan(&merchants)
    if err != nil {
        return nil, err
    }
    return &v1.MerchantListRes{List: merchants}, nil
}

// MerchantCreate 创建商户
func (c *ControllerV1) MerchantCreate(ctx context.Context, req *v1.MerchantCreateReq) (res *v1.MerchantCreateRes, err error) {
    result, err := dao.Merchant.Ctx(ctx).Data(map[string]interface{}{
        "name":    req.Name,
        "contact": req.Contact,
        "address": req.Address,
    }).Insert()
    if err != nil {
        return nil, err
    }
    id, _ := result.LastInsertId()
    return &v1.MerchantCreateRes{MerchantId: id}, nil
}

// MerchantUpdate 更新商户
func (c *ControllerV1) MerchantUpdate(ctx context.Context, req *v1.MerchantUpdateReq) (res *v1.MerchantUpdateRes, err error) {
    data := make(map[string]interface{})
    if req.Name != "" {
        data["name"] = req.Name
    }
    if req.Contact != "" {
        data["contact"] = req.Contact
    }
    if req.Address != "" {
        data["address"] = req.Address
    }

    _, err = dao.Merchant.Ctx(ctx).Where("merchant_id", req.MerchantId).Data(data).Update()
    if err != nil {
        return nil, err
    }
    return &v1.MerchantUpdateRes{Success: true}, nil
}

// MerchantDelete 删除商户
func (c *ControllerV1) MerchantDelete(ctx context.Context, req *v1.MerchantDeleteReq) (res *v1.MerchantDeleteRes, err error) {
    _, err = dao.Merchant.Ctx(ctx).Where("merchant_id", req.MerchantId).Delete()
    if err != nil {
        return nil, err
    }
    return &v1.MerchantDeleteRes{Success: true}, nil
}

// MerchantGet 获取商户详情
func (c *ControllerV1) MerchantGet(ctx context.Context, req *v1.MerchantGetReq) (res *v1.MerchantGetRes, err error) {
    var merchant entity.Merchant
    err = dao.Merchant.Ctx(ctx).Where("merchant_id", req.MerchantId).Scan(&merchant)
    if err != nil {
        return nil, err
    }
    return &v1.MerchantGetRes{
        MerchantId: merchant.MerchantId,
        Name:       merchant.Name,
        Contact:    merchant.Contact,
        Address:    merchant.Address,
    }, nil
}