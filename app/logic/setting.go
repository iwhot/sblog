package logic

import (
	"encoding/json"
	"io/ioutil"
)

// Basic 基础配置
type Basic struct {
	SiteName   string `json:"site_name"`
	Bei        string `json:"bei"`
	Footer     string `json:"footer"`
	Wechat     string `json:"wechat"`
	QQ         string `json:"qq"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Address    string `json:"address"`
	Occupation string `json:"occupation"`
	SiteUrl    string `json:"site_url"`
	EmailMe    string `json:"email_me"`
	Reprint    string `json:"reprint"`
}

// SEO seo配置
type SEO struct {
	SeoTitle string `json:"seo_title"`
	SeoKey   string `json:"seo_key"`
	SeoDesc  string `json:"seo_desc"`
}

type Setting struct {
	Basic Basic `json:"basic"`
	SEO   SEO   `json:"seo"`
}

// GetAllSetting 获取所有配置
func GetAllSetting() Setting {
	//初始化
	var sts = Setting{
		Basic: Basic{
			SiteName:   "sblog",
			Bei:        "",
			Footer:     "",
			Wechat:     "",
			QQ:         "",
			Email:      "",
			Username:   "",
			Address:    "",
			Occupation: "",
			SiteUrl:    "",
			EmailMe:    "",
			Reprint:    "",
		},
		SEO: SEO{
			SeoTitle: "",
			SeoKey:   "",
			SeoDesc:  "",
		},
	}
	b, err := ioutil.ReadFile("config/setting.json")
	if err != nil {
		return sts
	}
	err = json.Unmarshal(b, &sts)
	if err != nil {
		return sts
	}

	return sts
}
