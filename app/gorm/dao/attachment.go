package dao

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"math"
	"os"
	"sblog/app/gorm/model"
	"sblog/app/logic"
	"sblog/system/common"
	"strconv"
	"time"
)

type attachment struct {
}

var DefaultAttachmentDao = attachment{}

// Upload 上传
func (d attachment) Upload(ctx *fiber.Ctx, uid uint32) (string, error) {

	ufa, err := logic.UploadFile(ctx)
	if err != nil {
		return "", err
	}

	//保存文件到数据库
	var atm = model.Attachment{
		UserID:     uid,
		Url:        ufa.Path,
		Filesize:   uint32(ufa.Size),
		Width:      ufa.Width,
		Height:     ufa.Height,
		Createtime: logic.LocalTime{Time: time.Now()},
		Desc:       ufa.FileName,
		Mold:       ufa.Mold,
		Updatetime: logic.LocalTime{Time: time.Now()},
	}

	if err = atm.Create(masterDB); err != nil {
		//创建失败，删除文件
		if common.IsExists("." + ufa.Path) {
			_ = os.Remove("." + ufa.Path)
		}
		return "", err
	}

	return ufa.Path, nil
}

// Index 首页
func (d attachment) Index(ctx *fiber.Ctx) ([]*model.Attachment, map[string]interface{}, error) {
	var ats []*model.Attachment
	var err error
	var rp = make(map[string]interface{})

	p := new(listParam)
	if err = common.ParseParam(ctx, p); err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return ats, rp, err
	}

	var at = model.Attachment{
		Mold: ctx.Params("mold"),
	}

	log.Println("类型", ctx.Params("mold"))

	ats, err = at.GetList(masterDB, p.Page, p.PageSize)
	if err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return ats, rp, err
	}

	count := at.Count(masterDB)
	allPage := int(math.Ceil(float64(count) / float64(p.PageSize)))
	rp["count"] = count
	rp["all"] = allPage
	rp["page"] = p.Page
	rp["page_size"] = p.PageSize

	return ats, rp, nil
}

// Delete 删除
func (d attachment) Delete(ctx *fiber.Ctx) error {
	type param struct {
		ID uint32 `json:"id" xml:"id" form:"id"`
	}
	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var atm = model.Attachment{
		ID: p.ID,
	}
	return atm.Delete(masterDB)
}

// Download 下载
func (d attachment) Download(ctx *fiber.Ctx) (string, string) {
	var id, _ = strconv.Atoi(ctx.Params("id"))

	if id == 0 {
		return "", ""
	}

	var amt = model.Attachment{
		ID: uint32(id),
	}

	amts, _ := amt.FindOne(masterDB)
	return amts.Url, amts.Desc
}

// Edit 编辑图片
func (d attachment) Edit(ctx *fiber.Ctx) error {

	return nil
}
