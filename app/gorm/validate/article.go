package validate

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"sblog/system/common"
)

// articleValidate 简单验证
type articleValidate struct {
}

var DefaultArticleValidate = articleValidate{}

type ArticleParam struct {
	ID       uint32 `json:"id"`
	CID      uint32 `json:"cid"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Content  string `json:"content"`
	Attr     uint8  `json:"attr"`
	Status   uint8  `json:"status"`
	IsTop    uint8  `json:"is_top"`
	IsNice   uint8  `json:"is_nice"`
	IsIndex  uint8  `json:"is_index"`
	Sort     uint32 `json:"sort"`
	State    uint8  `json:"state"`
	Thumb    string `json:"thumb"`
	ThumbExt string `json:"thumb_ext"`
	SeoTitle string `json:"seo_title"`
	SeoDesc  string `json:"seo_desc"`
	SeoKey   string `json:"seo_key"`
	Tags     string `json:"tags"`
	IsDel    uint8  `json:"is_del"`
	Path     string `json:"path"`
}

// CheckSave 简单验证一下保存，先快速写完，懒得用验证器了
func (a articleValidate) CheckSave(ctx *fiber.Ctx) (*ArticleParam, error) {
	p := new(ArticleParam)
	if err := common.ParseParam(ctx, p); err != nil {
		return p, err
	}

	if p.Title == "" {
		return p, errors.New("标题不能为空")
	}

	if p.Content == "" {
		return p, errors.New("内容不能为空")
	}

	if p.Attr == 0 && p.CID == 0 {
		return p, errors.New("分类不能为空")
	}

	if p.Attr == 0 && p.Thumb == "" {
		return p, errors.New("缩略图必须上传")
	}

	return p, nil
}
