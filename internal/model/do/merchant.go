// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
    "github.com/gogf/gf/v2/frame/g"
)

// Merchant is the golang structure of table merchant for DAO operations like Where/Data.
type Merchant struct {
    g.Meta     `orm:"table:merchant, do:true"`
    MerchantId interface{} // 商户ID
    Name       interface{} // 商户名称
    Contact    interface{} // 联系方式
    Address    interface{} // 商户地址
}