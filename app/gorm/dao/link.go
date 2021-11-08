package dao

import (
	"github.com/gofiber/fiber/v2"
	"math"
	"sblog/app/gorm/model"
	"sblog/app/logic"
	"sblog/system/common"
	"time"
)

type link struct {
}

var DefaultLinkDao = link{}

type linkSaveParam struct {
	ID     uint32 `json:"id" xml:"id" form:"id"`
	Title  string `json:"title" xml:"title" form:"title"`
	Url    string `json:"url" xml:"url" form:"url"`
	Status uint8  `json:"status" xml:"status" form:"status"`
	Sort   uint32 `json:"sort" xml:"sort" form:"sort"`
}

// Save 保存
func (d link) Save(ctx *fiber.Ctx) error {
	p := new(linkSaveParam)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var lk = model.Link{
		Title:      p.Title,
		Url:        p.Url,
		Status:     p.Status,
		Sort:       p.Sort,
		Updatetime: logic.LocalTime{Time: time.Now()},
	}

	//log.Println(lk)

	if p.ID > 0 {
		//编辑
		lk.ID = p.ID
		return lk.Updates(masterDB)
	}

	lk.Createtime = logic.LocalTime{Time: time.Now()}
	return lk.Create(masterDB)
}

// Delete 删除
func (d link) Delete(ctx *fiber.Ctx) error {
	type param struct {
		ID uint32 `json:"id" xml:"id" form:"id"`
	}
	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}
	var lk = model.Link{
		ID: p.ID,
	}
	return lk.Delete(masterDB)
}

// GetList 列表
func (d link) GetList(ctx *fiber.Ctx) ([]*model.Link, map[string]interface{}, error) {
	var links []*model.Link
	var err error
	var rp = make(map[string]interface{})

	p := new(listParam)
	if err = common.ParseParam(ctx, p); err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return links, rp, err
	}

	var lk = model.Link{}

	links, err = lk.GetList(masterDB, p.Page, p.PageSize)
	if err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return links, rp, err
	}

	count := lk.Count(masterDB)
	allPage := int(math.Ceil(float64(count) / float64(p.PageSize)))
	rp["count"] = count
	rp["all"] = allPage
	rp["page"] = p.Page
	rp["page_size"] = p.PageSize

	return links, rp, nil
}
