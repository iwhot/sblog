package frontend

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
	"sblog/app/logic"
)

type Index struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *Index) NewRouter(g fiber.Router) {
	g.Get("/", c.Index)
	g.Get("/link", c.GetLink)
	g.Get("/tags/:id", c.GetTagsArticle)
	g.Get("/tags-all", c.GetTagsAll)
	g.Get("/search", c.Search)
	g.Post("/article-like/:id", c.ArticleLike)
	g.Get("/:detail", c.Detail)
}

// Index 接口首页
func (c *Index) Index(ctx *fiber.Ctx) error {
	var headerMenu = dao.DefaultFrontendIndex.GetMenu(ctx, "top","/")                   //顶部菜单
	var setting = logic.GetAllSetting()                                             //配置
	var banner, _ = dao.DefaultFrontendIndex.GetBanner(ctx, 1)                      //banner
	var articleTop, _ = dao.DefaultFrontendIndex.GetIndexArticle(ctx, 2)            //顶部文章
	var GetTabArticle, _ = dao.DefaultFrontendIndex.GetTabArticle(ctx, 7)           //选项卡文章
	var GetNiceArticleList, _ = dao.DefaultFrontendIndex.GetNiceArticleList(ctx, 6) //推荐文章
	var GetTopArticleList, _ = dao.DefaultFrontendIndex.GetTopArticleList(ctx, 1)   //一个置顶
	var GetNewArticleList, _ = dao.DefaultFrontendIndex.GetNewArticleList(ctx, 5)   //最新

	var GetClickArticleList, _ = dao.DefaultFrontendIndex.GetClickArticleList(ctx, 10) //点击排行
	var GetLickArticleList, _ = dao.DefaultFrontendIndex.GetLickArticleList(ctx, 10)   //站长推荐
	var GetNiceArticleList2, _ = dao.DefaultFrontendIndex.GetNiceArticleList2(ctx, 10) //猜你喜欢

	var banner1, _ = dao.DefaultFrontendIndex.GetBanner(ctx, 2) //banner1
	var banner2, _ = dao.DefaultFrontendIndex.GetBanner(ctx, 3) //banner2
	var banner3, _ = dao.DefaultFrontendIndex.GetBanner(ctx, 4) //banner3

	var GetIndexLinkList, _ = dao.DefaultFrontendIndex.GetIndexLinkList(ctx, 8)

	//log.Printf("%+v",setting)
	return ctx.Render("index", fiber.Map{
		"HeaderMenu":          headerMenu,
		"Setting":             setting,
		"Banner":              banner,
		"Banner1":             banner1,
		"Banner2":             banner2,
		"Banner3":             banner3,
		"IsBanner":            len(banner.SlideBanner) > 0,
		"IsBanner1":           len(banner1.SlideBanner) > 0,
		"IsBanner2":           len(banner2.SlideBanner) > 0,
		"IsBanner3":           len(banner3.SlideBanner) > 0,
		"ArticleTop":          articleTop,
		"GetTabArticle":       GetTabArticle,
		"GetNiceArticleList":  GetNiceArticleList,
		"GetTopArticleList":   GetTopArticleList,
		"GetNewArticleList":   GetNewArticleList,
		"GetClickArticleList": GetClickArticleList,
		"GetLickArticleList":  GetLickArticleList,
		"GetNiceArticleList2": GetNiceArticleList2,
		"GetIndexLinkList":    GetIndexLinkList,
	})
}

// GetLink 获取链接
func (c *Index) GetLink(ctx *fiber.Ctx) error {
	var headerMenu = dao.DefaultFrontendIndex.GetMenu(ctx, "top","")
	var setting = logic.GetAllSetting()
	var link, _ = dao.DefaultFrontendIndex.GetLinks(ctx)

	return ctx.Render("link", fiber.Map{
		"HeaderMenu": headerMenu,
		"Link":       link,
		"Setting":    setting,
	})
}

