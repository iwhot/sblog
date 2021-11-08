package v1

import (
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

type Config struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *Config) NewRouter(g fiber.Router) {
	g.Post("/config/index", c.Index)
	g.Post("/config/save", c.Save)
}

// Index 获取配置列表
func (c *Config) Index(ctx *fiber.Ctx) error {
	cfg, _ := dao.DefaultConfigsDao.GetList(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": cfg,
	})
}

// Save 保存配置
func (c *Config) Save(ctx *fiber.Ctx) error {
	_ = dao.DefaultConfigsDao.Save(ctx)

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "保存成功",
	})
}
