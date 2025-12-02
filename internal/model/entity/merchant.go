// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Merchant is the golang structure for table merchant.
type Merchant struct {
    MerchantId int    `json:"merchantId" orm:"merchant_id" description:"商户ID"`
    Name       string `json:"name"       orm:"name"        description:"商户名称"`
    Contact    string `json:"contact"    orm:"contact"     description:"联系方式"`
    Address    string `json:"address"    orm:"address"     description:"商户地址"`
}