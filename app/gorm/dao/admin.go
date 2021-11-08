package dao

import (
	"errors"
	"github.com/dchest/captcha"
	"github.com/gofiber/fiber/v2"
	"log"
	"math"
	"sblog/app/gorm/model"
	"sblog/system/common"
	"strings"
)

type admin struct {
}

//解析参数 todo application/json用FormValue拿不到，只能解构后再用
type loginParam struct {
	Username string `json:"username" xml:"username" form:"name"`
	Password string `json:"password" xml:"password" form:"password"`
	Captcha  string `json:"captcha" xml:"captcha" form:"captcha"`
	Sess     string `json:"sess" xml:"sess" form:"sess"`
}

var DefaultAdminDao = admin{}

// Login 登录
func (a admin) Login(ctx *fiber.Ctx) (*model.Admin, string, error) {
	p := new(loginParam)
	if err := common.ParseParam(ctx, p); err != nil {
		return nil, "", err
	}
	//校验验证码
	if captcha.VerifyString(p.Sess, p.Captcha) == false {
		return nil, "", errors.New("验证码输入错误")
	}
	var au = model.Admin{
		Username: p.Username,
	}
	//通过用户名取出信息
	var adminInfo, err = au.Login(masterDB)
	if err != nil {
		return nil, "", errors.New("用户名或者密码错误")
	}
	//获取密码
	pwd := common.GetPassword(p.Password, adminInfo.Salt)

	log.Println(pwd,adminInfo.Password)
	//校验密码
	if pwd != adminInfo.Password {
		return nil, "", errors.New("用户名或者密码错误")
	}

	//创建Token
	var ut = model.UserToken{
		UserID: adminInfo.ID,
	}

	token, err := ut.Create(masterDB)
	if err != nil {
		return nil, "", errors.New("用户名或者密码错误")
	}

	adminInfo.Token = token
	err = adminInfo.Update(masterDB)
	if err != nil {
		return nil, "", err
	}

	return adminInfo, token, nil
}

// CheckToken 校验token
func (a admin) CheckToken(token string) error {
	var ut = model.UserToken{
		Token: token,
	}

	return ut.CheckToken(masterDB)
}

// GetAdminList 获取列表
func (a admin) GetAdminList(ctx *fiber.Ctx) ([]*model.Admin, map[string]interface{}, error) {
	var admins []*model.Admin
	var err error
	var rp = make(map[string]interface{})

	p := new(listParam)
	if err = common.ParseParam(ctx, p); err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return admins, rp, err
	}

	var adminUser = model.Admin{}

	admins, err = adminUser.GetAdminList(masterDB, p.Page, p.PageSize)
	if err != nil {
		rp["count"] = 0
		rp["all"] = 0
		rp["page"] = p.Page
		rp["page_size"] = p.PageSize
		return admins, rp, err
	}

	count := adminUser.Count(masterDB)
	allPage := int(math.Ceil(float64(count) / float64(p.PageSize)))
	rp["count"] = count
	rp["all"] = allPage
	rp["page"] = p.Page
	rp["page_size"] = p.PageSize

	return admins, rp, nil
}

// GetAdminInfoByToken 通过token获取当前用户
func (a admin) GetAdminInfoByToken(ctx *fiber.Ctx) (*model.Admin, error) {
	var token = ctx.Get("Authorization")
	var am = model.Admin{
		Token: token,
	}

	return am.GetAdminInfoByToken(masterDB)
}

// Save 因为懒得用关联模式就随便写了个保存
func (d admin) Save(ctx *fiber.Ctx) error {
	type adminParam struct {
		ID       uint32   `json:"id" xml:"id" form:"id"`
		Username string   `json:"username" xml:"username" form:"username"`
		Nickname string   `json:"nickname" xml:"nickname" form:"nickname"`
		Password string   `json:"password" xml:"password" form:"password"`
		Avatar   string   `json:"avatar" xml:"avatar" form:"avatar"`
		Email    string   `json:"email" xml:"email" form:"email"`
		Status   uint8    `json:"status" xml:"status" form:"status"`
		Group    []uint32 `json:"group" xml:"group" form:"group"`
	}

	p := new(adminParam)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}
	var admin1 = model.Admin{
		Username: p.Username,
		Nickname: p.Nickname,
		Password: p.Password,
		Avatar:   p.Avatar,
		Email:    p.Email,
		Status:   p.Status,
		Sort:     1000,
	}
	if p.ID > 0 {
		admin1.ID = p.ID
		if p.Username == "" {
			return errors.New("用户名不得为空")
		}

		//编辑
		return admin1.Update2(masterDB, p.Group)
	} else {
		//创建
		if p.Username == "" {
			return errors.New("用户名不得为空")
		}

		if p.Password == "" {
			return errors.New("密码不得为空")
		}

		return admin1.Create(masterDB, p.Group)
	}
}

// Save2 保存信息
func (d admin) Save2(ctx *fiber.Ctx) error {
	type adminParam struct {
		ID       uint32   `json:"id" xml:"id" form:"id"`
		Nickname string   `json:"nickname" xml:"nickname" form:"nickname"`
		Password string   `json:"password" xml:"password" form:"password"`
		Avatar   string   `json:"avatar" xml:"avatar" form:"avatar"`
		Email    string   `json:"email" xml:"email" form:"email"`
	}

	p := new(adminParam)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var updateInfo = make(map[string]interface{})
	updateInfo["id"] = p.ID
	updateInfo["nickname"] = p.Nickname
	updateInfo["password"] = p.Password
	updateInfo["avatar"] = p.Avatar
	updateInfo["email"] = p.Email
	var admin1 = model.Admin{ID: p.ID}
	return admin1.Update3(masterDB, updateInfo)
}



// Delete 删除
func (d admin) Delete(ctx *fiber.Ctx) error {
	type param struct {
		ID uint32 `json:"id" xml:"id" form:"id"`
	}

	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var admin1 = model.Admin{
		ID: p.ID,
	}

	return admin1.Delete(masterDB)
}

// DeleteAll Delete 批量删除
func (d admin) DeleteAll(ctx *fiber.Ctx) error {
	type param struct {
		ID string `json:"id" xml:"id" form:"id"`
	}

	p := new(param)
	if err := common.ParseParam(ctx, p); err != nil {
		return err
	}

	var admin1 = model.Admin{}
	ids := strings.Split(p.ID, ",")

	return admin1.DeleteAll(masterDB, ids)
}

// GetMyGroupID 获取自己的组id
func (d admin) GetMyGroupID(ctx *fiber.Ctx) ([]uint32, error) {
	var GroupIDs []uint32

	admin, err := d.GetAdminInfoByToken(ctx)
	if err != nil {
		return GroupIDs, err
	}

	var aga = model.AuthGroupAccess{
		UID: admin.ID,
	}

	agas, err := aga.GetGroupIdByUserId(masterDB)
	if err != nil {
		return GroupIDs, err
	}

	for _, v := range agas {
		GroupIDs = append(GroupIDs, v.GroupID)
	}

	return GroupIDs, nil
}
