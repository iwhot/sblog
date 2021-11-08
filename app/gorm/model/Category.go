package model

import (
	"errors"
	"gorm.io/gorm"
	"sblog/app/logic"
)

type Category struct {
	ID          uint32          `gorm:"column:id;primaryKey" json:"id"`
	PID         uint32          `gorm:"column:pid" json:"pid"`
	Name        string          `gorm:"column:name" json:"name"`
	Key         string          `gorm:"column:key" json:"key"`
	Status      uint8           `gorm:"column:status" json:"status"`
	Sort        uint32          `gorm:"column:sort" json:"sort"`
	Createtime  logic.LocalTime `gorm:"column:createtime" json:"createtime"`
	Updatetime  logic.LocalTime `gorm:"column:updatetime" json:"updatetime"`
	HasChildren bool            `gorm:"column:has_children" json:"hasChildren"`
	///Article     []Article       `gorm:"foreignKey:CID" json:"article"` //一对多，指定外键CID
	Desc  string `gorm:"column:desc" json:"desc"`
	Thumb string `gorm:"column:thumb" json:"thumb"`
}

// Create 创建
func (m Category) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// Update 更新
func (m Category) Update(db *gorm.DB) error {
	return db.Model(&m).Omit("Createtime").Save(m).Error
}

// GetList 获取列表
func (m Category) GetList(db *gorm.DB) ([]*Category, error) {
	var cates []*Category
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Model(&m).Where("pid=?", m.PID).Order("sort desc,id asc").Find(&cates).Error; err != nil {
		return cates, err
	}

	for _, v := range cates {
		v.HasChildren = v.IsChildren(db)
	}

	return cates, nil
}

// IsChildren 检测是否有下级
func (m Category) IsChildren(db *gorm.DB) bool {
	var ct *Category
	if err := db.Model(&m).Where("pid=?", m.ID).First(&ct).Error; err != nil {
		return false
	}

	return true
}

// Delete 删除
func (m Category) Delete(db *gorm.DB) error {
	tx := db.Begin()
	//检测下级
	if m.IsChildren(tx) == true {
		tx.Rollback()
		return errors.New("有下级分类不可删除")
	}
	//检测文章
	var at = Article{CID: m.ID}
	if at.IsArticleByCategory(tx) == true {
		tx.Rollback()
		return errors.New("分类下有文章不可删除")
	}
	//删除
	if err := tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// GetParentList 获取顶级菜单
func (m Category) GetParentList(db *gorm.DB) ([]*Category, error) {
	var cts []*Category
	if err := db.Model(&m).Where("pid=?", 0).Find(&cts).Error; err != nil {
		return cts, err
	}
	return cts, nil
}

// UpdateInfo 保存信息
func (m Category) UpdateInfo(db *gorm.DB, mp map[string]interface{}) error {
	if mp["id"] == nil {
		return errors.New("分类不存在")
	}
	return db.Model(m).Where("id=?", mp["id"]).Updates(mp).Error
}

// CascadeCategory 获取级联分类
func (m Category) CascadeCategory(db *gorm.DB) ([]Cascade, error) {
	var cas []Cascade
	var allCate []*Category
	if err := db.Where("status=?", 1).Find(&allCate).Error; err != nil {
		return cas, err
	}
	//传入的名称
	var cname = CascadeName{
		ID:    "ID",
		PID:   "PID",
		Title: "Name",
	}
	cas = GetCascadeTree(allCate, 0, cname)
	return cas, nil
}

// FindOne 获取一个菜单
func (m Category) FindOne(db *gorm.DB) (*Category, error) {
	var ar *Category
	if err := db.Model(&m).Where("id=?", m.ID).First(&ar).Error; err != nil {
		return ar, err
	}
	return ar, nil
}

// FindOneByPID 获取一个菜单
func (m Category) FindOneByPID(db *gorm.DB) (*Category, error) {
	var ar *Category
	if err := db.Model(&m).Where("id=?", m.PID).First(&ar).Error; err != nil {
		return ar, err
	}
	return ar, nil
}

// GetAllCategoryId 获取所有上级id
func (m Category) GetAllCategoryId(db *gorm.DB, nowId []uint32) []uint32 {
	var ret []uint32
	ret = append(ret, nowId...)

	ar, err := m.FindOneByPID(db)
	if err != nil {
		return ret
	}
	ret = append(ret, ar.ID)
	if ar.PID > 0 {
		return m.GetAllCategoryId(db, ret)
	}

	return ret
}

//todo 分类====================================================================================

// GetCategoryByKey 获取分类
func (m Category) GetCategoryByKey(db *gorm.DB) (*Category, error) {
	var cat *Category
	if err := db.Where("`status` = ? AND `key` = ?", 1, m.Key).First(&cat).Error; err != nil {
		return nil, err
	}

	return cat, nil
}

// GetCrumbsByID 获取所有上级
func (m Category) GetCrumbsByID(db *gorm.DB, nowCategory []*Category) ([]*Category, error) {
	var ret []*Category
	ret = append(ret, nowCategory...)

	ar, err := m.FindOne(db)
	if err != nil {
		return ret, err
	}

	var ret2 []*Category
	ret2 = append(ret2, ar)
	ret2 = append(ret2, ret...)

	if ar.PID > 0 {
		return m.GetCrumbsByID(db, ret2)
	}

	return ret2, nil
}

// GetAllChildList 获取子级栏目
func (m Category) GetAllChildList(db *gorm.DB) ([]*Category, error) {
	var ret []*Category
	if err := db.Model(&m).Where("`pid`=?", m.ID).Order("`sort` desc,`id` desc").Find(&ret).Error; err != nil {
		return ret, err
	}

	return ret, nil
}
