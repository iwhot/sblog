package model

import (
	"gorm.io/gorm"
	"sblog/app/logic"
	"sblog/system/common"
	"sblog/system/page"
	"time"
)

type Admin struct {
	ID           uint32          `gorm:"column:id;primaryKey" json:"id"`
	Username     string          `gorm:"column:username" json:"username"`
	Nickname     string          `gorm:"column:nickname" json:"nickname"`
	Password     string          `gorm:"column:password" json:"password"`
	Salt         string          `gorm:"column:salt" json:"salt"`
	Avatar       string          `gorm:"column:avatar" json:"avatar"`
	Email        string          `gorm:"column:email" json:"email"`
	Loginfailure uint8           `gorm:"column:loginfailure" json:"loginfailure"`
	Logintime    logic.LocalTime `gorm:"column:logintime" json:"logintime"`
	Loginip      string          `gorm:"column:loginip" json:"loginip"`
	Createtime   logic.LocalTime `gorm:"column:createtime" json:"createtime"`
	Updatetime   logic.LocalTime `gorm:"column:updatetime" json:"updatetime"`
	Token        string          `gorm:"column:token" json:"token"`
	Status       uint8           `gorm:"column:status" json:"status"`
	Sort         uint32          `gorm:"column:sort" json:"sort"`
}

// Create 创建
func (m Admin) Create(db *gorm.DB, groupIDS []uint32) error {
	var tx = db.Begin()

	m.Salt = common.CreateSalt(6)
	m.Password = common.GetPassword(m.Password, m.Salt)
	m.Createtime = logic.LocalTime{Time: time.Now()}
	m.Updatetime = logic.LocalTime{Time: time.Now()}
	m.Logintime = logic.LocalTime{Time: time.Now()}

	//创建用户
	if err := tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	//根据用户id删除组关联
	var ags1 = AuthGroupAccess{
		UID: m.ID,
	}
	if err := ags1.DeleteByUID(tx); err != nil {
		tx.Rollback()
		return err
	}

	//创建组
	for _, v := range groupIDS {
		var ags = AuthGroupAccess{
			UID:     m.ID,
			GroupID: v,
		}

		if err := ags.Create(tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

// Update 简单更新
func (m Admin) Update(db *gorm.DB) error {
	return db.Model(&m).Omit("Createtime").Save(m).Error
}

// Update2 更新
func (m Admin) Update2(db *gorm.DB, groupIDS []uint32) error {
	var tx = db.Begin()

	am, err := m.FindOneAdmin(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	m.Username = am.Username

	if m.Password == "" {
		m.Password = am.Password
	} else {
		m.Password = common.GetPassword(m.Password, am.Salt)
	}

	m.Logintime = am.Logintime
	m.Loginip = am.Loginip
	m.Loginfailure = am.Loginfailure
	m.Token = am.Token
	m.Salt = am.Salt

	//忽略更新创建时间
	if err := tx.Model(&m).Omit("Createtime").Save(m).Error; err != nil {
		tx.Rollback()
		return err
	}

	//根据用户id删除组关联
	var ags1 = AuthGroupAccess{
		UID: m.ID,
	}
	if err := ags1.DeleteByUID(tx); err != nil {
		tx.Rollback()
		return err
	}

	//创建组
	for _, v := range groupIDS {
		var ags = AuthGroupAccess{
			UID:     m.ID,
			GroupID: v,
		}

		if err := ags.Create(tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

// Update3 更新3
func (m Admin) Update3(db *gorm.DB, mp map[string]interface{}) error {
	var am,err = m.FindOneAdmin(db)
	if err != nil{
		return err
	}

	var updateInfo = make(map[string]interface{})
	updateInfo["id"] = mp["id"]
	updateInfo["nickname"] = mp["nickname"]
	updateInfo["avatar"] = mp["avatar"]
	updateInfo["email"] = mp["email"]

	if mp["password"] == "" {
		updateInfo["password"] = am.Password
	} else {
		updateInfo["password"] = common.GetPassword(mp["password"].(string), am.Salt)
	}

	return db.Model(&m).Where("id=?", m.ID).Updates(updateInfo).Error
}

// Delete 删除
func (m Admin) Delete(db *gorm.DB) error {
	var tx = db.Begin()

	if err := tx.Where("id = ?", m.ID).Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	//根据用户id删除组关联
	var ags1 = AuthGroupAccess{
		UID: m.ID,
	}
	if err := ags1.DeleteByUID(tx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// DeleteAll 批量删除
func (a Admin) DeleteAll(db *gorm.DB, ids []string) error {
	var tx = db.Begin()

	if err := db.Where("id in ?", ids).Delete(&a).Error; err != nil {
		tx.Rollback()
		return err
	}

	//根据用户id删除组关联
	var ags1 = AuthGroupAccess{}
	if err := ags1.DeleteByUIDs(tx, ids); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// FindOneAdmin 获取一个会员
func (a Admin) FindOneAdmin(db *gorm.DB) (*Admin, error) {
	var admin *Admin
	tx := db.Model(&a)

	if a.ID != 0 {
		tx = tx.Where("id=?", a.ID)
	} else if a.Username != "" {
		tx = tx.Where("username=?", a.Username)
	}

	if err := tx.First(&admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

// Login 登录
func (a Admin) Login(db *gorm.DB) (*Admin, error) {
	var adminInfo, err = a.FindOneAdmin(db)
	if err != nil {
		return nil, err
	}

	//校验密码

	return adminInfo, err
}

// GetAdminList 获取管理员列表
func (a Admin) GetAdminList(db *gorm.DB, p, pz int) ([]*Admin, error) {
	var admins []*Admin
	var offset = page.GetOffset(p, pz)
	var tx = db.Session(&gorm.Session{PrepareStmt: true})

	if err := tx.Offset(offset).Limit(pz).Order("sort desc,id desc").Find(&admins).Error; err != nil {
		return nil, err
	}

	return admins, nil
}

// Count 统计管理员数量
func (a Admin) Count(db *gorm.DB) int64 {
	var count int64
	if err := db.Model(&a).Count(&count).Error; err != nil {
		return 0
	}

	return count
}

// GetAdminInfoByToken 通过token定位用户
func (m Admin) GetAdminInfoByToken(db *gorm.DB) (*Admin, error) {
	var ut = UserToken{
		Token: m.Token,
	}

	tu2, err := ut.FindOneToken(db)
	if err != nil {
		return nil, err
	}

	var am = Admin{
		ID: tu2.UserID,
	}

	//取出用户
	am2, err := am.FindOneAdmin(db)
	if err != nil {
		return nil, err
	}

	return am2, nil
}
