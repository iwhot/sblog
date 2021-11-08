package model

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"sblog/app/logic"
	"sblog/system/common"
	"sblog/system/page"
)

type Attachment struct {
	ID         uint32          `gorm:"column:id;primaryKey" json:"id"`
	UserID     uint32          `gorm:"column:user_id" json:"user_id"`
	Url        string          `gorm:"column:url" json:"url"`
	Filesize   uint32          `gorm:"column:filesize" json:"filesize"`
	Width      uint32          `gorm:"column:width" json:"width"`
	Height     uint32          `gorm:"column:height" json:"height"`
	Createtime logic.LocalTime `gorm:"column:createtime" json:"createtime"`
	Desc       string          `gorm:"column:desc" json:"desc"`
	Mold       string          `gorm:"column:mold" json:"mold"`
	Updatetime logic.LocalTime `gorm:"column:updatetime" json:"updatetime"`
}

// Create 上传
func (m Attachment) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// GetList 列表
func (m Attachment) GetList(db *gorm.DB, p, pz int) ([]*Attachment, error) {
	var atms []*Attachment
	var offset = page.GetOffset(p, pz)
	var tx = db.Session(&gorm.Session{PrepareStmt: true})

	if m.Mold != "" && m.Mold != "all" {
		tx = tx.Where("mold=?", m.Mold)
	}

	if err := tx.Offset(offset).Limit(pz).Order("id desc").Find(&atms).Error; err != nil {
		return nil, err
	}

	return atms, nil
}

// FindOne 获取某个文件
func (m Attachment) FindOne(db *gorm.DB) (*Attachment, error) {
	var at *Attachment
	if err := db.Where("id=?", m.ID).First(&at).Error; err != nil {
		return at, err
	}

	return at, nil
}

// Delete 删除
func (m Attachment) Delete(db *gorm.DB) error {
	var at, err = m.FindOne(db)
	if err != nil {
		return err
	}

	//删除对应数据
	if err := db.Where("id=?", m.ID).Delete(&m).Error; err != nil {
		return err
	}

	//删除对应文件
	if common.IsExists(fmt.Sprintf(".%s", at.Url)) {
		_ = os.Remove(fmt.Sprintf(".%s", at.Url))
	}

	return nil
}

// Count 统计
func (m Attachment) Count(db *gorm.DB) int64 {
	var count int64
	if err := db.Model(&m).Count(&count).Error; err != nil {
		return 0
	}

	return count
}

// Edit 编辑图片
func (m Attachment) Edit(db *gorm.DB) error {
	return db.Model(&m).Select("url,filesize,width,height,desc,mold,updatetime").Updates(m).Error
}
