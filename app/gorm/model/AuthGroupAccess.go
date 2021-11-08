package model

import "gorm.io/gorm"

// AuthGroupAccess 此处相关操作懒得用关联模式
type AuthGroupAccess struct {
	UID     uint32 `gorm:"column:uid" json:"uid"`
	GroupID uint32 `gorm:"column:group_id" json:"group_id"`
}

// Create 创建
func (m AuthGroupAccess) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// Update 更新
func (m AuthGroupAccess) Update(db *gorm.DB) error {
	return db.Model(&m).Save(m).Error
}

// DeleteByUID 根据用户id删除
func (m AuthGroupAccess) DeleteByUID(db *gorm.DB) error {
	return db.Where("uid=?", m.UID).Delete(&m).Error
}

// DeleteByUIDs 根据用户id删除
func (m AuthGroupAccess) DeleteByUIDs(db *gorm.DB, uids []string) error {
	return db.Where("uid IN ?", uids).Delete(&m).Error
}

// DeleteByGroupID 根据group_id删除
func (m AuthGroupAccess) DeleteByGroupID(db *gorm.DB) error {
	return db.Where("group_id=?", m.GroupID).Delete(&m).Error
}

// GetGroupIdByUserId 根据用户id获取所有组id
func (m AuthGroupAccess) GetGroupIdByUserId(db *gorm.DB) ([]*AuthGroupAccess, error) {
	var aga []*AuthGroupAccess
	if err := db.Model(&m).Where("uid=?", m.UID).Find(&aga).Error; err != nil {
		return aga, err
	}

	return aga, nil
}
