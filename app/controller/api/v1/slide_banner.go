package v1

import (
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

type SlideBanner struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *SlideBanner) NewRouter(g fiber.Router) {
	g.Post("/slide-banner/index/:cid", c.Index)
	g.Post("/slide-banner/save", c.Save)
	g.Post("/slide-banner/delete", c.Delete)
}

// Index 首页
func (c *SlideBanner) Index(ctx *fiber.Ctx) error {
	aus, rp, _ := dao.DefaultSlideBannerDao.GetList(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": map[string]interface{}{
			"list": aus,
			"page": rp,
		},
	})
}

// Save 保存
func (c *SlideBanner) Save(ctx *fiber.Ctx) error {
	err := dao.DefaultSlideBannerDao.Save(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  "保存失败",
		})
	}
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "保存成功",
	})
}

// Delete 删除
func (c *SlideBanner) Delete(ctx *fiber.Ctx) error {
	err := dao.DefaultSlideBannerDao.Delete(ctx)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  "删除失败",
		})
	}

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "删除成功",
	})
}

