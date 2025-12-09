package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-demo-user/internal/model/entity"
)

// ItemListReq 获取物品列表
type ItemListReq struct {
	g.Meta `path:"/items" method:"get" tags:"Item" summary:"获取物品列表"`
}

type ItemListRes struct {
	List []entity.Item `json:"list"`
}

// ItemCreateReq 创建物品
type ItemCreateReq struct {
	g.Meta      `path:"/items" method:"post" tags:"Item" summary:"创建物品"`
	Name        string  `json:"name" v:"required#名称必填"`
	Description string  `json:"description"`
	Price       float64 `json:"price" v:"required#价格必填"`
	Stock       int     `json:"stock" v:"min:0#库存不能为负"`
}

type ItemCreateRes struct {
	Item entity.Item `json:"item"`
}

// ItemGetReq 获取单个物品
type ItemGetReq struct {
	g.Meta `path:"/items/{id}" method:"get" tags:"Item" summary:"获取物品详情"`
	Id     int `json:"id" in:"path" v:"required|min:1"`
}

type ItemGetRes struct {
	Item entity.Item `json:"item"`
}

// ItemUpdateReq 更新物品
type ItemUpdateReq struct {
	g.Meta      `path:"/items/{id}" method:"put" tags:"Item" summary:"更新物品"`
	Id          int     `json:"id" in:"path" v:"required|min:1"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock" v:"min:0"`
}

type ItemUpdateRes struct {
	Item entity.Item `json:"item"`
}

// ItemDeleteReq 删除物品
type ItemDeleteReq struct {
	g.Meta `path:"/items/{id}" method:"delete" tags:"Item" summary:"删除物品"`
	Id     int `json:"id" in:"path" v:"required|min:1"`
}

type ItemDeleteRes struct {
	Success bool `json:"success"`
}
