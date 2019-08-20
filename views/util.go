package views

import (
	"github.com/daiguadaidai/haechi/controllers"
	"github.com/daiguadaidai/haechi/middlewares"
	"github.com/daiguadaidai/haechi/views/form"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	handler := new(UtilHandler)
	AddHandlerV1("/utils", handler) // 添加当前页面的uri路径之前都要添加上这个
}

// 注册route
func (this *UtilHandler) RegisterV1(group *gin.RouterGroup) {
	group.GET("/encrypt", this.Encrypt)
	group.GET("/decrypt", this.Decrypt)
	// group.GET("/use_db_sample", this.UseDBSample)
}

// 实现接口, 是否需要进行认证
func (this *UtilHandler) IsAuth() bool { return true }

type UtilHandler struct{}

// 加密
func (this *UtilHandler) Encrypt(c *gin.Context) {
	serverCtx, err := middlewares.GetServerContext(c)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	f, err := form.NewForm(c, &form.UtilEncreptForm{})
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	data, err := controllers.NewUtilController(serverCtx).Encrypt(f.(*form.UtilEncreptForm))
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
	}
	returnSuccess(c, data)
}

// 解密
func (this *UtilHandler) Decrypt(c *gin.Context) {
	serverCtx, err := middlewares.GetServerContext(c)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	f, err := form.NewForm(c, &form.UtilDecryptForm{})
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	data, err := controllers.NewUtilController(serverCtx).Decrypt(f.(*form.UtilDecryptForm))
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
	}
	returnSuccess(c, data)
}

/*
// 使用数据库样例
func (this *UtilHandler) UseDBSample(c *gin.Context) {
	serverCtx, err := middlewares.GetServerContext(c)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	f, err := form.NewForm(c, &form.UtilUseDBSampleForm{})
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	data, err := controllers.NewUtilController(serverCtx).UseDBSample(f.(*form.UtilUseDBSampleForm))
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
	}
	returnSuccess(c, data)
}
*/
