package model

import (
	"gorm.io/gorm"
	"sblog/app/logic"
	"sblog/system/page"
)

type SlideBanner struct {
	ID         uint32          `gorm:"column:id;primaryKey" json:"id"`
	Title      string          `gorm:"column:title" json:"title"`
	Remark     string          `gorm:"column:remark" json:"remark"`
	Thumb      string          `gorm:"column:thumb" json:"thumb"`
	Url        string          `gorm:"column:url" json:"url"`
	Status     uint8           `gorm:"column:status" json:"status"`
	Createtime logic.LocalTime `gorm:"column:createtime" gorm:"column:createtime" json:"createtime"`
	Updatetime logic.LocalTime `gorm:"column:updatetime" gorm:"column:updatetime" json:"updatetime"`
	CID        uint32          `gorm:"column:cid" json:"cid"`
	Sort       uint32          `gorm:"column:sort" json:"sort"`
}

// Create 创建
func (m SlideBanner) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// Update 更新
func (m SlideBanner) Update(db *gorm.DB) error {
	return db.Model(&m).Save(m).Error
}

// Delete 删除
func (m SlideBanner) Delete(db *gorm.DB) error {
	return db.Delete(&m).Error
}

// UpdateInfo 更新某些字段
func (m SlideBanner) UpdateInfo(db *gorm.DB, mp map[string]interface{}) error {
	return db.Model(&m).Where("id=?", mp["id"]).Updates(mp).Error
}

// GetList 获取列表
func (m SlideBanner) GetList(db *gorm.DB, p, pz int) ([]*SlideBanner, error) {
	var sld []*SlideBanner
	var offset = page.GetOffset(p, pz)
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Where("cid=?", m.CID).Offset(offset).Limit(pz).Order("sort desc,id desc").Find(&sld).Error; err != nil {
		return nil, err
	}
	return sld, nil
}

// Count 获取统计
func (m SlideBanner) Count(db *gorm.DB) int64 {
	var count int64
	if err := db.Model(&m).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// BannerISExists 判断分类下是否有幻灯
func (m SlideBanner) BannerISExists(db *gorm.DB) bool {
	var c int64
	if err := db.Model(&m).Where("cid=?", m.CID).Count(&c).Error; err != nil {
		return false
	}

	if c <= 0 {
		return false
	}

	return true
}
