package model

import (
	"errors"
	"gorm.io/gorm"
	"sblog/app/logic"
)

type Menu struct {
	ID          uint32          `gorm:"column:id;primaryKey" json:"id"`
	PID         uint32          `gorm:"column:pid" json:"pid"`
	Title       string          `gorm:"column:title" json:"title"`
	Url         string          `gorm:"column:url" json:"url"`
	Target      string          `gorm:"column:target" json:"target"`
	Status      uint8           `gorm:"column:status" json:"status"`
	Sort        uint32          `gorm:"column:sort" json:"sort"`
	Createtime  logic.LocalTime `gorm:"column:createtime" json:"createtime"`
	Updatetime  logic.LocalTime `gorm:"column:updatetime" json:"updatetime"`
	HasChildren bool            `gorm:"column:has_children" json:"hasChildren"`
	Group       string          `gorm:"column:group" json:"group"`
}

// Create 创建
func (m Menu) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// Updates 更新
func (m Menu) Updates(db *gorm.DB) error {
	return db.Model(&m).Omit("Createtime").Save(m).Error
}

// Delete 删除
func (m Menu) Delete(db *gorm.DB) error {
	tx := db.Begin()
	//检测下级
	if m.IsChildren(tx) == true {
		tx.Rollback()
		return errors.New("有下级分类不可删除")
	}

	if err := db.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// Count 统计
func (m Menu) Count(db *gorm.DB) int64 {
	var count int64

	if err := db.Model(&m).Count(&count).Error; err != nil {
		return 0
	}

	return count
}

// GetList 后台获取列表
func (m Menu) GetList(db *gorm.DB) ([]*Menu, error) {
	var ms []*Menu

	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Model(&m).Where("pid=?", m.PID).Order("sort desc,id desc").Find(&ms).Error; err != nil {
		return nil, err
	}

	for _, v := range ms {
		v.HasChildren = v.IsChildren(db)
	}

	return ms, nil
}

// IsChildren 检测是否有下级
func (m Menu) IsChildren(db *gorm.DB) bool {
	var ct *Menu
	if err := db.Model(&m).Where("pid=?", m.ID).First(&ct).Error; err != nil {
		return false
	}

	return true
}

// GetParentList 获取顶级菜单
func (m Menu) GetParentList(db *gorm.DB) ([]*Menu, error) {
	var cts []*Menu
	if err := db.Model(&m).Where("pid=?", 0).Find(&cts).Error; err != nil {
		return cts, err
	}
	return cts, nil
}

// CascadeCategory 获取级联分类
func (m Menu) CascadeCategory(db *gorm.DB) ([]Cascade, error) {
	var cas []Cascade
	var allCate []*Menu
	if err := db.Where("status=?", 1).Find(&allCate).Error; err != nil {
		return cas, err
	}
	//传入的名称
	var cname = CascadeName{
		ID:    "ID",
		PID:   "PID",
		Title: "Title",
	}
	cas = GetCascadeTree(allCate, 0, cname)
	return cas, nil
}

// FindOne 获取一个菜单
func (m Menu) FindOne(db *gorm.DB) (*Menu, error) {
	var ar *Menu
	if err := db.Model(&m).Where("id=?", m.ID).First(&ar).Error; err != nil {
		return ar, err
	}
	return ar, nil
}

// GetAllMenuId 获取所有上级id
func (m Menu) GetAllMenuId(db *gorm.DB, nowId []uint32) []uint32 {
	var ret []uint32
	ret = append(ret, nowId...)

	ar, err := m.FindOne(db)
	if err != nil {
		return ret
	}
	ret = append(ret, ar.ID)
	if ar.PID > 0 {
		return m.GetAllMenuId(db, ret)
	}

	return ret
}

// UpdateInfo 保存信息
func (m Menu) UpdateInfo(db *gorm.DB, mp map[string]interface{}) error {
	if mp["id"] == nil {
		return errors.New("分类不存在")
	}
	return db.Model(m).Where("id=?", mp["id"]).Updates(mp).Error
}

//todo 前台部分==============================================================================

// GetMenu 获取菜单
func (m Menu) GetMenu(db *gorm.DB, active string) []*TreeMenu {
	var cas []*TreeMenu
	var ms []*Menu
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Model(&m).Where("`group`=? AND `status` = ?", m.Group, 1).Order("sort desc,id desc").Find(&ms).Error; err != nil {
		return cas
	}
	cas = GetTreeMenu(ms, 0)
	for _, v := range cas {
		if v.Url == active {
			v.Active = true
		} else {
			v.Active = false
		}
	}
	return cas
}
