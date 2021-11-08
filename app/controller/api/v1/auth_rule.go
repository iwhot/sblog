package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

type AuthRule struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *AuthRule) NewRouter(g fiber.Router) {
	g.Post("/auth-rule/index", c.Index)
	g.Post("/auth-rule/save", c.Save)
	g.Post("/auth-rule/delete", c.Delete)
	g.Post("/auth-rule/delete-all", c.DeleteAll)
	g.Post("/auth-rule/get-menu", c.GetMyMenu)
	g.Post("/auth-rule/cascade-menu", c.CascadeMenu)
	g.Post("/auth-rule/get-menu-pid", c.GetMenuPID)
	g.Post("/auth-rule/update-menu-status", c.UpdateMenuStatus)
	g.Post("/auth-rule/update-menu-sort", c.UpdateMenuSort)
}

// Index 列表页，菜单不分页
func (c *AuthRule) Index(ctx *fiber.Ctx) error {
	ars, _ := dao.DefaultAuthRuleDao.GetList(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": ars,
	})
}

// Save 保存页面(添加、编辑)
func (c *AuthRule) Save(ctx *fiber.Ctx) error {
	err := dao.DefaultAuthRuleDao.Save(ctx)
	//log.Println("保存错误",err)
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
func (c *AuthRule) Delete(ctx *fiber.Ctx) error {
	err := dao.DefaultAuthRuleDao.Delete(ctx)
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

// DeleteAll 批量删除
func (c *AuthRule) DeleteAll(ctx *fiber.Ctx) error {
	err := dao.DefaultAuthRuleDao.DeleteAll(ctx)
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

// GetMyMenu 获取属于自己的菜单
func (c *AuthRule) GetMyMenu(ctx *fiber.Ctx) error {
	//获取用户
	admin, err := dao.DefaultAdminDao.GetAdminInfoByToken(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  fmt.Sprintf("get admin info err : %v", err),
		})
	}
	//取出对应用户的菜单
	menu, err := dao.DefaultAuthRuleDao.GetMyMenu(admin.ID)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  fmt.Sprintf("get menu err : %v", err),
		})
	}

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "获取菜单成功",
		"data": menu,
	})
}

// CascadeMenu 级联菜单,用于element无限极
func (c *AuthRule) CascadeMenu(ctx *fiber.Ctx) error {
	cas, _ := dao.DefaultAuthRuleDao.CascadeMenu(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "获取级联菜单成功",
		"data": cas,
	})
}

// GetMenuPID 获取菜单上级pid，递归获取 todo 后面补充
func (c *AuthRule) GetMenuPID(ctx *fiber.Ctx) error {
	pid := dao.DefaultAuthRuleDao.GetPid(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "获取菜单成功",
		"data": pid,
	})
}

// UpdateMenuStatus 更新菜单状态
func (c *AuthRule) UpdateMenuStatus(ctx *fiber.Ctx) error {
	err := dao.DefaultAuthRuleDao.UpdateStatus(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  "更新状态失败",
		})
	}

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "更新状态成功",
	})
}

// UpdateMenuSort 更新菜单状态
func (c *AuthRule) UpdateMenuSort(ctx *fiber.Ctx) error {
	err := dao.DefaultAuthRuleDao.UpdateSort(ctx)
	log.Println(err)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  "更新排序失败",
		})
	}

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "更新排序成功",
	})
}
