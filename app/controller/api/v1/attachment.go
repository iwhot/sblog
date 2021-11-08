package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sblog/app/controller"
	"sblog/app/gorm/dao"
)

// Attachment 附件
type Attachment struct {
	controller.Base
}

// NewRouter 本页面相关路由
func (c *Attachment) NewRouter(g fiber.Router) {
	g.Post("/attachment/index/:mold", c.Index)
	g.Post("/attachment/upload", c.Upload)
	g.Post("/attachment/upload-editor", c.EditorUpload)
	g.Post("/attachment/delete", c.Delete)
	g.Post("/attachment/download/:id", c.Download)
	g.Post("/attachment/edit", c.Edit)
}

// Index 附件列表
func (c *Attachment) Index(ctx *fiber.Ctx) error {
	aus, rp, _ := dao.DefaultAttachmentDao.Index(ctx)
	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "请求成功",
		"data": map[string]interface{}{
			"list": aus,
			"page": rp,
		},
	})
}

// Upload 附件上传
func (c *Attachment) Upload(ctx *fiber.Ctx) error {
	admin, _ := dao.DefaultAdminDao.GetAdminInfoByToken(ctx)
	paths, err := dao.DefaultAttachmentDao.Upload(ctx, admin.ID)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  "上传失败",
		})
	}

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "上传成功",
		"data": paths,
	})
}

// EditorUpload 专门为富文本上传封装一个
func (c *Attachment) EditorUpload(ctx *fiber.Ctx) error {
	admin, _ := dao.DefaultAdminDao.GetAdminInfoByToken(ctx)
	paths, err := dao.DefaultAttachmentDao.Upload(ctx, admin.ID)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": 0,
			"msg":  "上传失败",
			"path": "",
		})
	}

	return ctx.JSON(fiber.Map{
		"code": 1,
		"msg":  "上传成功",
		"path": paths,
	})
}

// Delete 附件删除
func (c *Attachment) Delete(ctx *fiber.Ctx) error {
	err := dao.DefaultAttachmentDao.Delete(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  "删除失败",
		})
	}

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "删除成功",
	})
}

// Download 附件下载
func (c *Attachment) Download(ctx *fiber.Ctx) error {
	path, name := dao.DefaultAttachmentDao.Download(ctx)
	return ctx.Download(fmt.Sprintf(".%s", path), name)
}

// Edit 编辑,先只是简单编辑信息,后面加入图片剪裁
func (c *Attachment) Edit(ctx *fiber.Ctx) error {
	err := dao.DefaultAttachmentDao.Edit(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code": "500",
			"msg":  "编辑失败",
		})
	}

	return ctx.JSON(fiber.Map{
		"code": "200",
		"msg":  "编辑成功",
	})
}

// Tailoring 图片剪裁
func (c *Attachment) Tailoring(ctx *fiber.Ctx) error {
	//todo 乱七八糟
	return nil
}
