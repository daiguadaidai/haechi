package controllers

import (
	"github.com/daiguadaidai/haechi/contexts"
	"github.com/daiguadaidai/haechi/views/form"
	"github.com/daiguadaidai/peep"
)

type UtilController struct {
	ctx *contexts.HttpContext
}

func NewUtilController(ctx *contexts.HttpContext) *UtilController {
	return &UtilController{
		ctx: ctx,
	}
}

// 加密
func (this *UtilController) Encrypt(f *form.UtilEncreptForm) (string, error) {
	return peep.Encrypt(f.Data)
}

// 解密
func (this *UtilController) Decrypt(f *form.UtilDecryptForm) (string, error) {
	return peep.Decrypt(f.Data)
}

// 使用数据库链接
/*
func (this *UtilController) UseDBSample(f *form.UtilUseDBSampleForm) (*models.Util, error) {
	// 数据库
	// this.ctx.DB

	utilDao := dao.NewUtilDao(this.ctx.DB)
	return utilDao.GetOne(f.Id)
}
*/
