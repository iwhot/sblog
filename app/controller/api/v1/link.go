package v1

import (
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

// Link 链接
type Link struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *Link) NewRouter(g fiber.Router) {
	g.Post("/link/index", c.Index)
	g.Post("/link/save", c.Save)
	g.Post("/link/delete", c.Delete)
}

// Index 首页
func (c *Link) Index(ctx *fiber.Ctx) error {
	aus, rp, _ := dao.DefaultLinkDao.GetList(ctx)
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
func (c *Link) Save(ctx *fiber.Ctx) error {
	err := dao.DefaultLinkDao.Save(ctx)
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
func (c *Link) Delete(ctx *fiber.Ctx) error {
	err := dao.DefaultLinkDao.Delete(ctx)

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
