package controller

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Base struct {
}

func (c *Base) NotFound(ctx *fiber.Ctx) error {
	//如果访问的是前台页面则返回json格式
	apiIndex := strings.Index(string(ctx.Request().URI().Path()), "/api")
	//log.Println(ctx.Request().URI().String(),apiIndex)
	if apiIndex == 0 {
		return ctx.JSON(fiber.Map{
			"code": "404",
			"msg":  "page not found!",
		})
	}

	//如果访问的其他则返回前台的404模板
	return ctx.Render("404",fiber.Map{
		//todo 写一些乱七八糟的东西
	})
}
