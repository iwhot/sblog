package dao

import (
	"github.com/gofiber/fiber/v2"
	"math"
	"sblog/app/gorm/model"
	"sblog/app/logic"
	"sblog/system/common"
	"time"
)

type tags struct {
}

var DefaultTagsDao = tags{}

// Save 保存
func (d tags) Save(ctx *fiber.Ctx) error {
	type param struct {
		ID      uint32 `json:"id"`
		TagName string `json:"tag_name"`
	}

	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var tg = model.Tags{
		TagName:    p.TagName,
		Updatetime: logic.LocalTime{Time: time.Now()},
	}

	if p.ID > 0 {
		//编辑
		tg.ID = p.ID
		return tg.Update(masterDB)
	}

	tg.Createtime = logic.LocalTime{Time: time.Now()}
	return tg.Create(masterDB)
}

// GetList 获取列表
func (d tags) GetList(ctx *fiber.Ctx) ([]*model.Tags, map[string]interface{}, error) {
	var scs []*model.Tags
	var err error
	var rp = make(map[string]interface{})

	p := new(listParam)
	if err = common.ParseParam(ctx, p); err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return scs, rp, err
	}

	var sc = model.Tags{}
	scs, err = sc.GetList(masterDB, p.Page, p.PageSize)
	if err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return scs, rp, err
	}

	count := sc.Count(masterDB)

	allPage := int(math.Ceil(float64(count) / float64(p.PageSize)))
	rp["count"] = count
	rp["all"] = allPage
	rp["page"] = p.Page
	rp["page_size"] = p.PageSize

	return scs, rp, nil
}

// Delete 删除
func (d tags) Delete(ctx *fiber.Ctx) error{
	type param struct {
		ID uint32 `json:"id" xml:"id" form:"id"`
	}
	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}
	var sc = model.Tags{
		ID: p.ID,
	}
	return sc.Delete(masterDB)
}
