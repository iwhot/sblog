package dao

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"html/template"
	"sblog/app/gorm/model"
	"sblog/app/logic"
	"sblog/system/page"
	"strconv"
	"time"
)

type frontIndex struct {
}

var DefaultFrontendIndex = frontIndex{}

// GetMenu 菜单
func (d frontIndex) GetMenu(ctx *fiber.Ctx, group,active string) []*model.TreeMenu {
	var m = model.Menu{Group: group}
	return m.GetMenu(masterDB,active)
}

// GetArticleByPath 获取文章或单页
func (d frontIndex) GetArticleByPath(ctx *fiber.Ctx) (*model.Article, error) {
	var at = model.Article{Path: ctx.Params("detail")}
	return at.GetArticleInfoByPath(masterDB)
}

// GetCategoryInfoByKey 获取分类信息
func (d frontIndex) GetCategoryInfoByKey(ctx *fiber.Ctx) (*model.Category, error) {
	var cat = model.Category{Key: ctx.Params("detail")}
	return cat.GetCategoryByKey(masterDB)
}

// GetBanner 获取banner
func (d frontIndex) GetBanner(ctx *fiber.Ctx, id uint32) (*model.SlideCategory, error) {
	var sc = model.SlideCategory{ID: id}
	return sc.GetCategoryList(masterDB)
}

// GetIndexArticle 获取首页顶部文章
func (d frontIndex) GetIndexArticle(ctx *fiber.Ctx, num int) ([]*model.Article, error) {
	var at = model.Article{
		IsIndex: 1,
	}
	return at.GetArticleByNum(masterDB, num, "sort desc,id desc")
}

// GetTabArticle 获取选项卡文章
func (d frontIndex) GetTabArticle(ctx *fiber.Ctx, num int) (model.TabArticle, error) {
	return model.GetTabArticle(masterDB, num)
}

// GetNiceArticleList 获取推荐文章
func (d frontIndex) GetNiceArticleList(ctx *fiber.Ctx, num int) ([]*model.TabArticleList, error) {
	return model.GetNiceArticleList(masterDB, num)
}

// GetTopArticleList 获取置顶文章
func (d frontIndex) GetTopArticleList(ctx *fiber.Ctx, num int) ([]*model.TabArticleList, error) {
	return model.GetTopArticleList(masterDB, num)
}

// GetNewArticleList 最新文章
func (d frontIndex) GetNewArticleList(ctx *fiber.Ctx, num int) ([]*model.TabArticleList, error) {
	return model.GetNewArticleList(masterDB, num)
}

// GetClickArticleList 点击排行
func (d frontIndex) GetClickArticleList(ctx *fiber.Ctx, num int) ([]*model.TabArticleList, error) {
	return model.GetClickArticleList(masterDB, num)
}

// GetLickArticleList 猜你喜欢
func (d frontIndex) GetLickArticleList(ctx *fiber.Ctx, num int) ([]*model.TabArticleList, error) {
	return model.GetLickArticleList(masterDB, num)
}

// GetNiceArticleList2 站长推荐
func (d frontIndex) GetNiceArticleList2(ctx *fiber.Ctx, num int) ([]*model.TabArticleList, error) {
	return model.GetNiceArticleList2(masterDB, num)
}

// GetIndexLinkList 获取友情链接
func (d frontIndex) GetIndexLinkList(ctx *fiber.Ctx, num int) ([]*model.Link, error) {
	var lk = model.Link{}
	return lk.GetIndexLink(masterDB, num)
}

// GetCrumbsByID 获取文章的面包屑导航
func (d frontIndex) GetCrumbsByID(ctx *fiber.Ctx, id uint32) ([]*model.Category, error) {
	var cat = model.Category{
		ID: id,
	}

	var ret []*model.Category
	return cat.GetCrumbsByID(masterDB, ret)
}

// GetArticlePrevNext 获取上下页
func (d frontIndex) GetArticlePrevNext(ctx *fiber.Ctx, id uint32) (*model.TabArticleList, *model.TabArticleList) {
	return model.GetArticlePrevNext(masterDB, id)
}

// GetArticleByCategoryID 获取某一分类文章
func (d frontIndex) GetArticleByCategoryID(ctx *fiber.Ctx, id uint32, num int) ([]*model.TabArticleList, error) {
	return model.GetArticleByCategoryID(masterDB, id, num)
}

