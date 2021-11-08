package dao

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"math"
	"sblog/app/gorm/model"
	"sblog/app/logic"
	"sblog/system/common"
	"strings"
	"time"
)

type authGroup struct {
}

//保存参数
type saveParam struct {
	ID        uint32 `json:"id" xml:"id" form:"id"`
	PID       uint32 `json:"pid" xml:"pid" form:"pid"`
	Name      string `json:"name" xml:"name" form:"name"`
	Rules     string `json:"rules" xml:"rules" form:"rules"`
	Status    uint8  `json:"status" xml:"status" form:"status"`
	Sort      uint32 `json:"sort" xml:"sort" form:"sort"`
	Halfcheck string `json:"halfcheck" xml:"halfcheck" form:"halfcheck"`
}

var DefaultAuthGroupDao = authGroup{}

// Save 保存
func (d authGroup) Save(ctx *fiber.Ctx) error {
	var p = new(saveParam)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	if p.Name == "" {
		return errors.New("角色组名称不得为空")
	}

	var ag = model.AuthGroup{
		PID:       p.PID,
		Name:      p.Name,
		Rules:     p.Rules,
		Status:    p.Status,
		Sort:      p.Sort,
		Halfcheck: p.Halfcheck,
	}
	ag.Updatetime = logic.LocalTime{Time: time.Now()}
	if p.ID > 0 {
		ag.ID = p.ID
		return ag.Update(masterDB)
	} else {
		ag.Createtime = logic.LocalTime{Time: time.Now()}
		return ag.Create(masterDB)
	}
}

// GetList 获取列表
func (d authGroup) GetList(ctx *fiber.Ctx) ([]*model.AuthGroup, map[string]interface{}, error) {
	var ags []*model.AuthGroup
	var err error
	p := new(listParam)
	var rp = make(map[string]interface{})
	if err = common.ParseParam(ctx, p); err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return ags, rp, err
	}

	var ag = model.AuthGroup{}
	ags, err = ag.GetList(masterDB, p.Page, p.PageSize)
	if err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return ags, rp, err
	}

	count := ag.Count(masterDB)
	allPage := int(math.Ceil(float64(count) / float64(p.PageSize)))
	rp["count"] = count
	rp["all"] = allPage
	rp["page"] = p.Page
	rp["page_size"] = p.PageSize

	return ags, rp, nil
}

// Delete 删除
func (d authGroup) Delete(ctx *fiber.Ctx) error {
	type param = struct {
		ID uint32 `json:"id" xml:"id" form:"id"`
	}
	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var ag = model.AuthGroup{
		ID: p.ID,
	}

	return ag.Delete(masterDB)
}

// DeleteAll 批量删除
func (d authGroup) DeleteAll(ctx *fiber.Ctx) error {
	type param struct {
		ID string `json:"id" xml:"id" form:"id"`
	}
	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var ag = model.AuthGroup{}
	var ids = strings.Split(p.ID, ",")

	return ag.DeleteAll(masterDB, ids)
}

// GetAllGroup 获取所有组
func (d authGroup) GetAllGroup(ctx *fiber.Ctx) ([]map[string]interface{}, error) {
	var ag = model.AuthGroup{}
	return ag.GetAllList(masterDB)
}
