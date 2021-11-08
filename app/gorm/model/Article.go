package model

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sblog/app/logic"
	"sblog/system/common"
	"sblog/system/page"
	"strconv"
)

type Article struct {
	ID         uint32          `gorm:"column:id;primaryKey" json:"id"`
	Token      string          `gorm:"column:token" json:"token"`
	CID        uint32          `gorm:"column:cid" json:"cid"`
	Category   Category        `gorm:"foreignKey:CID" json:"category"` //cid作为外键
	UserID     uint32          `gorm:"column:user_id" json:"user_id"`
	Admin      Admin           `gorm:"foreignKey:UserID" json:"admin"` //user_id作为外键
	Title      string          `gorm:"column:title" json:"title"`
	Desc       string          `gorm:"column:desc" json:"desc"`
	Content    string          `gorm:"column:content" json:"content"`
	Like       uint32          `gorm:"column:like" json:"like"`
	Read       uint32          `gorm:"column:read" json:"read"`
	Attr       uint8           `gorm:"column:attr" json:"attr"`
	Status     uint8           `gorm:"column:status" json:"status"`
	IsTop      uint8           `gorm:"column:is_top" json:"is_top"`
	IsNice     uint8           `gorm:"column:is_nice" json:"is_nice"`
	IsIndex    uint8           `gorm:"column:is_index" json:"is_index"`
	Sort       uint32          `gorm:"column:sort" json:"sort"`
	Createtime logic.LocalTime `gorm:"column:createtime" json:"createtime"`
	Updatetime logic.LocalTime `gorm:"column:updatetime" json:"updatetime"`
	State      uint8           `gorm:"column:state" json:"state"`
	Thumb      string          `gorm:"column:thumb" json:"thumb"`         //缩略图
	ThumbExt   string          `gorm:"column:thumb_ext" json:"thumb_ext"` //多图
	SeoTitle   string          `gorm:"column:seo_title" json:"seo_title"`
	SeoDesc    string          `gorm:"column:seo_desc" json:"seo_desc"`
	SeoKey     string          `gorm:"column:seo_key" json:"seo_key"`
	IsDel      uint8           `gorm:"column:is_del" json:"is_del"`
	Tags       []Tags          `gorm:"many2many:article_tags;"` //多对多
	Path       string          `gorm:"column:path" json:"path"` //路径
}

// Create 创建
func (m Article) Create(db *gorm.DB) error {
	tx := db.Begin()

	if m.SeoTitle == "" {
		m.SeoTitle = m.Title
	}

	if m.SeoKey == "" {
		m.SeoKey = m.Title
	}

	if m.SeoDesc == "" {
		m.SeoDesc = m.Title
	}

	//如果摘要没写就从内容中复制
	if m.Desc == ""{
		m.Desc = common.SubString(common.TrimHtml(m.Content), 0, 120) //取100个字
	}

	if err := tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	//路径为空
	if m.Path == "" {
		var pt = strconv.Itoa(int(m.ID)) + ".html"
		if err := tx.Model(&m).Where("id=?", m.ID).Updates(map[string]interface{}{"path": pt}).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}

	//存在路径
	var at *Article
	_ = tx.Model(&m).Where("path=?", m.Path).First(&at)

	if at.ID != 0 {
		tx.Rollback()
		return errors.New("路径已经存在")
	}

	tx.Commit()
	return nil
}

