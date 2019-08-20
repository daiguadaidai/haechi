package views

import (
	"github.com/daiguadaidai/haechi/middlewares"
	"github.com/daiguadaidai/haechi/views/form"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	handler := new(ReviewHandler)
	AddHandlerV1("/review", handler) // 添加当前页面的uri路径之前都要添加上这个
}

type ReviewHandler struct{}

// 注册route
func (this *ReviewHandler) RegisterV1(group *gin.RouterGroup) {
	group.POST("", this.Review)
	group.GET("", this.Review)
}

// 实现接口, 是否需要进行认证
func (this *ReviewHandler) IsAuth() bool { return false }

func (this *ReviewHandler) Review(c *gin.Context) {
	serverCtx, err := middlewares.GetServerContext(c)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	f := form.NewReviewForm(serverCtx.RuleConfig.Clone())
	if err = c.ShouldBind(f); err != nil {
		returnError(c, http.StatusInternalServerError, err)
	}
	f.CheckDBInfo()
}
