package config

const (
	RULE_NAME_LENGTH = 64
)

type RuleConfig struct {
	// 通用名字长度
	RuleNameLength int `json:"rule_name_length" form:"rule_name_length"`
}

func (this *RuleConfig) Clone() *RuleConfig {
	tmpRuleConfig := *this
	return &tmpRuleConfig
}
