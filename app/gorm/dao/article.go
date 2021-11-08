package dao

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"math"
	"sblog/app/gorm/model"
	"sblog/app/gorm/validate"
	"sblog/app/logic"
	"sblog/system/common"
	"strconv"
	"strings"
	"time"
)

type article struct {
}

var DefaultArticleDao = article{}

// Save 保存
func (d article) Save(ctx *fiber.Ctx) error {
	p, err := validate.DefaultArticleValidate.CheckSave(ctx)
	if err != nil {
		return err
	}

	var tgs []model.Tags
	if p.Tags != "" {
		//取出所有标签
		allTag, err := (model.Tags{}).GetAllTags(masterDB)
		if err != nil {
			return err
		}

		//分割字符串
		t := strings.Split(p.Tags, ",")
		for _, v := range t {
			tag1 := GetThisTagId(v, allTag) //检测标签是否存在
			var tg2 model.Tags
			if tag1 != nil { //标签存在则更新
				tg2 = model.Tags{
					ID:         tag1.ID,
					TagName:    v,
					Createtime: tag1.Createtime,
					Updatetime: logic.LocalTime{Time: time.Now()},
				}
			} else { //标签不存在则添加
				tg2 = model.Tags{
					TagName:    v,
					Createtime: logic.LocalTime{Time: time.Now()},
					Updatetime: logic.LocalTime{Time: time.Now()},
				}
			}

			tgs = append(tgs, tg2)
		}
	}

	var at = model.Article{
		CID:        p.CID,
		Title:      p.Title,
		Desc:       p.Desc,
		Content:    p.Content,
		Attr:       p.Attr,
		Status:     p.Status,
		IsTop:      p.IsTop,
		IsNice:     p.IsNice,
		IsIndex:    p.IsIndex,
		Sort:       p.Sort,
		Updatetime: logic.LocalTime{Time: time.Now()},
		State:      p.State,
		Thumb:      p.Thumb,
		ThumbExt:   p.ThumbExt,
		SeoTitle:   p.SeoTitle,
		SeoDesc:    p.SeoDesc,
		SeoKey:     p.SeoKey,
		IsDel:      p.IsDel,
		Tags:       tgs,
		Path:       p.Path,
	}

	if p.ID > 0 {
		at.ID = p.ID
		at.Updatetime = logic.LocalTime{Time: time.Now()}
		//更新
		return at.Update(masterDB)
	} else {
		//获取用户
		var token = ctx.Get("Authorization")
		var ue = model.UserToken{Token: token}
		u, err := ue.FindOneToken(masterDB)
		if err != nil {
			return err
		}

		at.UserID = u.UserID
		at.Token = utils.UUIDv4()
		at.Createtime = logic.LocalTime{Time: time.Now()}
		return at.Create(masterDB)
	}
}

// Delete 软删除
func (d article) Delete(ctx *fiber.Ctx) error {
	var id, _ = strconv.Atoi(ctx.Params("id"))
	if id <= 0 {
		return errors.New("删除失败")
	}

	var at = model.Article{ID: uint32(id)}
	var mp = make(map[string]interface{})
	mp["is_del"] = 1 //放回收站
	return at.UpdateByMap(masterDB, mp)
}

// RealDelete 真实删除
func (d article) RealDelete(ctx *fiber.Ctx) error {
	var id, _ = strconv.Atoi(ctx.Params("id"))
	if id <= 0 {
		return errors.New("删除失败")
	}
	var at = model.Article{ID: uint32(id)}
	return at.Delete(masterDB)
}

// FindOne 获取一条记录
func (d article) FindOne(ctx *fiber.Ctx) (*model.Article, error) {
	var id, _ = strconv.Atoi(ctx.Params("id"))
	if id <= 0 {
		return nil, errors.New("文章不存在")
	}

	var at = model.Article{
		ID: uint32(id),
	}
	return at.FindOne(masterDB)
}

// GetList 获取列表
func (d article) GetList(ctx *fiber.Ctx) ([]*model.Article, map[string]interface{}, error) {
	type param struct {
		Page     int    `json:"page" xml:"page" form:"page"`
		PageSize int    `json:"page_size" xml:"page_size" form:"page_size"`
		Count    int    `json:"count" xml:"count" form:"count"`
		Start    string `json:"start" xml:"start" form:"start"`
		End      string `json:"end" xml:"end" form:"end"`
		State    uint8  `json:"state" xml:"state" form:"state"`
		Title    string `json:"title" xml:"title" form:"title"`
		CateName string `json:"cate_name" xml:"cate_name" form:"cate_name"`
		Username string `json:"username" xml:"username" form:"username"`
		IsDel    uint8  `json:"is_del" xml:"is_del" form:"is_del"`
		Attr     uint8  `json:"attr" xml:"attr" form:"attr"`
	}
	var arts []*model.Article
	var err error
	var rp = make(map[string]interface{})

	p := new(param)
	if err = common.ParseParam(ctx, p); err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return arts, rp, err
	}

	var at = model.Article{
		State:    p.State,
		Title:    p.Title,
		Category: model.Category{Name: p.CateName},
		Admin:    model.Admin{Username: p.Username},
		IsDel:    p.IsDel,
		Attr:     p.Attr,
	}

	//log.Println("传参",p.State)
	arts, err = at.GetArticleList(masterDB, p.Page, p.PageSize, p.Start, p.End)
	if err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return arts, rp, err
	}

	count := at.Count(masterDB, p.Start, p.End)
	allPage := int(math.Ceil(float64(count) / float64(p.PageSize)))
	rp["count"] = count
	rp["all"] = allPage
	rp["page"] = p.Page
	rp["page_size"] = p.PageSize

	return arts, rp, nil
}

func GetThisTagId(tg string, allTag []*model.Tags) *model.Tags {

	for _, v := range allTag {
		if v.TagName == tg {
			return v
		}
	}

	return nil
}

// UpdateInfo 更新部分信息
func (d article) UpdateInfo(ctx *fiber.Ctx) error {
	var mp = make(map[string]interface{})
	if err := common.ParseParam(ctx, &mp); err != nil {
		return err
	}

	if mp["id"] == nil {
		return errors.New("文章不存在")
	}

	return (model.Article{}).UpdateInfo(masterDB, mp)
}
