package model

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"reflect"
	"sblog/system/common"
	"strings"
)

var prefix string

func init() {
	prefix = viper.GetString("database.prefix")
}

type LeftMenu struct {
	ID        uint32     `json:"id"`
	PID       uint32     `json:"pid"`
	Path      string     `json:"path"`
	Title     string     `json:"title"`
	Icon      string     `json:"icon"`
	Component string     `json:"component"`
	IsShow    uint8      `json:"isShow"`
	Redirect  string     `json:"redirect"`
	Children  []LeftMenu `json:"children"`
}

// GetGroupMenuIDByUserID 获取用户所拥有的菜单
func GetGroupMenuIDByUserID(db *gorm.DB, uid uint32) ([]string, error) {
	var ruleIDS []string
	//如果是第一个用户直接放行
	if uid == 1 {
		ruleIDS = append(ruleIDS, "super")
		return ruleIDS, nil
	}

	//结果集
	type Result struct {
		ID        uint32 `json:"id"`
		Rules     string `json:"rules"`
		Halfcheck string `json:"halfcheck"`
	}

	var ret []*Result
	if err := db.Select("ag.rules,ag.halfcheck,ag.id").Table(prefix+"auth_group.ag").Joins("left join "+prefix+"auth_group_access as aga on aga.group_id=ag.id").Where("aga.uid=? and ag.status=?", uid, 1).Scan(&ret).Error; err != nil {
		return ruleIDS, err
	}

	//如果不在角色组就返回
	if len(ret) == 0 {
		return ruleIDS, nil
	}

	for _, v := range ret {
		if v.ID == 1 {
			//如果存在第一个角色则是超级管理员直接放行
			ruleIDS = append(ruleIDS, "super")
			return ruleIDS, nil
		} else {
			var rules = strings.Split(v.Rules, ",")
			var halfcheck = strings.Split(v.Halfcheck, ",")
			ruleIDS = append(ruleIDS, rules...)
			ruleIDS = append(ruleIDS, halfcheck...)
		}
	}

	var ruleIds2 []string
	//去重
	for _, v := range ruleIDS {
		if !common.InArray2(v, ruleIds2) {
			ruleIds2 = append(ruleIds2, v)
		}
	}

	return ruleIds2, nil
}

// GetLeftMenuTree 生成树形菜单
func GetLeftMenuTree(node []*AuthRule, pid uint32) []LeftMenu {
	var leftMenuInfo []LeftMenu

	for _, v := range node {
		if v.PID == pid {
			child := GetLeftMenuTree(node, v.ID)
			var lm = LeftMenu{
				ID:        v.ID,
				PID:       v.PID,
				Path:      v.Url,
				Title:     v.Title,
				Icon:      v.Icon,
				Component: v.Component,
				IsShow:    v.IsMenu,
				Redirect:  v.Redirect,
			}
			lm.Children = child
			leftMenuInfo = append(leftMenuInfo, lm)
		}
	}

	return leftMenuInfo
}

// Cascade 级联菜单
type Cascade struct {
	Label    string    `json:"label" xml:"label"`
	Value    uint32    `json:"value" xml:"value"`
	PID      uint32    `json:"pid" xml:"pid"`
	Children []Cascade `json:"children" xml:"children"`
}

// CascadeName 级联菜单反射名称
type CascadeName struct {
	ID    string `json:"id"`
	PID   string `json:"pid"`
	Title string `json:"title"`
}

/*// GetCascadeTree 获取级联菜单树形结构
func GetCascadeTree(node []*AuthRule, pid uint32) []Cascade {
	var cas []Cascade

	for _, v := range node {
		if v.PID == pid {
			child := GetCascadeTree(node, v.ID)
			var ca = Cascade{
				Label:    v.Title,
				Value:    v.ID,
				PID:      v.PID,
				Children: child,
			}
			cas = append(cas, ca)
		}
	}

	return cas
}*/

// GetCascadeTree 获取级联菜单树形结构
func GetCascadeTree(node interface{}, pid uint32, casName CascadeName) []Cascade {
	var cas []Cascade

	//获取值类型
	ty := reflect.TypeOf(node)
	if ty.Kind() != reflect.Slice {
		return cas
	}

	obj := reflect.ValueOf(node)
	//遍历数组
	for i := 0; i < obj.Len(); i++ {
		var val = obj.Index(i)
		//如果不是结构体则返回
		if reflect.TypeOf(val).Kind() != reflect.Struct {
			return cas
		}
		ptr := val.Elem()

		var vpid = ptr.FieldByName(casName.PID).Interface().(uint32)
		var vid = ptr.FieldByName(casName.ID).Interface().(uint32)
		if vpid == pid {

			child := GetCascadeTree(node, vid, casName)
			var ca = Cascade{
				Label:    ptr.FieldByName(casName.Title).String(),
				Value:    vid,
				PID:      vpid,
				Children: child,
			}
			cas = append(cas, ca)
		}
	}

	return cas
}

// TreeMenu 树形菜单
type TreeMenu struct {
	ID       uint32     `json:"id"`
	PID      uint32     `json:"pid"`
	Title    string     `json:"title"`
	Url      string     `json:"url"`
	Target   string     `json:"target"`
	IsChild  bool       `json:"is_child"`
	Children []*TreeMenu `json:"children"`
	Active   bool       `json:"active"`
}

func GetTreeMenu(ms []*Menu, pid uint32) []*TreeMenu {
	var tms []*TreeMenu
	for _, v := range ms {
		if v.PID == pid {
			child := GetTreeMenu(ms, v.ID)
			var tm = &TreeMenu{
				ID:       v.ID,
				PID:      v.PID,
				Title:    v.Title,
				Url:      v.Url,
				Target:   v.Target,
				IsChild:  child != nil,
				Children: child,
			}
			tms = append(tms, tm)
		}
	}

	return tms
}
