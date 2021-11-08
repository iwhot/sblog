package router

import (
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	v1 "sblog/app/controller/api/v1"
	"sblog/app/controller/frontend"
	"sblog/app/middleware"
)

func NewRouter(r *fiber.App) {
	//后台API
	apiV1 := r.Group("/api/v1", middleware.CheckLogin())
	{
		new(v1.Index).NewRouter(apiV1)         //首页接口
		new(v1.AuthGroup).NewRouter(apiV1)     //用户组
		new(v1.AuthRule).NewRouter(apiV1)      //菜单
		new(v1.Admin).NewRouter(apiV1)         //管理员
		new(v1.Config).NewRouter(apiV1)        //配置
		new(v1.Attachment).NewRouter(apiV1)    //附件
		new(v1.Link).NewRouter(apiV1)          //链接
		new(v1.Category).NewRouter(apiV1)      //分类
		new(v1.Article).NewRouter(apiV1)       //文章
		new(v1.SlideCategory).NewRouter(apiV1) //幻灯片分类
		new(v1.SlideBanner).NewRouter(apiV1)   //幻灯片
		new(v1.Tags).NewRouter(apiV1)          //标签
		new(v1.Menu).NewRouter(apiV1)          //目录
	}

	//前台套模板(为了能SEO放弃vue页面)
	front := r.Group("")
	{
		new(frontend.Index).NewRouter(front)
	}

	//最后的最后设置404页面，相当于所有路由都不走的时候就是404,这个不能在最前面设置否则所有都成404页面了
	r.Use(new(controller.Base).NotFound)
}