// GetTagsArticle 获取标签文章
func (c *Index) GetTagsArticle(ctx *fiber.Ctx) error {
	var headerMenu = dao.DefaultFrontendIndex.GetMenu(ctx, "top","")
	var setting = logic.GetAllSetting()

	var GetClickArticleList, _ = dao.DefaultFrontendIndex.GetClickArticleList(ctx, 10) //点击排行
	var GetLickArticleList, _ = dao.DefaultFrontendIndex.GetLickArticleList(ctx, 10)   //站长推荐
	var GetNiceArticleList2, _ = dao.DefaultFrontendIndex.GetNiceArticleList2(ctx, 10) //猜你喜欢
	var banner2, _ = dao.DefaultFrontendIndex.GetBanner(ctx, 3)                        //banner2
	var banner3, _ = dao.DefaultFrontendIndex.GetBanner(ctx, 4)                        //banner3
	var GetTagsByNum, _ = dao.DefaultFrontendIndex.GetTagsByNum(ctx, 16)               //标签

	//获取列表
	var articleList, pHtml = dao.DefaultFrontendIndex.GetTagsArticleList(ctx, 10)

	var tagInfo, _ = dao.DefaultFrontendIndex.GetOneTagInfo(ctx)

	return ctx.Render("article/list_tags", fiber.Map{
		"HeaderMenu":          headerMenu,
		"Setting":             setting,
		"Banner2":             banner2,
		"Banner3":             banner3,
		"IsBanner2":           len(banner2.SlideBanner) > 0,
		"IsBanner3":           len(banner3.SlideBanner) > 0,
		"GetClickArticleList": GetClickArticleList,
		"GetLickArticleList":  GetLickArticleList,
		"GetNiceArticleList2": GetNiceArticleList2,
		"GetTagsByNum":        GetTagsByNum,
		"ArticleList":         articleList,
		"PageHtml":            pHtml,
		"TagInfo":             tagInfo,
	})
}

// GetTagsAll 获取所有标签
func (c *Index) GetTagsAll(ctx *fiber.Ctx) error {
	var headerMenu = dao.DefaultFrontendIndex.GetMenu(ctx, "top","")
	var setting = logic.GetAllSetting()
	var tags, _ = dao.DefaultFrontendIndex.GetTags(ctx)

	return ctx.Render("tags", fiber.Map{
		"HeaderMenu": headerMenu,
		"Tags":       tags,
		"Setting":    setting,
	})
}

// Search 获取搜索
func (c *Index) Search(ctx *fiber.Ctx) error {
	var headerMenu = dao.DefaultFrontendIndex.GetMenu(ctx, "top","")
	var setting = logic.GetAllSetting()

	var GetClickArticleList, _ = dao.DefaultFrontendIndex.GetClickArticleList(ctx, 10) //点击排行
	var GetLickArticleList, _ = dao.DefaultFrontendIndex.GetLickArticleList(ctx, 10)   //站长推荐
	var GetNiceArticleList2, _ = dao.DefaultFrontendIndex.GetNiceArticleList2(ctx, 10) //猜你喜欢
	var banner2, _ = dao.DefaultFrontendIndex.GetBanner(ctx, 3)                        //banner2
	var banner3, _ = dao.DefaultFrontendIndex.GetBanner(ctx, 4)                        //banner3
	var GetTagsByNum, _ = dao.DefaultFrontendIndex.GetTagsByNum(ctx, 16)               //标签

	//获取列表
	var articleList, pHtml = dao.DefaultFrontendIndex.GetSearchArticleList(ctx, 10)

	return ctx.Render("article/list_search", fiber.Map{
		"HeaderMenu":          headerMenu,
		"Setting":             setting,
		"Banner2":             banner2,
		"Banner3":             banner3,
		"IsBanner2":           len(banner2.SlideBanner) > 0,
		"IsBanner3":           len(banner3.SlideBanner) > 0,
		"GetClickArticleList": GetClickArticleList,
		"GetLickArticleList":  GetLickArticleList,
		"GetNiceArticleList2": GetNiceArticleList2,
		"GetTagsByNum":        GetTagsByNum,
		"ArticleList":         articleList,
		"PageHtml":            pHtml,
		"SearchName":          ctx.Query("keywords"),
	})
}

