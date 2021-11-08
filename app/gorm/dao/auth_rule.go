package dao

import (
	"github.com/gofiber/fiber/v2"
	"sblog/app/gorm/model"
	"sblog/app/logic"
	"sblog/system/common"
	"strconv"
	"strings"
	"time"
)

type authRule struct {
}

// saveAuthRuleParam 保存菜单参数
type saveAuthRuleParam struct {
	ID        uint32 `json:"id" xml:"id" form:"id"`
	PID       uint32 `json:"pid" xml:"pid" form:"pid"`
	Url       string `json:"url" xml:"url" form:"url"`
	Title     string `json:"title" xml:"title" form:"title"`
	Icon      string `json:"icon" xml:"icon" form:"icon"`
	Component string `json:"component" xml:"component" form:"component"`
	Status    uint8  `json:"status" xml:"status" form:"status"`
	Remark    string `json:"remark" xml:"remark" form:"remark"`
	Redirect  string `json:"redirect" xml:"redirect" form:"redirect"`
	Sort      uint32 `json:"sort" xml:"sort" form:"sort"`
	IsMenu    uint8  `json:"is_menu" xml:"is_menu" form:"is_menu"`
	IsNotDel  uint8  `json:"is_not_del" xml:"is_not_del" form:"is_not_del"`
}

//删除参数
type saveAuthDeleteParam struct {
	ID uint32 `json:"id" xml:"id" form:"id"`
}

//删除参数
type saveAuthDeleteAllParam struct {
	ID string `json:"id" xml:"id" form:"id"`
}

var DefaultAuthRuleDao = authRule{}

// GetMyMenu 获取我的菜单
func (d authRule) GetMyMenu(uid uint32) ([]model.LeftMenu, error) {
	var ar = model.AuthRule{}
	m1, err := ar.GetMyMenu(masterDB, uid)
	if err != nil {
		return nil, err
	}

	//获取无限极
	m2 := model.GetLeftMenuTree(m1, 0)
	return m2, nil
}

// GetList 获取菜单列表
func (d authRule) GetList(ctx *fiber.Ctx) ([]*model.AuthRule, error) {
	var p = new(saveAuthDeleteParam)
	if err := common.ParseParam(ctx, p); err != nil {
		return nil, err
	}
	var ar = model.AuthRule{
		PID: p.ID,
	}

	return ar.GetList(masterDB)
}

// Save 保存菜单
func (d authRule) Save(ctx *fiber.Ctx) error {
	p := new(saveAuthRuleParam)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}
	var ar = model.AuthRule{
		PID:       p.PID,
		Url:       p.Url,
		Title:     p.Title,
		Icon:      p.Icon,
		Component: p.Component,
		Status:    p.Status,
		Remark:    p.Remark,
		Redirect:  p.Redirect,
		Sort:      p.Sort,
		IsMenu:    p.IsMenu,
		IsNotDel:  p.IsNotDel,
	}

	ar.Updatetime = logic.LocalTime{Time: time.Now()}
	if p.ID > 0 {
		ar.ID = p.ID
		return ar.Update(masterDB)
	} else {
		ar.Createtime = logic.LocalTime{Time: time.Now()}
		return ar.Create(masterDB)
	}
}

// Delete 删除
func (d authRule) Delete(ctx *fiber.Ctx) error {
	p := new(saveAuthDeleteParam)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var ar = model.AuthRule{
		ID: p.ID,
	}
	return ar.Delete(masterDB)
}

// DeleteAll 批量删除
func (d authRule) DeleteAll(ctx *fiber.Ctx) error {
	p := new(saveAuthDeleteAllParam)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var ar = model.AuthRule{}
	ids := strings.Split(p.ID, ",")
	return ar.DeleteAll(masterDB, ids)
}

// GetPid 获取菜单所有的pid
func (d authRule) GetPid(ctx *fiber.Ctx) []uint32 {
	var ret []uint32
	p := new(saveAuthDeleteParam)
	if err := common.ParseParam(ctx, p); err != nil {
		return ret
	}

	var ar = model.AuthRule{
		ID: p.ID,
	}

	ars, err := ar.GetAllParentId(masterDB, ret)
	if err != nil {
		return ret
	}

	return ars
}

// UpdateStatus 更新菜单状态
func (d authRule) UpdateStatus(ctx *fiber.Ctx) error {
	type updateStatus struct {
		ID     uint32 `json:"id" xml:"id" form:"id"`
		Status uint8  `json:"status" xml:"status" form:"status"`
	}
	p := new(updateStatus)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var ar = model.AuthRule{
		ID:     p.ID,
		Status: p.Status,
	}

	var dt = make(map[string]interface{})
	dt["status"] = p.Status
	return ar.UpdatePartial(masterDB, dt)
}

// UpdateSort 更新排序
func (d authRule) UpdateSort(ctx *fiber.Ctx) error {
	type updateSort struct {
		ID   uint32 `json:"id" xml:"id" form:"id"`
		Sort string `json:"sort" xml:"sort" form:"sort"`
	}

	p := new(updateSort)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var s, _ = strconv.Atoi(p.Sort)

	var ar = model.AuthRule{
		ID:   p.ID,
		Sort: uint32(s),
	}

	var dt = make(map[string]interface{})
	dt["sort"] = p.Sort
	return ar.UpdatePartial(masterDB, dt)
}

// CascadeMenu 获取级联菜单
func (d authRule) CascadeMenu(ctx *fiber.Ctx) ([]model.Cascade, error) {
	var ar = model.AuthRule{}
	return ar.CascadeMenu(masterDB)
}
