package item

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "gf-demo-user/api/item/v1"
	"gf-demo-user/internal/dao"
	"gf-demo-user/internal/model/do"
	"gf-demo-user/internal/model/entity"
)

// List 获取物品列表
func (c *ControllerV1) ItemList(ctx context.Context, req *v1.ItemListReq) (res *v1.ItemListRes, err error) {
	var items []entity.Item
	if err = dao.Item.Ctx(ctx).Scan(&items); err != nil {
		return nil, err
	}
	res = &v1.ItemListRes{List: items}
	return
}

// ItemCreate 创建物品
func (c *ControllerV1) ItemCreate(ctx context.Context, req *v1.ItemCreateReq) (res *v1.ItemCreateRes, err error) {
	data := do.Item{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}
	result, err := dao.Item.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	res = &v1.ItemCreateRes{Item: entity.Item{
		Id:          uint(id),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}}
	return
}

// ItemGet 获取物品详情
func (c *ControllerV1) ItemGet(ctx context.Context, req *v1.ItemGetReq) (res *v1.ItemGetRes, err error) {
	var item entity.Item
	if err = dao.Item.Ctx(ctx).Where(dao.Item.Columns().Id, req.Id).Scan(&item); err != nil {
		return nil, err
	}
	res = &v1.ItemGetRes{Item: item}
	return
}

// ItemUpdate 更新物品
func (c *ControllerV1) ItemUpdate(ctx context.Context, req *v1.ItemUpdateReq) (res *v1.ItemUpdateRes, err error) {
	data := g.Map{}
	if req.Name != "" {
		data[dao.Item.Columns().Name] = req.Name
	}
	if req.Description != "" {
		data[dao.Item.Columns().Description] = req.Description
	}
	if req.Price != 0 {
		data[dao.Item.Columns().Price] = req.Price
	}
	if req.Stock != 0 {
		data[dao.Item.Columns().Stock] = req.Stock
	}
	if len(data) == 0 {
		return nil, gerror.New("没有需要更新的字段")
	}
	if _, err = dao.Item.Ctx(ctx).Where(dao.Item.Columns().Id, req.Id).Data(data).Update(); err != nil {
		return nil, err
	}
	return c.ItemGet(ctx, &v1.ItemGetReq{Id: req.Id})
}

// ItemDelete 删除物品
func (c *ControllerV1) ItemDelete(ctx context.Context, req *v1.ItemDeleteReq) (res *v1.ItemDeleteRes, err error) {
	if _, err = dao.Item.Ctx(ctx).Where(dao.Item.Columns().Id, req.Id).Delete(); err != nil {
		return nil, err
	}
	res = &v1.ItemDeleteRes{Success: true}
	return
}
