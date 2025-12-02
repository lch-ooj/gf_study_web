package v1

import "github.com/gogf/gf/v2/frame/g"

// ========== 商户列表 ==========
type MerchantListReq struct {
    g.Meta `path:"/merchant/list" method:"get" tags:"商户管理" summary:"获取商户列表"`
}
type MerchantListRes struct {
    List interface{} `json:"list" dc:"商户列表"`
}

// ========== 创建商户 ==========
type MerchantCreateReq struct {
    g.Meta  `path:"/merchant/create" method:"post" tags:"商户管理" summary:"创建商户"`
    Name    string `json:"name" v:"required#商户名称不能为空" dc:"商户名称"`
    Contact string `json:"contact" dc:"联系方式"`
    Address string `json:"address" dc:"商户地址"`
}
type MerchantCreateRes struct {
    MerchantId int64 `json:"merchantId" dc:"商户ID"`
}

// ========== 更新商户 ==========
type MerchantUpdateReq struct {
    g.Meta     `path:"/merchant/update" method:"put" tags:"商户管理" summary:"更新商户"`
    MerchantId int    `json:"merchantId" v:"required#商户ID不能为空" dc:"商户ID"`
    Name       string `json:"name" dc:"商户名称"`
    Contact    string `json:"contact" dc:"联系方式"`
    Address    string `json:"address" dc:"商户地址"`
}
type MerchantUpdateRes struct {
    Success bool `json:"success" dc:"是否成功"`
}

// ========== 删除商户 ==========
type MerchantDeleteReq struct {
    g.Meta     `path:"/merchant/delete" method:"delete" tags:"商户管理" summary:"删除商户"`
    MerchantId int `json:"merchantId" v:"required#商户ID不能为空" dc:"商户ID"`
}
type MerchantDeleteRes struct {
    Success bool `json:"success" dc:"是否成功"`
}

// ========== 获取单个商户 ==========
type MerchantGetReq struct {
    g.Meta     `path:"/merchant/get" method:"get" tags:"商户管理" summary:"获取商户详情"`
    MerchantId int `json:"merchantId" in:"query" v:"required#商户ID不能为空" dc:"商户ID"`
}
type MerchantGetRes struct {
    MerchantId int    `json:"merchantId" dc:"商户ID"`
    Name       string `json:"name" dc:"商户名称"`
    Contact    string `json:"contact" dc:"联系方式"`
    Address    string `json:"address" dc:"商户地址"`
}