package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/spf13/viper"
	"log"
	"sblog/app/router"
	"sblog/system/common"
	"sblog/system/global"
	"time"
)

func init() {
	//载入全局配置
	global.Init()
}

func main() {
	//设置模板引擎
	engine := html.New("./views", ".html")
	//模板中引入函数
	engine.AddFunc("InArray",common.InArray2)
	engine.AddFunc("StringToTemplate",common.StringToTemplate)
	//设置读写时间
	rd, _ := time.ParseDuration(viper.GetString("server.readTimeout"))
	wd, _ := time.ParseDuration(viper.GetString("server.writeTimeout"))
	app := fiber.New(fiber.Config{
		Views: engine, //使用模板引擎
		StrictRouting: true,//严格的路由
		ReadTimeout:    rd,//读取超时时间
		WriteTimeout:   wd,//写入超时时间
		ReadBufferSize: 1 << 20,//
	})
	//设置一些东西
	app.Use(recover.New()) //恢复
	app.Use(logger.New())  //日志
	//app.Use(cache.New())   //缓存
	//app.Use(cors.New())//跨域设置,vue部分开了代理可不用这个
	app.Static("/storage", "./storage") //设置静态路径
	//载入路由
	router.NewRouter(app)
	var addr = viper.GetString("server.host") + ":" + viper.GetString("server.port")
	log.Fatal(app.Listen(addr))
}