// Detail 详情页面
func (c *Index) Detail(ctx *fiber.Ctx) error {
	var setting = logic.GetAllSetting()

	//获取详情
	article, _ := dao.DefaultFrontendIndex.GetArticleByPath(ctx)

	//单页
	if article != nil && article.Attr == 1 {
		var active = "/" + article.Path
		var headerMenu = dao.DefaultFrontendIndex.GetMenu(ctx, "top",active)
		return ctx.Render("page/about", fiber.Map{
			"HeaderMenu": headerMenu,
			"Article":    article,
			"Setting":    setting,
			"Active":    active,
		})
	}

	var GetClickArticleList, _ = dao.DefaultFrontendIndex.GetClickArticleList(ctx, 10) //点击排行
	var GetLickArticleList, _ = dao.DefaultFrontendIndex.GetLickArticleList(ctx, 10)   //站长推荐
	var GetNiceArticleList2, _ = dao.DefaultFrontendIndex.GetNiceArticleList2(ctx, 10) //猜你喜欢
	var banner2, _ = dao.DefaultFrontendIndex.GetBanner(ctx, 3)                        //banner2
	var banner3, _ = dao.DefaultFrontendIndex.GetBanner(ctx, 4)                        //banner3
	var GetTagsByNum, _ = dao.DefaultFrontendIndex.GetTagsByNum(ctx, 16)               //标签

	//文章
	if article != nil && article.Attr == 0 {
		//log.Println(article.Admin)
		//获取导航
		var crumbs, _ = dao.DefaultFrontendIndex.GetCrumbsByID(ctx, article.Category.ID)
		var prev, next = dao.DefaultFrontendIndex.GetArticlePrevNext(ctx, article.ID)
		var GetArticleByCategoryID, _ = dao.DefaultFrontendIndex.GetArticleByCategoryID(ctx, article.Category.ID, 6)

		//阅读数+1
		_ = dao.DefaultFrontendIndex.ArticleReadInc(ctx, article.ID)
		article.Read += 1
		//log.Println("自增",err)
		//log.Println(prev, next)
		var active = "/" + dao.DefaultFrontendIndex.GetActive(ctx,article.CID)
		var headerMenu = dao.DefaultFrontendIndex.GetMenu(ctx, "top",active)
		//普通文章
		return ctx.Render("article/info", fiber.Map{
			"HeaderMenu":             headerMenu,
			"Article":                article,
			"Setting":                setting,
			"Crumbs":                 crumbs,
			"Prev":                   prev,
			"IsPrev":                 prev == nil,
			"Next":                   next,
			"IsNext":                 next == nil,
			"GetArticleByCategoryID": GetArticleByCategoryID,
			"Banner2":                banner2,
			"Banner3":                banner3,
			"IsBanner2":              len(banner2.SlideBanner) > 0,
			"IsBanner3":              len(banner3.SlideBanner) > 0,
			"GetClickArticleList":    GetClickArticleList,
			"GetLickArticleList":     GetLickArticleList,
			"GetNiceArticleList2":    GetNiceArticleList2,
			"GetTagsByNum":           GetTagsByNum,
		})
	}

	//分类列表
	category, _ := dao.DefaultFrontendIndex.GetCategoryInfoByKey(ctx)
	if category != nil {
		//获取下级
		var GetAllChildList, _ = dao.DefaultFrontendIndex.GetAllChildList(ctx, category.ID)
		//获取分类下文章列表
		var articleList, pHtml = dao.DefaultFrontendIndex.GetCategoryArticleList(ctx, category.ID, 10)
		var active = "/" + dao.DefaultFrontendIndex.GetActive(ctx,category.ID)

		var headerMenu = dao.DefaultFrontendIndex.GetMenu(ctx, "top",active)
		//分类
		return ctx.Render("article/list", fiber.Map{
			"HeaderMenu":          headerMenu,
			"Category":            category,
			"Setting":             setting,
			"Banner2":             banner2,
			"Banner3":             banner3,
			"IsBanner2":           len(banner2.SlideBanner) > 0,
			"IsBanner3":           len(banner3.SlideBanner) > 0,
			"GetClickArticleList": GetClickArticleList,
			"GetLickArticleList":  GetLickArticleList,
			"GetNiceArticleList2": GetNiceArticleList2,
			"GetTagsByNum":        GetTagsByNum,

			"GetAllChildList": GetAllChildList,
			"ArticleList":     articleList,
			"PageHtml":        pHtml,
		})
	}

	//什么都不是跳转到首页
	return ctx.Redirect("/", http.StatusMovedPermanently)
}

func (c *Index) ArticleLike(ctx *fiber.Ctx) error {
	_ = dao.DefaultFrontendIndex.ArticleLike(ctx)
	return ctx.JSON(fiber.Map{
		"code": "0",
		"msg":  "点赞成功",
	})
}
