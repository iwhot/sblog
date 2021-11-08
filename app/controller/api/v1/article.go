package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

// Article 文章
type Article struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *Article) NewRouter(g fiber.Router) {
	g.Post("/article/index", c.Index)
	g.Post("/article/info/:id", c.Edit)
	g.Post("/article/save", c.Save)
	g.Post("/article/delete/:id", c.Delete)
	g.Post("/article/real-delete/:id", c.RealDelete)
	g.Post("/article/update-info", c.UpdateInfo)
}

// Index 首页
func (c *Article) Index(ctx *fiber.Ctx) error {
	aus, rp, _ := dao.DefaultArticleDao.GetList(ctx)
	//log.Println(err)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": map[string]interface{}{
			"list": aus,
			"page": rp,
		},
	})
}

// Edit 编辑页面
func (c *Article) Edit(ctx *fiber.Ctx) error {
	info, _ := dao.DefaultArticleDao.FindOne(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": info,
	})
}

// Save 保存
func (c *Article) Save(ctx *fiber.Ctx) error {
	err := dao.DefaultArticleDao.Save(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  fmt.Sprintf("保存失败,错误:%v",err),
		})
	}
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "保存成功",
	})
}

// Delete 删除
func (c *Article) Delete(ctx *fiber.Ctx) error {
	err := dao.DefaultArticleDao.Delete(ctx)

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

// RealDelete 真删
func (c *Article) RealDelete(ctx *fiber.Ctx) error {
	err := dao.DefaultArticleDao.RealDelete(ctx)

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

// UpdateInfo 更新某个信息
func (c *Article) UpdateInfo(ctx *fiber.Ctx) error {
	err := dao.DefaultArticleDao.UpdateInfo(ctx)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  "操作失败",
		})
	}

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "操作成功",
	})
}
