package model

import (
	"gorm.io/gorm"
	"sblog/app/logic"
	"sblog/system/page"
)

type AuthGroup struct {
	ID         uint32          `gorm:"column:id;primaryKey" json:"id"`
	PID        uint32          `gorm:"column:pid" json:"pid"`
	Name       string          `gorm:"column:name" json:"name"`
	Rules      string          `gorm:"column:rules" json:"rules"`
	Createtime logic.LocalTime `gorm:"column:createtime" json:"createtime"`
	Updatetime logic.LocalTime `gorm:"column:updatetime" json:"updatetime"`
	Status     uint8           `gorm:"column:status" json:"status"`
	Sort       uint32          `gorm:"column:sort" json:"sort"`
	Halfcheck  string          `gorm:"column:halfcheck" json:"halfcheck"`
}

// Create 创建
func (m AuthGroup) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// Update 更新
func (m AuthGroup) Update(db *gorm.DB) error {
	return db.Model(&m).Omit("Createtime").Save(m).Error
}

// Delete 删除
func (m AuthGroup) Delete(db *gorm.DB) error {
	return db.Where("id = ?", m.ID).Delete(&m).Error
}

// DeleteAll 批量删除
func (m AuthGroup) DeleteAll(db *gorm.DB, ids []string) error {
	return db.Where("id in ?", ids).Delete(&m).Error
}

// GetList 获取列表
func (m AuthGroup) GetList(db *gorm.DB, p, pz int) ([]*AuthGroup, error) {
	var ags []*AuthGroup
	var offset = page.GetOffset(p, pz)
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Offset(offset).Limit(pz).Order("sort desc,id desc").Find(&ags).Error; err != nil {
		return nil, err
	}

	return ags, nil
}

// Count 统计
func (m AuthGroup) Count(db *gorm.DB) int64 {
	var count int64
	if err := db.Model(&m).Count(&count).Error; err != nil {
		return 0
	}

	return count
}

// GetAllList 获取所有的组
func (m AuthGroup) GetAllList(db *gorm.DB) ([]map[string]interface{}, error) {
	var ret []map[string]interface{}
	var ags []*AuthGroup
	if err := db.Model(&m).Where("status=?", 1).Order("sort desc,id desc").Find(&ags).Error; err != nil {
		return ret, err
	}

	for _, v := range ags {
		var val = make(map[string]interface{})
		val["label"] = v.Name
		val["value"] = v.ID
		ret = append(ret, val)
	}

	return ret, nil
}
