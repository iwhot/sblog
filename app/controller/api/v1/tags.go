package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

type Tags struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *Tags) NewRouter(g fiber.Router) {
	g.Post("/tags/index", c.Index)
	g.Post("/tags/save", c.Save)
	g.Post("/tags/delete", c.Delete)
}

// Index 首页
func (c *Tags) Index(ctx *fiber.Ctx) error {
	aus, rp, _ := dao.DefaultTagsDao.GetList(ctx)
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
func (c *Tags) Save(ctx *fiber.Ctx) error {
	err := dao.DefaultTagsDao.Save(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  fmt.Sprintf("保存失败,err = %v", err),
		})
	}
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "保存成功",
	})
}

// Delete 删除
func (c *Tags) Delete(ctx *fiber.Ctx) error {
	err := dao.DefaultTagsDao.Delete(ctx)

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
