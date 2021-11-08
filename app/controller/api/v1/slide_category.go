package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

type SlideCategory struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *SlideCategory) NewRouter(g fiber.Router) {
	g.Post("/slide-category/index", c.Index)
	g.Post("/slide-category/save", c.Save)
	g.Post("/slide-category/delete", c.Delete)
	g.Post("/slide-category/find-one/:id", c.FindOne)
}

// Index 首页
func (c *SlideCategory) Index(ctx *fiber.Ctx) error {
	aus, rp, _ := dao.DefaultSlideCategoryDao.GetList(ctx)
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
func (c *SlideCategory) Save(ctx *fiber.Ctx) error {
	err := dao.DefaultSlideCategoryDao.Save(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  fmt.Sprintf("保存失败,err = %v",err),
		})
	}
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "保存成功",
	})
}

// Delete 删除
func (c *SlideCategory) Delete(ctx *fiber.Ctx) error {
	err := dao.DefaultSlideCategoryDao.Delete(ctx)

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

// FindOne 获取一个分类
func (c *SlideCategory) FindOne(ctx *fiber.Ctx) error {
	info, _ := dao.DefaultSlideCategoryDao.FindOne(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": info,
	})
}
