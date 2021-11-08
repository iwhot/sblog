package dao

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"sblog/app/gorm/model"
	"sblog/app/logic"
	"sblog/system/common"
	"strconv"
	"time"
)

type category struct {
}

var DefaultCategoryDao = category{}

// Save 保存
func (d category) Save(ctx *fiber.Ctx) error {
	type param struct {
		ID     uint32 `json:"id" xml:"id" form:"id"`
		PID    uint32 `json:"pid" xml:"pid" form:"pid"`
		Name   string `json:"name" xml:"name" form:"name"`
		Key    string `json:"key" xml:"key" form:"key"`
		Status uint8  `json:"status" xml:"status" form:"status"`
		Sort   uint32 `json:"sort" xml:"sort" form:"sort"`
		Desc   string `json:"desc" xml:"desc"`
		Thumb  string `xml:"thumb" xml:"thumb"`
	}
	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}
	var ct = model.Category{
		PID:        p.PID,
		Name:       p.Name,
		Key:        p.Key,
		Status:     p.Status,
		Sort:       p.Sort,
		Updatetime: logic.LocalTime{Time: time.Now()},
		Desc:       p.Desc,
		Thumb:      p.Thumb,
	}
	if p.ID > 0 {
		ct.ID = p.ID

		if p.ID == p.PID {
			ct.PID = 0
		}

		return ct.Update(masterDB)
	} else {
		ct.Createtime = logic.LocalTime{Time: time.Now()}
		return ct.Create(masterDB)
	}
}

// GetList 列表
func (d category) GetList(ctx *fiber.Ctx) ([]*model.Category, error) {
	type param struct {
		ID uint32 `json:"id" xml:"id" form:"id"`
	}
	var p = new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return nil, err
	}

	var ct = model.Category{
		PID: p.ID,
	}
	return ct.GetList(masterDB)
}

// Delete 删除
func (d category) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	if id <= 0 {
		return errors.New("删除失败")
	}
	var ct = model.Category{ID: uint32(id)}
	return ct.Delete(masterDB)
}

// GetParentCategory 获取上级菜单,只允许一级下拉
func (d category) GetParentCategory(ctx *fiber.Ctx) ([]*model.Category, error) {
	var ct = model.Category{}
	return ct.GetParentList(masterDB)
}

// UpdateInfo 更改信息
func (d category) UpdateInfo(ctx *fiber.Ctx) error {
	var info = make(map[string]interface{})
	if err := common.ParseParam(ctx, &info); err != nil {
		return err
	}

	var ct = model.Category{}
	return ct.UpdateInfo(masterDB, info)
}

// CascadeCategory 获取级联菜单
func (d category) CascadeCategory(ctx *fiber.Ctx) ([]model.Cascade, error) {
	var ar = model.Category{}
	return ar.CascadeCategory(masterDB)
}

// GetAllCategoryId 获取所有上级id
func (d category) GetAllCategoryId(ctx *fiber.Ctx) []uint32 {
	var ret []uint32
	var id, _ = strconv.Atoi(ctx.Params("id"))
	if id <= 0 {
		return ret
	}

	var cat = model.Category{
		ID: uint32(id),
	}

	ret = append(ret,uint32(id))
	return cat.GetAllCategoryId(masterDB, ret)
}
