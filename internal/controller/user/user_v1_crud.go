package user

import (
    "context"

    v1 "gf-demo-user/api/user/v1"
    "gf-demo-user/internal/dao"
    "gf-demo-user/internal/model/entity"
)

// CustomerList 获取外卖用户列表
func (c *ControllerV1) CustomerList(ctx context.Context, req *v1.CustomerListReq) (res *v1.CustomerListRes, err error) {
    var customers []entity.Customer
    err = dao.Customer.Ctx(ctx).Scan(&customers)
    if err != nil {
        return nil, err
    }
    return &v1.CustomerListRes{List: customers}, nil
}

// CustomerCreate 创建外卖用户
func (c *ControllerV1) CustomerCreate(ctx context.Context, req *v1.CustomerCreateReq) (res *v1.CustomerCreateRes, err error) {
    result, err := dao.Customer.Ctx(ctx).Data(map[string]interface{}{
        "name":     req.Name,
        "password": req.Password,
        "phone":    req.Phone,
        "address":  req.Address,
    }).Insert()
    if err != nil {
        return nil, err
    }
    id, _ := result.LastInsertId()
    return &v1.CustomerCreateRes{Uid: id}, nil
}

// CustomerUpdate 更新外卖用户
func (c *ControllerV1) CustomerUpdate(ctx context.Context, req *v1.CustomerUpdateReq) (res *v1.CustomerUpdateRes, err error) {
    data := make(map[string]interface{})
    if req.Name != "" {
        data["name"] = req.Name
    }
    if req.Phone != "" {
        data["phone"] = req.Phone
    }
    if req.Address != "" {
        data["address"] = req.Address
    }

    _, err = dao.Customer.Ctx(ctx).Where("uid", req.Uid).Data(data).Update()
    if err != nil {
        return nil, err
    }
    return &v1.CustomerUpdateRes{Success: true}, nil
}

// CustomerDelete 删除外卖用户
func (c *ControllerV1) CustomerDelete(ctx context.Context, req *v1.CustomerDeleteReq) (res *v1.CustomerDeleteRes, err error) {
    _, err = dao.Customer.Ctx(ctx).Where("uid", req.Uid).Delete()
    if err != nil {
        return nil, err
    }
    return &v1.CustomerDeleteRes{Success: true}, nil
}

// CustomerGet 获取外卖用户详情
func (c *ControllerV1) CustomerGet(ctx context.Context, req *v1.CustomerGetReq) (res *v1.CustomerGetRes, err error) {
    var customer entity.Customer
    err = dao.Customer.Ctx(ctx).Where("uid", req.Uid).Scan(&customer)
    if err != nil {
        return nil, err
    }
    return &v1.CustomerGetRes{
        Uid:     customer.Uid,
        Name:    customer.Name,
        Phone:   customer.Phone,
        Address: customer.Address,
    }, nil
}