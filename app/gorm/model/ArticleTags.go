package model

import "gorm.io/gorm"

type ArticleTags struct {
	ArticleID uint32 `gorm:"column:article_id" json:"article_id"`
	TagsID    uint32 `gorm:"column:tags_id" json:"tags_id"`
}

// DeleteByTagsID 删除所有标签id
func (m ArticleTags) DeleteByTagsID(db *gorm.DB) error {
	return db.Where("tags_id = ?", m.TagsID).Delete(m).Error
}

// DeleteByArticleID 删除文章id
func (m ArticleTags) DeleteByArticleID(db *gorm.DB) error {
	return db.Where("article_id = ?", m.ArticleID).Delete(m).Error
}
