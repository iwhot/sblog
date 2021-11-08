package v1

import (
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

type Admin struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *Admin) NewRouter(g fiber.Router) {
	g.Post("/admin/index", c.Index)
	g.Post("/admin/save", c.Save)
	g.Post("/admin/save2", c.Save2)
	g.Post("/admin/delete", c.Delete)
	g.Post("/admin/delete-all", c.DeleteAll)
	g.Post("/admin/get-group-id", c.GetMyGroupID)
}

// Index 列表页
func (c *Admin) Index(ctx *fiber.Ctx) error {
	aus, rp, _ := dao.DefaultAdminDao.GetAdminList(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": map[string]interface{}{
			"list": aus,
			"page": rp,
		},
	})
}

// Save 保存页面(添加、编辑)
func (c *Admin) Save(ctx *fiber.Ctx) error {
	err := dao.DefaultAdminDao.Save(ctx)

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

// Save2 保存页面(添加、编辑)
func (c *Admin) Save2(ctx *fiber.Ctx) error {
	err := dao.DefaultAdminDao.Save2(ctx)

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
func (c *Admin) Delete(ctx *fiber.Ctx) error {
	err := dao.DefaultAdminDao.Delete(ctx)

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
func (c *Admin) DeleteAll(ctx *fiber.Ctx) error {
	err := dao.DefaultAdminDao.DeleteAll(ctx)

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

// GetMyGroupID 获取分组id
func (c *Admin) GetMyGroupID(ctx *fiber.Ctx) error {
	gids, _ := dao.DefaultAdminDao.GetMyGroupID(ctx)

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "获取成功",
		"data": gids,
	})
}
