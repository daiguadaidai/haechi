package views

import "github.com/gin-gonic/gin"

var handlerV1 map[string]RegistrantV1

// 注册路由
func Register(engine *gin.Engine) {
	v1Group := engine.Group("/api/v1")
	registerV1(v1Group)
}

type RegistrantV1 interface {
	RegisterV1(group *gin.RouterGroup)
	IsAuth() bool
}

// 注册 v1 版本
func registerV1(group *gin.RouterGroup) {
	// 注册任务的路由
	for prefix, handler := range handlerV1 {
		handler.RegisterV1(group.Group(prefix))
	}
}

// 添加 v1 版 handler
func AddHandlerV1(prefix string, h RegistrantV1) {
	if handlerV1 == nil {
		handlerV1 = make(map[string]RegistrantV1)
	}

	handlerV1[prefix] = h
}
