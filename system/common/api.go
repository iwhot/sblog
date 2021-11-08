package common

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// ParseParam 解析参数
func ParseParam(ctx *fiber.Ctx, out interface{}) error {
	//如果是文本则作解密
	if ctx.Get("Content-Type") == "text/plain" {
		var key = []byte(viper.GetString("encryption.key"))
		var iv = []byte(Md5V(ctx.Get("Timestamp"))[0:16]) //取16位
		//解密参数
		ret, err := Decrypt(string(ctx.Body()), key, iv)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(ret), &out)
		return err
	}

	//正常走
	return ctx.BodyParser(out)
}
