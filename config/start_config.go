package config

import "fmt"

const (
	LISTEN_HOST = "0.0.0.0"
	LISTEN_PORT = 19527
)

type StartConfig struct {
	ListenHost string `json:"listen_host" form:"listen_host"`
	ListenPort int    `json:"listen_port" form:"listen_port"`
	RuleConfig *RuleConfig
}

func NewStartConfig() *StartConfig {
	return &StartConfig{
		RuleConfig: new(RuleConfig),
	}
}

func (this *StartConfig) Addr() string {
	return fmt.Sprintf("%s:%d", this.ListenHost, this.ListenPort)
}
