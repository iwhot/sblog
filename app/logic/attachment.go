package logic

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"path/filepath"
	"sblog/system/common"
	"time"
)

type UploadFileInfo struct {
	Path     string
	FileName string
	Ext      string
	Size     int64
	Mold     string
	Width    uint32
	Height   uint32
}

// UploadFile 上传文件
func UploadFile(ctx *fiber.Ctx) (*UploadFileInfo, error) {
	var ret = new(UploadFileInfo)

	file, err := ctx.FormFile("file")
	if err != nil {
		return ret, err
	}
	name := file.Filename
	size := file.Size
	ext := filepath.Ext(name) //获取后缀

	//允许列表
	var allowList = []string{
		".jpg", ".jpeg", ".png", ".bmp", ".gif",
		".txt", ".doc", ".docx",
		".xls", ".xlsx", ".csv",
		".mp4", ".m3u8", ".3gp", ".flv", ".avi",
		".mp3",
		".zip", ".tar", ".gz", ".tar.gz",
	}

	if !common.InArray2(ext, allowList) {
		return ret, errors.New("不允许上传的类型")
	}

	paths := "/storage/uploads/" + time.Now().Format("200601") + "/"
	//创建文件夹
	err = common.CreateFile(fmt.Sprintf(".%s", paths))
	if err != nil {
		return ret, err
	}
	//获取文件名称
	newFileName := paths + common.Md5V(time.Now().Format("20060102150405")+common.CreateCaptcha()) + ext
	//保存文件
	err = ctx.SaveFile(file, fmt.Sprintf(".%s", newFileName))
	if err != nil {
		return ret, err
	}

	ret.FileName = name
	ret.Path = newFileName
	ret.Ext = ext
	ret.Size = size
	ret.Mold = common.GetFileType(ext)
	ret.Width = 0
	ret.Height = 0

	return ret, nil
}