// GetTagsByNum 标签云
func (d frontIndex) GetTagsByNum(ctx *fiber.Ctx, num int) ([]*model.Tags, error) {
	var tag = model.Tags{}
	return tag.GetTagsByNum(masterDB, num)
}

// ArticleReadInc 文章自增1
func (d frontIndex) ArticleReadInc(ctx *fiber.Ctx, id uint32) error {
	var art = model.Article{ID: id}
	return art.ArticleReadInc(masterDB)
}

// ArticleLikeInc 文章自增1
func (d frontIndex) ArticleLikeInc(ctx *fiber.Ctx, id uint32) error {
	var art = model.Article{ID: id}
	return art.ArticleReadInc(masterDB)
}

// GetLinks 获取所有链接
func (d frontIndex) GetLinks(ctx *fiber.Ctx) ([]*model.Link, error) {
	var lk = model.Link{}
	return lk.GetAllLink(masterDB)
}

// GetTags 获取所有标签
func (d frontIndex) GetTags(ctx *fiber.Ctx) ([]*model.Tags, error) {
	var tg = model.Tags{}
	return tg.GetAllTags(masterDB)
}

// GetAllChildList 获取所有下级分类(只获取一级)
func (d frontIndex) GetAllChildList(ctx *fiber.Ctx, id uint32) ([]*model.Category, error) {
	var cat = model.Category{ID: id}
	return cat.GetAllChildList(masterDB)
}

// GetCategoryArticleList 获取分类下列表
func (d frontIndex) GetCategoryArticleList(ctx *fiber.Ctx, id uint32, pz int) ([]*model.TabArticleList, template.HTML) {
	var p int
	p, _ = strconv.Atoi(ctx.Query("p", "0"))

	if p < 1 {
		p = 1
	}

	var tab, _ = model.GetCategoryArticleList(masterDB, id, p, pz)
	var count = model.GetCategoryArticleCount(masterDB, id)
	var pages = page.GetPagination(p, pz, int(count), "")
	return tab, pages
}

// GetOneTagInfo 获取一个标签
func (d frontIndex) GetOneTagInfo(ctx *fiber.Ctx) (*model.Tags, error) {
	var id, _ = strconv.Atoi(ctx.Params("id", "0"))
	var tag = model.Tags{ID: uint32(id)}
	return tag.FindOne(masterDB)
}

// GetTagsArticleList 获取标签列表
func (d frontIndex) GetTagsArticleList(ctx *fiber.Ctx, pz int) ([]*model.TabArticleList, template.HTML) {
	var p int
	var id, _ = strconv.Atoi(ctx.Params("id", "0"))
	p, _ = strconv.Atoi(ctx.Query("p", "0"))

	if p < 1 {
		p = 1
	}

	var tab, _ = model.GetTagsArticleList(masterDB, uint32(id), p, pz)
	var count = model.GetTagsArticleCount(masterDB, uint32(id))
	var pages = page.GetPagination(p, pz, int(count), "")
	return tab, pages
}

// GetSearchArticleList 获取搜索列表
func (d frontIndex) GetSearchArticleList(ctx *fiber.Ctx, pz int) ([]*model.TabArticleList, template.HTML) {
	var p int
	var keywords = ctx.Query("keywords", "")
	p, _ = strconv.Atoi(ctx.Query("p", "0"))

	if p < 1 {
		p = 1
	}

	//log.Println("关键字",keywords)

	var tab, _ = model.GetSearchArticleList(masterDB, keywords, p, pz)
	var count = model.GetSearchArticleCount(masterDB, keywords)
	var pages = page.GetPagination(p, pz, int(count), "keywords="+keywords)
	return tab, pages
}

// ArticleLike 点赞
func (d frontIndex) ArticleLike(ctx *fiber.Ctx) error {
	var articleId,_ = strconv.Atoi(ctx.Params("id","0"))
	if articleId == 0 {
		return errors.New("点赞失败")
	}
	var ip = ctx.IP()
	var al = model.ArticleLike{
		ArticleID: uint32(articleId),
		Ip:         ip,
		Createtime: logic.LocalTime{Time:time.Now()},
	}

	return al.Create(masterDB)
}

// GetActive 获取选中
func (d frontIndex) GetActive(ctx *fiber.Ctx,id uint32) string {
	return model.GetActive(masterDB, id)
}