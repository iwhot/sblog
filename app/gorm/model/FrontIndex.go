package model

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"sblog/system/page"
	"strings"
)

type TabArticleList struct {
	ID         uint32   `json:"id"`
	Title      string   `json:"title"`
	Desc       string   `json:"desc"`
	Thumb      string   `json:"thumb"`
	Path       string   `json:"path"`
	ThumbExt   string   `json:"thumb_ext"`
	Createtime string   `json:"createtime"`
	Like       uint32   `json:"like"`
	Read       uint32   `json:"read"`
	Pics       []string `json:"pics"`
	IsTop      uint8    `json:"is_top"`
}

type TabArticle struct {
	Category []string            `json:"cate"`
	Article  [][]*TabArticleList `json:"article"`
}

// GetTabArticle 获取选项卡文章
func GetTabArticle(db *gorm.DB, num int) (TabArticle, error) {
	var ta TabArticle
	//取出菜单
	var ms []*Menu
	if err := db.Where("`group`=? AND `status` = ?", "center", 1).Order("sort desc,id desc").Find(&ms).Error; err != nil {
		return ta, err
	}

	var urls []string
	for _, v := range ms {
		urls = append(urls, strings.TrimLeft(v.Url, "/"))
	}

	//通过地址获取分类
	var cas []*Category
	if err := db.Where("`key` IN ? AND `status` = ?", urls, 1).Find(&cas).Error; err != nil {
		return ta, err
	}

	var category []string

	var clist [][]*TabArticleList
	for _, v := range cas {
		category = append(category, v.Name)

		var tbs []*TabArticleList
		_ = db.Model(Article{}).Select("`id`,`title`,`desc`,`thumb`,`path`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=? AND `cid` = ?", 0, 0, 2, 1, v.ID).Limit(num).Order("sort desc,id desc").Scan(&tbs).Error
		clist = append(clist, tbs)
	}

	ta.Category = category
	ta.Article = clist

	return ta, nil
}

// GetNiceArticleList 推荐文章
func GetNiceArticleList(db *gorm.DB, num int) ([]*TabArticleList, error) {
	var tbs []*TabArticleList
	if err := db.Model(Article{}).Select("`id`,`title`,`desc`,`thumb`,`path`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=? AND `is_nice` = ?", 0, 0, 2, 1, 1).Limit(num).Order("sort desc,id desc").Scan(&tbs).Error; err != nil {
		return nil, err
	}
	return tbs, nil
}

// GetTopArticleList 置顶文章
func GetTopArticleList(db *gorm.DB, num int) ([]*TabArticleList, error) {
	var tbs []*TabArticleList
	if err := db.Model(Article{}).Select("`id`,`title`,`desc`,`thumb`,`path`,`thumb_ext`,`createtime`,`like`,`read`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=? AND `is_top` = ?", 0, 0, 2, 1, 1).Limit(num).Order("sort desc,id desc").Scan(&tbs).Error; err != nil {
		return nil, err
	}

	for _, v := range tbs {
		if v.ThumbExt != "" {
			arr := strings.Split(v.ThumbExt, ",")
			v.Pics = append(v.Pics, arr...)
		}
	}

	return tbs, nil
}

// GetNewArticleList 获取最新列表
func GetNewArticleList(db *gorm.DB, num int) ([]*TabArticleList, error) {
	var tbs []*TabArticleList
	if err := db.Model(Article{}).Select("`id`,`title`,`desc`,`thumb`,`path`,`thumb_ext`,`createtime`,`like`,`read`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=?", 0, 0, 2, 1).Limit(num).Order("id desc").Scan(&tbs).Error; err != nil {
		return nil, err
	}

	for _, v := range tbs {
		if v.ThumbExt != "" {
			arr := strings.Split(v.ThumbExt, ",")
			v.Pics = append(v.Pics, arr...)
		}
	}

	return tbs, nil
}

// GetClickArticleList 点击排行
func GetClickArticleList(db *gorm.DB, num int) ([]*TabArticleList, error) {
	var tbs []*TabArticleList
	if err := db.Model(Article{}).Select("`id`,`title`,`desc`,`thumb`,`path`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=?", 0, 0, 2, 1).Limit(num).Order("`read` desc").Scan(&tbs).Error; err != nil {
		return nil, err
	}

	return tbs, nil
}

// GetLickArticleList 喜欢排行
func GetLickArticleList(db *gorm.DB, num int) ([]*TabArticleList, error) {
	var tbs []*TabArticleList
	if err := db.Model(Article{}).Select("`id`,`title`,`desc`,`thumb`,`path`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=?", 0, 0, 2, 1).Limit(num).Order("`like` desc").Scan(&tbs).Error; err != nil {
		return nil, err
	}

	return tbs, nil
}

// GetNiceArticleList2  点击排行
func GetNiceArticleList2(db *gorm.DB, num int) ([]*TabArticleList, error) {
	var tbs []*TabArticleList
	if err := db.Model(Article{}).Select("`id`,`title`,`desc`,`thumb`,`path`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=?", 0, 0, 2, 1).Limit(num).Order("`is_nice` desc").Scan(&tbs).Error; err != nil {
		return nil, err
	}

	return tbs, nil
}

