package model

import (
	"gorm.io/gorm"
	"sblog/app/logic"
	"sblog/system/page"
)

type Link struct {
	ID         uint32          `gorm:"column:id;primaryKey" json:"id"`
	Title      string          `gorm:"column:title" json:"title"`
	Url        string          `gorm:"column:url" json:"url"`
	Createtime logic.LocalTime `gorm:"column:createtime" json:"createtime"`
	Updatetime logic.LocalTime `gorm:"column:updatetime" json:"updatetime"`
	Status     uint8           `gorm:"column:status" json:"status"`
	Sort       uint32          `gorm:"column:sort" json:"sort"`
}

// Create 创建
func (m Link) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// Updates 更新
func (m Link) Updates(db *gorm.DB) error {
	return db.Model(&m).Omit("Createtime").Save(m).Error
}

// Delete 删除
func (m Link) Delete(db *gorm.DB) error {
	return db.Delete(&m).Error
}

// Count 统计
func (m Link) Count(db *gorm.DB) int64 {
	var count int64

	if err := db.Model(&m).Count(&count).Error; err != nil {
		return 0
	}

	return count
}

// GetList 后台获取列表
func (m Link) GetList(db *gorm.DB, p, pz int) ([]*Link, error) {
	var lks []*Link
	var offset = page.GetOffset(p, pz)
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Offset(offset).Limit(pz).Order("sort desc,id desc").Find(&lks).Error; err != nil {
		return nil, err
	}
	return lks, nil
}

//todo 前台==============================================================================================================

// GetIndexLink 获取首页连接
func (m Link) GetIndexLink(db *gorm.DB, num int) ([]*Link, error) {
	var lks []*Link
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Where("`status`=?", 1).Limit(num).Order("sort desc,id desc").Find(&lks).Error; err != nil {
		return nil, err
	}
	return lks, nil
}

// GetAllLink 获取所有链接
func (m Link) GetAllLink(db *gorm.DB) ([]*Link, error) {
	var lks []*Link
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Where("`status`=?", 1).Order("sort desc,id desc").Find(&lks).Error; err != nil {
		return nil, err
	}
	return lks, nil
}
