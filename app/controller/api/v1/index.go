package v1

import (
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

type Index struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *Index) NewRouter(g fiber.Router) {
	g.Get("/index", c.Index)             //首页
	g.Post("/login", c.Login)            //登陆
	g.Post("/get-captcha", c.GetCaptcha) //获取验证码信息
	g.Get("/captcha", c.Captcha)         //获取验证码图片
	g.Post("/user-info", c.Info)          //获取用户信息
}

// Index 接口首页
func (c *Index) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"code": "000000",
		"msg":  "请求成功1",
	})
}

// Login 登录页面
func (c *Index) Login(ctx *fiber.Ctx) error {
	adminInfo, token, err := dao.DefaultAdminDao.Login(ctx)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  fmt.Sprintf("Login err : %v", err),
		})
	}

	var user = make(map[string]interface{})
	user["id"] = adminInfo.ID
	user["username"] = adminInfo.Username
	user["nickname"] = adminInfo.Nickname
	user["avatar"] = adminInfo.Avatar
	user["email"] = adminInfo.Email

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "登陆成功",
		"data": map[string]interface{}{
			"token": token,
			"user":  user,
		},
	})
}

// GetCaptcha 获取验证码，参考地址 https://studygolang.com/articles/30279?fr=sidebar
func (c *Index) GetCaptcha(ctx *fiber.Ctx) error {
	sess := captcha.NewLen(5)
	return ctx.JSON(fiber.Map{
		"code": 1,
		"msg":  "successfully",
		"data": map[string]string{
			"sess": sess,
			"url":  "/captcha?sess=" + sess,
		},
	})
}

// Captcha 验证码
func (c *Index) Captcha(ctx *fiber.Ctx) error {
	var sess = ctx.FormValue("sess")
	if sess == "" {
		return ctx.SendStatus(404)
	}
	ctx.Set("Content-Type", "image/png")
	err := captcha.WriteImage(ctx.Response().BodyWriter(), sess, 160, 40)
	return err
}

// Info 获取用户信息
func (c *Index) Info(ctx *fiber.Ctx) error {
	admin, _ := dao.DefaultAdminDao.GetAdminInfoByToken(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "获取成功",
		"data": admin,
	})
}
