package v1

import "github.com/gogf/gf/v2/frame/g"

// ========== 外卖用户列表 ==========
type CustomerListReq struct {
    g.Meta `path:"/customer/list" method:"get" tags:"外卖用户管理" summary:"获取外卖用户列表"`
}
type CustomerListRes struct {
    List interface{} `json:"list" dc:"用户列表"`
}

// ========== 创建外卖用户 ==========
type CustomerCreateReq struct {
    g.Meta   `path:"/customer/create" method:"post" tags:"外卖用户管理" summary:"创建外卖用户"`
    Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
    Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
    Phone    string `json:"phone" dc:"手机号"`
    Address  string `json:"address" dc:"地址"`
}
type CustomerCreateRes struct {
    Uid int64 `json:"uid" dc:"用户ID"`
}

// ========== 更新外卖用户 ==========
type CustomerUpdateReq struct {
    g.Meta  `path:"/customer/update" method:"put" tags:"外卖用户管理" summary:"更新外卖用户"`
    Uid     int    `json:"uid" v:"required#用户ID不能为空" dc:"用户ID"`
    Name    string `json:"name" dc:"用户名"`
    Phone   string `json:"phone" dc:"手机号"`
    Address string `json:"address" dc:"地址"`
}
type CustomerUpdateRes struct {
    Success bool `json:"success" dc:"是否成功"`
}

// ========== 删除外卖用户 ==========
type CustomerDeleteReq struct {
    g.Meta `path:"/customer/delete" method:"delete" tags:"外卖用户管理" summary:"删除外卖用户"`
    Uid    int `json:"uid" v:"required#用户ID不能为空" dc:"用户ID"`
}
type CustomerDeleteRes struct {
    Success bool `json:"success" dc:"是否成功"`
}

// ========== 获取单个外卖用户 ==========
type CustomerGetReq struct {
    g.Meta `path:"/customer/get" method:"get" tags:"外卖用户管理" summary:"获取外卖用户详情"`
    Uid    int `json:"uid" in:"query" v:"required#用户ID不能为空" dc:"用户ID"`
}
type CustomerGetRes struct {
    Uid     int    `json:"uid" dc:"用户ID"`
    Name    string `json:"name" dc:"用户名"`
    Phone   string `json:"phone" dc:"手机号"`
    Address string `json:"address" dc:"地址"`
}