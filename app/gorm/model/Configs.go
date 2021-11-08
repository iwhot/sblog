package model

import (
	"gorm.io/gorm"
)

type Configs struct {
	ID      uint32 `gorm:"column:id;primaryKey" json:"id"`
	Title   string `gorm:"column:title" json:"title"`
	Group   string `gorm:"column:group" json:"group"`
	Name    string `gorm:"column:name" json:"name"`
	Content string `gorm:"column:content" json:"content"`
	Type    string `gorm:"column:type" json:"type"`
	Option  string `gorm:"column:option" json:"option"`
}

// GetList 获取列表
func (c Configs) GetList(db *gorm.DB) (map[string]map[string]interface{}, error) {
	var ret = make(map[string]map[string]interface{})
	var allConfigs []*Configs
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Find(&allConfigs).Error; err != nil {
		return ret, err
	}

	for _, v := range allConfigs {
		ret[v.Group] = make(map[string]interface{})
	}

	for _, v := range allConfigs {
		ret[v.Group][v.Name] = v.Content
	}

	return ret, nil
}

// Save 保存一个配置
func (c Configs) Save(db *gorm.DB) error {
	return db.Model(c).Where("name=?", c.Name).Updates(c).Error
}

// GetOneConfigContent 获取一个配置内容
func (c Configs) GetOneConfigContent(db *gorm.DB) string {
	var cfg *Configs
	if err := db.Where("name=?", c.Name).First(&cfg).Error; err != nil {
		return ""
	}
	return cfg.Content
}
