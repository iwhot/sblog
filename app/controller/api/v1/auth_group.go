package v1

import (
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

type AuthGroup struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *AuthGroup) NewRouter(g fiber.Router) {
	g.Post("/auth-group/index", c.Index)
	g.Post("/auth-group/save", c.Save)
	g.Post("/auth-group/delete", c.Delete)
	g.Post("/auth-group/delete-all", c.DeleteAll)
	g.Post("/auth-group/get-all", c.GetAllList)
}

// Index 列表页
func (c *AuthGroup) Index(ctx *fiber.Ctx) error {
	ret, rp, _ := dao.DefaultAuthGroupDao.GetList(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": map[string]interface{}{
			"list": ret,
			"page": rp,
		},
	})
}

// Save 保存页面(添加、编辑)
func (c *AuthGroup) Save(ctx *fiber.Ctx) error {
	err := dao.DefaultAuthGroupDao.Save(ctx)
	//log.Println(err)
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
func (c *AuthGroup) Delete(ctx *fiber.Ctx) error {
	err := dao.DefaultAuthGroupDao.Delete(ctx)

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
func (c *AuthGroup) DeleteAll(ctx *fiber.Ctx) error {
	err := dao.DefaultAuthGroupDao.DeleteAll(ctx)

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

// GetAllList 获取所有角色组
func (c *AuthGroup) GetAllList(ctx *fiber.Ctx) error {
	ags, _ := dao.DefaultAuthGroupDao.GetAllGroup(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "获取成功",
		"data": ags,
	})
}
