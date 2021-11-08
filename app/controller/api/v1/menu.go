package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

type Menu struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *Menu) NewRouter(g fiber.Router) {
	g.Post("/menu/index", c.Index)
	g.Post("/menu/save", c.Save)
	g.Post("menu/delete/:id", c.Delete)
	g.Post("/menu/get-parent-list", c.GetParentList)
	g.Post("/menu/update-info", c.UpdateInfo)
	g.Post("/menu/cascade", c.GetCategoryCascade)
	g.Post("/menu/pid/:id", c.GetAllCategoryId)
}

// Index 首页
func (c *Menu) Index(ctx *fiber.Ctx) error {
	aus, _ := dao.DefaultMenuDao.GetList(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": aus,
	})
}

// Save 保存
func (c *Menu) Save(ctx *fiber.Ctx) error {
	err := dao.DefaultMenuDao.Save(ctx)
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
func (c *Menu) Delete(ctx *fiber.Ctx) error {
	err := dao.DefaultMenuDao.Delete(ctx)

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

// GetParentList 获取顶级分类
func (c *Menu) GetParentList(ctx *fiber.Ctx) error {
	data, _ := dao.DefaultMenuDao.GetParentMenu(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": data,
	})
}

// GetCategoryCascade 获取分类的级联菜单
func (c *Menu) GetCategoryCascade(ctx *fiber.Ctx) error {
	cas, _ := dao.DefaultMenuDao.CascadeMenu(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "获取级联分类成功",
		"data": cas,
	})
}

// GetAllCategoryId 获取所有上级id
func (c *Menu) GetAllCategoryId(ctx *fiber.Ctx) error {
	ret := dao.DefaultMenuDao.GetAllMenuId(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "获取上级分类成功",
		"data": ret,
	})
}

// UpdateInfo 更改状态
func (c *Menu) UpdateInfo(ctx *fiber.Ctx) error {
	err := dao.DefaultMenuDao.UpdateInfo(ctx)
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
