package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"sblog/app/gorm/dao"
	"sblog/system/common"
)

// CheckLogin 登录验证
func CheckLogin() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//获取uri
		var nowUri = ctx.Request().URI().Path()
		//放行列表
		var urls = []string{"/api/v1/login", "/api/v1/captcha", "/api/v1/get-captcha"}
		//log.Println(string(nowUri))
		//如果在放行列表直接放行
		if bol, _ := common.InArray(string(nowUri), urls); bol == true {
			return ctx.Next()
		}

		var token = ctx.Get("Authorization")
		//没传token则返回错误
		if token == "" {
			return ctx.JSON(fiber.Map{
				"code": "401",
				"msg":  "请登录",
			})
		}

		//token校验
		err := dao.DefaultAdminDao.CheckToken(token)
		if err != nil {
			log.Println("token校验错误", err)
			return ctx.JSON(fiber.Map{
				"code": "401",
				"msg":  "Token过期,请登录",
			})
		}

		//校验通过，放行
		return ctx.Next()
	}
}
