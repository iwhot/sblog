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

type menu struct {
}

var DefaultMenuDao = menu{}

// Save 保存
func (d menu) Save(ctx *fiber.Ctx) error {
	type param struct {
		ID     uint32 `json:"id"`
		PID    uint32 `json:"pid"`
		Title  string `json:"title"`
		Url    string `json:"url"`
		Target string `json:"target"`
		Status uint8  `json:"status"`
		Sort   uint32 `json:"sort"`
		Group string `json:"group"`
	}

	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var lk = model.Menu{
		PID:        p.PID,
		Title:      p.Title,
		Url:        p.Url,
		Target:     p.Target,
		Status:     p.Status,
		Sort:       p.Sort,
		Updatetime: logic.LocalTime{Time: time.Now()},
		Group: p.Group,
	}

	if p.ID > 0 {
		//编辑
		lk.ID = p.ID
		return lk.Updates(masterDB)
	}

	lk.Createtime = logic.LocalTime{Time: time.Now()}
	return lk.Create(masterDB)
}

// Delete 删除
func (d menu) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	if id <= 0 {
		return errors.New("删除失败")
	}
	var lk = model.Menu{
		ID: uint32(id),
	}
	return lk.Delete(masterDB)
}

// GetList 列表
func (d menu) GetList(ctx *fiber.Ctx) ([]*model.Menu, error) {
	type param struct {
		ID uint32 `json:"id" xml:"id" form:"id"`
	}
	var p = new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return nil, err
	}

	var ct = model.Menu{
		PID: p.ID,
	}
	return ct.GetList(masterDB)
}

// GetParentMenu 获取上级菜单,只允许一级下拉
func (d menu) GetParentMenu(ctx *fiber.Ctx) ([]*model.Menu, error) {
	var ct = model.Menu{}
	return ct.GetParentList(masterDB)
}

// CascadeMenu 获取级联菜单
func (d menu) CascadeMenu(ctx *fiber.Ctx) ([]model.Cascade, error) {
	var ar = model.Menu{}
	return ar.CascadeCategory(masterDB)
}

// GetAllMenuId 获取所有上级id
func (d menu) GetAllMenuId(ctx *fiber.Ctx) []uint32 {
	var ret []uint32
	var id, _ = strconv.Atoi(ctx.Params("id"))
	if id <= 0 {
		return ret
	}

	var cat = model.Menu{
		ID: uint32(id),
	}

	return cat.GetAllMenuId(masterDB,ret)
}

// UpdateInfo 更改信息
func (d menu) UpdateInfo(ctx *fiber.Ctx) error {
	var info = make(map[string]interface{})
	if err := common.ParseParam(ctx, &info); err != nil {
		return err
	}

	var ct = model.Menu{}
	return ct.UpdateInfo(masterDB, info)
}