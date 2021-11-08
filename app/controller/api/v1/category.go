package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

// Category 分类
type Category struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *Category) NewRouter(g fiber.Router) {
	g.Post("/category/index", c.Index)
	g.Post("/category/save", c.Save)
	g.Post("/category/delete/:id", c.Delete)
	g.Post("/category/get-parent-list", c.GetParentList)
	g.Post("/category/update-info", c.UpdateInfo)
	g.Post("/category/cascade", c.GetCategoryCascade)
	g.Post("/category/pid/:id", c.GetAllCategoryId)
}

// Index 首页
func (c *Category) Index(ctx *fiber.Ctx) error {
	ars, _ := dao.DefaultCategoryDao.GetList(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": ars,
	})
}

// Save 保存
func (c *Category) Save(ctx *fiber.Ctx) error {
	err := dao.DefaultCategoryDao.Save(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  fmt.Sprintf("保存失败,错误为 %v", err),
		})
	}
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "保存成功",
	})
}

// Delete 删除
func (c *Category) Delete(ctx *fiber.Ctx) error {
	err := dao.DefaultCategoryDao.Delete(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  fmt.Sprintf("删除失败,错误为 %v", err),
		})
	}

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "删除成功",
	})
}

// GetParentList 获取顶级分类
func (c *Category) GetParentList(ctx *fiber.Ctx) error {
	data, _ := dao.DefaultCategoryDao.GetParentCategory(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": data,
	})
}

// UpdateInfo 更改状态
func (c *Category) UpdateInfo(ctx *fiber.Ctx) error {
	err := dao.DefaultCategoryDao.UpdateInfo(ctx)
	//log.Println("信息", err)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  fmt.Sprintf("更新失败,错误为 %v", err),
		})
	}

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "更新成功",
	})
}

// GetCategoryCascade 获取分类的级联菜单
func (c *Category) GetCategoryCascade(ctx *fiber.Ctx) error {
	cas, _ := dao.DefaultCategoryDao.CascadeCategory(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "获取级联分类成功",
		"data": cas,
	})
}

// GetAllCategoryId 获取所有上级id
func (c *Category) GetAllCategoryId(ctx *fiber.Ctx) error {
	ret := dao.DefaultCategoryDao.GetAllCategoryId(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "获取上级分类成功",
		"data": ret,
	})
}
