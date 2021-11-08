package model

import (
	"errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"sblog/system/common"
	"strconv"
	"time"
)

type UserToken struct {
	Token      string `gorm:"column:token;primaryKey" json:"token"`
	UserID     uint32 `gorm:"column:user_id" json:"user_id"`
	Createtime int    `gorm:"column:createtime" json:"createtime"`
	Expiretime int    `gorm:"column:expiretime" json:"expiretime"`
}

// Create 创建一个Token
func (m UserToken) Create(db *gorm.DB) (string, error) {
	//开启事物
	var tx = db.Begin()
	//删除其他相关token，只允许一个地方登陆
	if err := tx.Where("user_id=?", m.UserID).Delete(&m).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	var ctime = int(time.Now().Unix())
	m.Createtime = ctime
	m.Expiretime = ctime + viper.GetInt("server.tokenExpiretime")
	//弄一个唯一的数字作为token
	var str = strconv.Itoa(int(m.UserID)) + strconv.Itoa(int(time.Now().UnixNano()))
	m.Token = common.Md5V(str)

	//创建token
	if err := tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()
	return m.Token, nil
}

// FindOneToken 获取一个Token
func (m UserToken) FindOneToken(db *gorm.DB) (*UserToken, error) {
	var token *UserToken
	var tx = db.Session(&gorm.Session{PrepareStmt: true})
	if err := tx.Model(&m).Where("token=?", m.Token).First(&token).Error; err != nil {
		return nil, err
	}

	return token, nil
}

// CheckToken 校验token
func (m UserToken) CheckToken(db *gorm.DB) error {
	u, e := m.FindOneToken(db)
	if e != nil || u == nil {
		return e
	}

	//校验过期
	var now = int(time.Now().Unix())
	//如果当前时间比过期时间大则token失效
	if now > u.Expiretime {
		return errors.New("token已过期")
	}

	return nil
}
