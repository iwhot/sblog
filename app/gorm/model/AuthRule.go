package model

import (
	"errors"
	"gorm.io/gorm"
	"sblog/app/logic"
	"sblog/system/common"
)

type AuthRule struct {
	ID          uint32          `gorm:"column:id;primaryKey" json:"id"`
	PID         uint32          `gorm:"column:pid" json:"pid"`
	Url         string          `gorm:"column:url" json:"url"`
	Title       string          `gorm:"column:title" json:"title"`
	Icon        string          `gorm:"column:icon" json:"icon"`
	Component   string          `gorm:"column:component" json:"component"`
	Status      uint8           `gorm:"column:status" json:"status"`
	Remark      string          `gorm:"column:remark" json:"remark"`
	Createtime  logic.LocalTime `gorm:"column:createtime" json:"createtime"`
	Updatetime  logic.LocalTime `gorm:"column:updatetime" json:"updatetime"`
	Redirect    string          `gorm:"column:redirect" json:"redirect"`
	Sort        uint32          `gorm:"column:sort" json:"sort"`
	IsMenu      uint8           `gorm:"column:is_menu" json:"is_menu"`
	IsNotDel    uint8           `gorm:"column:is_not_del" json:"is_not_del"`
	HasChildren bool            `gorm:"column:has_children" json:"hasChildren"`
}

// GetMyMenu 获取自己的菜单
func (m AuthRule) GetMyMenu(db *gorm.DB, uid uint32) ([]*AuthRule, error) {
	var ars []*AuthRule
	//取出菜单id
	menuID, err := GetGroupMenuIDByUserID(db, uid)
	if err != nil || len(menuID) <= 0 {
		return ars, nil
	}

	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	tx = tx.Where("status=?", 1)
	if !common.InArray2("super", menuID) {
		//取出所有
		tx = tx.Where("id IN ?", menuID)
	}

	if err = tx.Find(&ars).Error; err != nil {
		return nil, err
	}

	return ars, nil
}

// GetList 获取菜单列表,菜单不分页
func (m AuthRule) GetList(db *gorm.DB) ([]*AuthRule, error) {
	var ars []*AuthRule
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Model(&m).Where("pid=?", m.PID).Order("sort desc,id asc").Find(&ars).Error; err != nil {
		return ars, err
	}

	for _, v := range ars {
		v.HasChildren = v.FineOneByPID(db)
	}

	return ars, nil
}

// Create 创建
func (m AuthRule) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// Update 更新
func (m AuthRule) Update(db *gorm.DB) error {
	return db.Model(&m).Omit("Createtime").Save(m).Error
}

// FindOne 获取一个菜单
func (m AuthRule) FindOne(db *gorm.DB) (*AuthRule, error) {
	var ar *AuthRule
	if err := db.Model(&m).Where("id=?", m.ID).First(&ar).Error; err != nil {
		return ar, err
	}
	return ar, nil
}

// GetAllParentId 递归获取上级菜单
func (m AuthRule) GetAllParentId(db *gorm.DB,nowID []uint32) ([]uint32,error) {
	var ret []uint32
	ret = append(ret,nowID...)
	ar ,err := m.FindOne(db)
	if err != nil{
		return ret, err
	}
	ret = append(ret,ar.ID)
	if ar.PID != 0{
		return m.GetAllParentId(db,ret)
	}

	return ret, err
}

// FineOneByPID 通过pid获取是否有菜单
func (m AuthRule) FineOneByPID(db *gorm.DB) bool {
	var ar *AuthRule
	if err := db.Model(&m).Where("pid=?", m.ID).First(&ar).Error; err != nil {
		return false
	}
	return true
}

// Delete 删除
func (m AuthRule) Delete(db *gorm.DB) error {
	var tx = db.Begin()
	ar, err := m.FindOne(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if ar.IsNotDel == 1 {
		tx.Rollback()
		return errors.New("此菜单是基础菜单为不可删除类型")
	}

	//删除子菜单
	if err = tx.Where("pid=?", ar.ID).Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	//删除本菜单
	if err = tx.Where("id=?", ar.ID).Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// DeleteAll 批量删除
func (m AuthRule) DeleteAll(db *gorm.DB, ids []string) error {
	var tx = db.Begin()
	//删除子菜单
	if err := tx.Where("pid IN ?", ids).Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	//删除本菜单
	if err := tx.Where("id IN ?", ids).Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// UpdatePartial 根据id更新部分字段
func (m AuthRule) UpdatePartial(db *gorm.DB, dt map[string]interface{}) error {
	return db.Model(m).Where("id=?", m.ID).Updates(dt).Error
}

// CascadeMenu 获取级联菜单
func (m AuthRule) CascadeMenu(db *gorm.DB) ([]Cascade, error) {
	var cas []Cascade
	//获取所有菜单
	var allMenu []*AuthRule
	if err := db.Find(&allMenu).Error; err != nil {
		return cas, err
	}
	//传入的名称
	var cname = CascadeName{
		ID:    "ID",
		PID:   "PID",
		Title: "Title",
	}
	cas = GetCascadeTree(allMenu, 0, cname)
	return cas, nil
}
