package services

import (
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/haechi/config"
	"github.com/daiguadaidai/haechi/contexts"
	"github.com/daiguadaidai/haechi/middlewares"
	"github.com/daiguadaidai/haechi/views"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Start(sc *config.StartConfig) {
	defer seelog.Flush()
	logger, _ := seelog.LoggerFromConfigAsBytes([]byte(config.LogDefautConfig()))
	seelog.ReplaceLogger(logger)

	httpCtx := contexts.NewHttpContext(sc)

	// 注册路由
	router := gin.Default()
	router.Use(middlewares.Cors())
	router.Use(middlewares.WithServerContext(httpCtx))
	views.Register(router)

	// 获取pala启动配置信息
	s := &http.Server{
		Addr:           sc.Addr(),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	seelog.Infof("Haechi 监听地址为: %v", sc.Addr())
	if err := s.ListenAndServe(); err != nil {
		seelog.Errorf("Haechi 启动服务出错: %v", err.Error())
	}
}
