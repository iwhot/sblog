package dao

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"math"
	"sblog/app/gorm/model"
	"sblog/app/logic"
	"sblog/system/common"
	"strconv"
	"time"
)

type slideBanner struct {
}

var DefaultSlideBannerDao = slideBanner{}

// Save 保存
func (d slideBanner) Save(ctx *fiber.Ctx) error {
	type param struct {
		ID     uint32 `json:"id"`
		Title  string `json:"title"`
		Remark string `json:"remark"`
		Thumb  string `json:"thumb"`
		Url    string `json:"url"`
		Status uint8  `json:"status"`
		CID    uint32 `json:"cid"`
		Sort   uint32 `json:"sort"`
	}
	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}
	var cate = model.SlideCategory{
		ID: p.CID,
	}
	cts,err := cate.FindOne(masterDB)
	if err != nil{
		return err
	}

	var sb = model.SlideBanner{
		ID:         p.ID,
		Title:      p.Title,
		Remark:     p.Remark,
		Thumb:      p.Thumb,
		Url:        p.Url,
		Status:     p.Status,
		Updatetime: logic.LocalTime{Time: time.Now()},
		CID:        p.CID,
		Sort:       p.Sort,
	}

	if p.ID > 0 {
		//编辑
		sb.ID = p.ID
		return sb.Update(masterDB)
	}

	if cts.Attr == 1 && sb.BannerISExists(masterDB){
		return errors.New("该分类下只能上传一张图片")
	}

	sb.Createtime = logic.LocalTime{Time: time.Now()}
	return sb.Create(masterDB)
}

// GetList 获取列表
func (d slideBanner) GetList(ctx *fiber.Ctx) ([]*model.SlideBanner, map[string]interface{}, error) {
	var scs []*model.SlideBanner
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

	var id, _ = strconv.Atoi(ctx.Params("cid"))
	if id <= 0 {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return scs, rp, err
	}

	var sc = model.SlideBanner{CID: uint32(id)}
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
func (d slideBanner) Delete(ctx *fiber.Ctx) error {
	type param struct {
		ID uint32 `json:"id" xml:"id" form:"id"`
	}
	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}
	var sc = model.SlideBanner{
		ID: p.ID,
	}
	return sc.Delete(masterDB)
}
