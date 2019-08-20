package contexts

import (
	"github.com/daiguadaidai/haechi/config"
)

type HttpContext struct {
	ServerConfig *config.StartConfig
	RuleConfig   *config.RuleConfig
}

func NewHttpContext(sc *config.StartConfig) *HttpContext {
	return &HttpContext{
		ServerConfig: sc,
		RuleConfig:   sc.RuleConfig,
	}
}
