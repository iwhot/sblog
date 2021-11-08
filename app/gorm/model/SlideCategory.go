package model

import (
	"errors"
	"gorm.io/gorm"
	"sblog/app/logic"
	"sblog/system/page"
)

type SlideCategory struct {
	ID          uint32          `gorm:"column:id;primaryKey" json:"id"`
	Title       string          `gorm:"column:title" json:"title"`
	Width       uint32          `gorm:"column:width" json:"width"`
	Height      uint32          `gorm:"column:height" json:"height"`
	Remark      string          `gorm:"column:remark" json:"remark"`
	Attr        uint8           `gorm:"column:attr" json:"attr"`
	Createtime  logic.LocalTime `gorm:"column:createtime" gorm:"column:createtime" json:"createtime"`
	Updatetime  logic.LocalTime `gorm:"column:updatetime" gorm:"column:updatetime" json:"updatetime"`
	SlideBanner []SlideBanner   `gorm:"foreignKey:CID" json:"slide_banner"`
}

// Create 创建
func (m SlideCategory) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// Update 更新
func (m SlideCategory) Update(db *gorm.DB) error {
	return db.Model(&m).Omit("Createtime").Save(m).Error
}

// Delete 删除
func (m SlideCategory) Delete(db *gorm.DB) error {
	if (SlideBanner{}).BannerISExists(db) == false {
		return errors.New("分类下有幻灯片，不可删除")
	}

	return db.Delete(&m).Error
}

// GetList 获取列表
func (m SlideCategory) GetList(db *gorm.DB, p, pz int) ([]*SlideCategory, error) {
	var sld []*SlideCategory
	var offset = page.GetOffset(p, pz)
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Offset(offset).Limit(pz).Order("id desc").Find(&sld).Error; err != nil {
		return nil, err
	}
	return sld, nil
}

// Count 获取统计
func (m SlideCategory) Count(db *gorm.DB) int64 {
	var count int64
	if err := db.Model(&m).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// FindOne 获取一个分类
func (m SlideCategory) FindOne(db *gorm.DB) (*SlideCategory, error) {
	var sc *SlideCategory
	if err := db.Model(&m).Where("id=?",m.ID).First(&sc).Error; err != nil {
		return nil, err
	}
	return sc, nil
}

//todo 前台部分========================================================================================

// GetCategoryList 获取分类列表
func (m SlideCategory) GetCategoryList(db *gorm.DB) (*SlideCategory, error) {
	var sc *SlideCategory
	if err := db.Model(&m).Where("id=?",m.ID).Preload("SlideBanner", "`status` = ?", 1).First(&sc).Error; err != nil {
		return nil, err
	}
	return sc, nil
}
