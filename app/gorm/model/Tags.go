package model

import (
	"gorm.io/gorm"
	"sblog/app/logic"
	"sblog/system/page"
)

type Tags struct {
	ID         uint32          `gorm:"column:id;primaryKey" json:"id"`
	TagName    string          `gorm:"column:tag_name" json:"tag_name"`
	Createtime logic.LocalTime `gorm:"column:createtime" json:"createtime"`
	Updatetime logic.LocalTime `gorm:"column:updatetime" json:"updatetime"`
}

// GetAllTags 获取所有标签
func (m Tags) GetAllTags(db *gorm.DB) ([]*Tags, error) {
	var tgs []*Tags
	if err := db.Model(&m).Find(&tgs).Error; err != nil {
		return nil, err
	}
	return tgs, nil
}

// Create 创建
func (m Tags) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// Update 更新
func (m Tags) Update(db *gorm.DB) error {
	return db.Model(&m).Omit("Createtime").Save(m).Error
}

// Delete 删除
func (m Tags) Delete(db *gorm.DB) error {
	tx := db.Begin()

	var at = ArticleTags{
		TagsID: m.ID,
	}

	if err := at.DeleteByTagsID(tx); err != nil {
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

// GetList 获取列表
func (m Tags) GetList(db *gorm.DB, p, pz int) ([]*Tags, error) {
	var sld []*Tags
	var offset = page.GetOffset(p, pz)
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Offset(offset).Limit(pz).Order("id desc").Find(&sld).Error; err != nil {
		return nil, err
	}
	return sld, nil
}

// Count 获取统计
func (m Tags) Count(db *gorm.DB) int64 {
	var count int64
	if err := db.Model(&m).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

//todo 前台===============================================================================================

// GetTagsByNum 获取一定数量的标签
func (m Tags) GetTagsByNum(db *gorm.DB, num int) ([]*Tags, error) {
	var sld []*Tags

	if err := db.Model(&m).Limit(num).Order("id desc").Find(&sld).Error; err != nil {
		return nil, err
	}
	return sld, nil
}

// FindOne 获取一个标签
func (m Tags) FindOne(db *gorm.DB) (*Tags,error) {
	var tag *Tags
	if err := db.Model(&m).First(&tag).Error;err != nil{
		return tag, err
	}
	return tag,nil
}