// Update 更新
func (m Article) Update(db *gorm.DB) error {
	tx := db.Begin()

	if m.SeoTitle == "" {
		m.SeoTitle = m.Title
	}

	if m.SeoKey == "" {
		m.SeoKey = m.Title
	}

	if m.SeoDesc == "" {
		m.SeoDesc = m.Title
	}

	//路径为空
	if m.Path == "" {
		m.Path = strconv.Itoa(int(m.ID)) + ".html"
	}else{
		//存在路径
		var at *Article
		_ = tx.Model(&m).Where("`path`=? AND `id` <> ?", m.Path,m.ID).First(&at)
		//log.Println(at.ID,at.Title)
		if at.ID != 0 {
			tx.Rollback()
			return errors.New("路径已经存在")
		}
	}

	//如果摘要没写就从内容中复制
	if m.Desc == ""{
		m.Desc = common.SubString(common.TrimHtml(m.Content), 0, 120) //取100个字
	}

	if err:= tx.Model(&m).Omit("Createtime", "Token","Category","Admin").Save(m).Error;err != nil{
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// UpdateByMap 根据map更新
func (m Article) UpdateByMap(db *gorm.DB, mp map[string]interface{}) error {
	return db.Model(&m).Updates(mp).Error
}

// IsArticleByCategory 检测分类下是否有文章
func (m Article) IsArticleByCategory(db *gorm.DB) bool {
	var at *Article
	if err := db.Model(&m).Where("cid=?", m.CID).First(&at).Error; err != nil {
		return false
	}
	return true
}

// GetArticleList 获取列表
// @Params p int 页码
// @Params pz int 一页条数
// @Params start string 开始时间
// @Params end string 结束时间
func (m Article) GetArticleList(db *gorm.DB, p, pz int, start, end string) ([]*Article, error) {
	var ats []*Article
	var tx = db.Session(&gorm.Session{PrepareStmt: true})

	//属性：0文章，1单页
	tx = tx.Where("attr=? AND is_del = ?", m.Attr, m.IsDel)

	//1草稿箱，2发布
	if m.State == 0 {
		tx = tx.Where("state > ?", 0)
	} else {
		tx = tx.Where("state = ?", m.State)
	}

	//搜索标题
	if m.Title != "" {
		tx = tx.Where("title like ?", "%"+m.Title+"%")
	}

	//搜索创建时间
	if start != "" {
		tx = tx.Where("createtime >= ?", start)
	}

	if end != "" {
		tx = tx.Where("createtime <= ?", end)
	}

	//搜索用户名
	if m.Admin.Username != "" {
		tx = tx.Preload("Admin", "username like ?", "%"+m.Admin.Username+"%")
	} else {
		tx = tx.Preload("Admin")
	}

	//搜索分类名
	if m.Category.Name != "" {
		tx = tx.Preload("Category", "name like ?", "%"+m.Category.Name+"%")
	} else {
		tx = tx.Preload("Category")
	}

	var offset = page.GetOffset(p, pz)

	//预加载全部
	if err := tx.Offset(offset).Limit(pz).Order("sort desc,id desc").Find(&ats).Error; err != nil {
		return ats, err
	}

	return ats, nil
}

// FindOne 获取一条记录
func (m Article) FindOne(db *gorm.DB) (*Article, error) {
	var at *Article
	if err := db.Where("id=?", m.ID).Preload(clause.Associations).First(&at).Error; err != nil {
		return at, err
	}
	return at, nil
}

// Delete 真实删除
func (m Article) Delete(db *gorm.DB) error {
	var tx = db.Begin()
	//删除文章标签
	var at = ArticleTags{
		ArticleID: m.ID,
	}

	if err := at.DeleteByArticleID(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// Count 统计
func (m Article) Count(tx *gorm.DB, start, end string) int64 {
	var count int64
	//属性：0文章，1单页
	tx = tx.Where("attr=?", m.Attr)

	//0回收站，1草稿箱，2发布
	if m.State == 0 {
		tx = tx.Where("state > ?", 0)
	} else {
		tx = tx.Where("state = ?", 0)
	}

	//搜索标题
	if m.Title != "" {
		tx = tx.Where("title like ?", "%"+m.Title+"%")
	}

	//搜索创建时间
	if start != "" {
		tx = tx.Where("createtime >= ?", start)
	}

	if end != "" {
		tx = tx.Where("createtime <= ?", end)
	}

	//搜索用户名
	if m.Admin.Username != "" {
		tx = tx.Preload("Admin", "username like ?", "%"+m.Admin.Username+"%")
	} else {
		tx = tx.Preload("Admin")
	}

	//搜索分类名
	if m.Category.Name != "" {
		tx = tx.Preload("Category", "name like ?", "%"+m.Category.Name+"%")
	} else {
		tx = tx.Preload("Category")
	}

	if err := tx.Model(&m).Order("sort desc,id desc").Count(&count).Error; err != nil {
		return 0
	}

	return count
}

// UpdateInfo 更新信息
func (m Article) UpdateInfo(db *gorm.DB, mp map[string]interface{}) error {
	return db.Model(&m).Where("id=?", mp["id"]).Updates(mp).Error
}

//todo 前台======================================================================================

// GetArticleInfoByPath 获取文章
func (m Article) GetArticleInfoByPath(db *gorm.DB) (*Article, error) {
	var at *Article
	if err := db.Where("`is_del` = ? AND `path` = ? AND `status` = ?", 0, m.Path, 1).Preload(clause.Associations).First(&at).Error; err != nil {
		return nil, err
	}
	return at, nil
}

// GetArticleByNum 获取固定数量文章
func (m Article) GetArticleByNum(db *gorm.DB, num int,order string) ([]*Article, error) {
	var ats []*Article
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	tx = tx.Model(&m)
	if m.IsTop > 0 {
		tx = tx.Where("is_top=?", 1)
	}

	if m.IsNice > 0 {
		tx = tx.Where("is_nice=?", 1)
	}

	if m.IsIndex > 0 {
		tx = tx.Where("is_index=?", 1)
	}

	if err := tx.Where("`state`=? AND `status` = ? AND `attr` = ?",2,1,0).Limit(num).Order(order).Find(&ats).Error; err != nil {
		return nil, err
	}

	return ats, nil
}

//todo 前台======================================================================================================================================

// ArticleReadInc 自增1
func (m Article) ArticleReadInc(db *gorm.DB) error {
	return db.Model(&m).Update("read",gorm.Expr("`read`+?",1)).Error
}

// ArticleLikeInc 自增1
func (m Article) ArticleLikeInc(db *gorm.DB) error {
	return db.Model(&m).Update("like",gorm.Expr("`like`+?",1)).Error
}