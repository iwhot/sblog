package dao

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"sblog/app/gorm/model"
	"sblog/app/logic"
	"sblog/system/common"
)

type configs struct {
}

var DefaultConfigsDao = configs{}

// GetList 获取配置列表
func (c configs) GetList(ctx *fiber.Ctx) (map[string]map[string]interface{}, error) {
	var cfg = model.Configs{}
	return cfg.GetList(masterDB)
}

// Save 保存配置
func (c configs) Save(ctx *fiber.Ctx) error {
	var p = make(map[string]interface{})
	if err := common.ParseParam(ctx, &p); err != nil {
		return err
	}

	for _, v := range p {
		//值暂时先用string，后面出现复杂的再处理
		var mp = v.(map[string]interface{})
		if mp != nil {
			for key, val := range mp {

				var cfg = model.Configs{
					Name:    key,
					Content: val.(string),
				}
				_ = cfg.Save(masterDB)
			}
		}
	}

	//保存配置到文件
	var pj = new(logic.Setting)
	if err := common.ParseParam(ctx, pj); err != nil {
		return err
	}

	b, _ := json.MarshalIndent(pj, "", "\t")
	_ = ioutil.WriteFile("config/setting.json", b, 0777)
	return nil
}
