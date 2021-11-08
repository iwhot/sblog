package dao

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"math"
	"sblog/app/gorm/model"
	"sblog/app/logic"
	"sblog/system/common"
	"strconv"
	"time"
)

type slideCategory struct {
}

var DefaultSlideCategoryDao = slideCategory{}

// Save 保存
func (d slideCategory) Save(ctx *fiber.Ctx) error {
	type param struct {
		ID     uint32 `json:"id"`
		Title  string `json:"title"`
		Width  uint32 `json:"width"`
		Height uint32 `json:"height"`
		Remark string `json:"remark"`
		Attr   uint8  `json:"attr"`
	}

	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var sc = model.SlideCategory{
		ID:         p.ID,
		Title:      p.Title,
		Width:      p.Width,
		Height:     p.Height,
		Remark:     p.Remark,
		Attr:       p.Attr,
		Updatetime: logic.LocalTime{Time: time.Now()},
	}

	if p.ID > 0 {
		//编辑
		sc.ID = p.ID
		return sc.Update(masterDB)
	}

	sc.Createtime = logic.LocalTime{Time: time.Now()}
	return sc.Create(masterDB)
}

// GetList 获取列表
func (d slideCategory) GetList(ctx *fiber.Ctx) ([]*model.SlideCategory, map[string]interface{}, error) {
	var scs []*model.SlideCategory
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

	var sc = model.SlideCategory{}
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
func (d slideCategory) Delete(ctx *fiber.Ctx) error {
	type param struct {
		ID uint32 `json:"id" xml:"id" form:"id"`
	}
	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}
	var sc = model.SlideCategory{
		ID: p.ID,
	}
	return sc.Delete(masterDB)
}

// FindOne 获取一个分类
func (d slideCategory) FindOne(ctx *fiber.Ctx) (*model.SlideCategory, error) {
	var id, _ = strconv.Atoi(ctx.Params("id"))
	if id <= 0 {
		return nil, errors.New("获取失败")
	}

	log.Println("分类", id)

	var sc = model.SlideCategory{
		ID: uint32(id),
	}
	return sc.FindOne(masterDB)
}
