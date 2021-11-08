package model

import (
	"gorm.io/gorm"
	"sblog/app/logic"
)

type ArticleLike struct {
	ArticleID  uint32          `gorm:"column:article_id" json:"article_id"`
	Ip         string          `gorm:"column:ip" json:"ip"`
	Createtime logic.LocalTime `gorm:"column:createtime" json:"createtime"`
}

// FindOne 获取
func (m ArticleLike) FindOne(db *gorm.DB) bool {
	var al *ArticleLike
	if err := db.Model(&m).Where("`article_id`=? AND `ip`=?", m.ArticleID, m.Ip).First(&al).Error; err != nil {
		return false
	}

	if al.ArticleID == 0 {
		return false
	}

	return true
}

// Create 创建
func (m ArticleLike) Create(db *gorm.DB) error {
	var tx = db.Begin()

	if m.FindOne(tx){
		tx.Rollback()
		return nil
	}

	if err := tx.Model(&Article{ID: m.ArticleID}).Update("like",gorm.Expr("`like`+?",1)).Error;err != nil{
		tx.Rollback()
		return err
	}

	if err := tx.Create(&m).Error;err != nil{
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