// GetArticlePrevNext 获取文章上下页
func GetArticlePrevNext(db *gorm.DB, id uint32) (*TabArticleList, *TabArticleList) {
	var a1 []*TabArticleList
	var a2 []*TabArticleList
	//获取上页
	_ = db.Model(Article{}).Select("`id`,`title`,`desc`,`thumb`,`path`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=? AND `id`<?", 0, 0, 2, 1, id).Order("id desc").Limit(1).Scan(&a1).Error
	//获取下页
	_ = db.Model(Article{}).Select("`id`,`title`,`desc`,`thumb`,`path`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=? AND `id`>?", 0, 0, 2, 1, id).Order("id asc").Limit(1).Scan(&a2).Error

	var b1 *TabArticleList
	var b2 *TabArticleList

	if len(a1) > 0 {
		b1 = a1[0]
	}

	if len(a2) > 0 {
		b2 = a2[0]
	}

	return b1, b2
}

// GetArticleByCategoryID 获取某一分类几篇文章
func GetArticleByCategoryID(db *gorm.DB, id uint32, num int) ([]*TabArticleList, error) {
	var tbs []*TabArticleList
	if err := db.Model(Article{}).Select("`id`,`title`,`desc`,`thumb`,`path`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=? AND `cid` = ?", 0, 0, 2, 1, id).Limit(num).Order("`sort` desc , `id` desc").Scan(&tbs).Error; err != nil {
		return nil, err
	}

	return tbs, nil
}

// GetCategoryArticleList 获取分类下文章列表
func GetCategoryArticleList(db *gorm.DB, id uint32, p, pz int) ([]*TabArticleList, error) {
	var tbs []*TabArticleList
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	var offset = page.GetOffset(p, pz)

	//获取分类
	var cts []uint32
	cts = append(cts,id)
	var cats = (Category{ID: id}).GetAllCategoryId(db,cts)

	if err := tx.Model(Article{}).Offset(offset).Limit(pz).Select("`id`,`title`,`desc`,`thumb`,`path`,`thumb_ext`,`createtime`,`like`,`read`,`is_top`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=? AND `cid` IN ?", 0, 0, 2, 1, cats).Order("`createtime` desc,`id` desc").Scan(&tbs).Error; err != nil {
		return nil, err
	}

	for _, v := range tbs {
		if v.ThumbExt != "" {
			arr := strings.Split(v.ThumbExt, ",")
			v.Pics = append(v.Pics, arr...)
		}
	}

	return tbs, nil

}

// GetCategoryArticleCount 获取统计
func GetCategoryArticleCount(db *gorm.DB, id uint32) int64 {
	var count int64
	//获取分类
	var cts []uint32
	cts = append(cts,id)
	var cats = (Category{ID: id}).GetAllCategoryId(db,cts)

	if err := db.Model(&Article{}).Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=? AND `cid` IN ?", 0, 0, 2, 1, cats).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// GetTagsArticleList 获取标签文章列表
func GetTagsArticleList(db *gorm.DB, id uint32, p, pz int) ([]*TabArticleList, error) {
	var tbs []*TabArticleList
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	var offset = page.GetOffset(p, pz)
	var prefix = viper.GetString("database.prefix")
	if err := tx.Model(Article{}).Offset(offset).Limit(pz).Joins("left join "+prefix+"article_tags on "+prefix+"article_tags.article_id = "+prefix+"article.id").Select("`id`,`title`,`desc`,`thumb`,`path`,`thumb_ext`,`createtime`,`like`,`read`,`is_top`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=? AND `tags_id` = ?", 0, 0, 2, 1, id).Order("`createtime` desc,`id` desc").Scan(&tbs).Error; err != nil {
		return nil, err
	}

	for _, v := range tbs {
		if v.ThumbExt != "" {
			arr := strings.Split(v.ThumbExt, ",")
			v.Pics = append(v.Pics, arr...)
		}
	}

	return tbs, nil
}

// GetTagsArticleCount 获取标签文章列表
func GetTagsArticleCount(db *gorm.DB, id uint32) int64 {
	var count int64
	var prefix = viper.GetString("database.prefix")
	if err := db.Model(&Article{}).Joins("left join "+prefix+"article_tags  on "+prefix+"article_tags.article_id = "+prefix+"article.id").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=? AND `tags_id` = ?", 0, 0, 2, 1, id).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// GetSearchArticleList 获取搜索列表
func GetSearchArticleList(db *gorm.DB, keywords string, p, pz int) ([]*TabArticleList, error) {
	var tbs []*TabArticleList
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	var offset = page.GetOffset(p, pz)

	tx = tx.Model(Article{})
	if keywords != "" {
		tx = tx.Where("`title` like ? ", "%"+keywords+"%")
	}

	if err := tx.Offset(offset).Limit(pz).Select("`id`,`title`,`desc`,`thumb`,`path`,`thumb_ext`,`createtime`,`like`,`read`,`is_top`").Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=?", 0, 0, 2, 1).Order("`createtime` desc,`id` desc").Scan(&tbs).Error; err != nil {
		return nil, err
	}

	for _, v := range tbs {
		if v.ThumbExt != "" {
			arr := strings.Split(v.ThumbExt, ",")
			v.Pics = append(v.Pics, arr...)
		}
	}

	return tbs, nil
}

// GetSearchArticleCount 获取统计
func GetSearchArticleCount(db *gorm.DB, keywords string) int64 {
	var count int64
	var tx = db.Model(&Article{})
	if keywords != "" {
		tx = tx.Where("`title` like ? ", "%"+keywords+"%")
	}
	if err := tx.Where("`attr`=? AND `is_del`=? AND `state`=? AND `status`=?", 0, 0, 2, 1).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// GetActive 获取选中
func GetActive(db *gorm.DB,id uint32) string {
	var cat,_ = (Category{ID: id}).FindOne(db)
	if cat.ID == 0{
		return ""
	}

	if cat.PID != 0{
		return GetActive(db,cat.PID)
	}

	return cat.Key
}