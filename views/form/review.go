package form

import (
	"github.com/daiguadaidai/haechi/config"
	"strings"
)

type ReviewForm struct {
	DBHost     string `json:"db_host" form:"db_host"`
	DBPort     int    `json:"db_port" form:"db_port"`
	DBUsername string `json:"db_username" form:"db_username"`
	DBPassword string `json:"db_password" form:"db_password"`
	Database   string `json:"database" form:"database"`

	*config.RuleConfig
}

func NewReviewForm(rc *config.RuleConfig) *ReviewForm {
	return &ReviewForm{
		RuleConfig: rc,
	}
}

func (this *ReviewForm) CheckDBInfo() {
	if strings.TrimSpace(this.DBHost) == "" {
		this.DBHost = config.MYSQL_HOST
	}

	if this.DBPort < 1 {
		this.DBPort = config.MYSQL_PORT
	}

	if strings.TrimSpace(this.DBUsername) == "" {
		this.DBUsername = config.MYSQL_USERNAME
	}

	if strings.TrimSpace(this.DBPassword) == "" {
		this.DBPassword = config.MYSQL_PASSWORD
	}

	if strings.TrimSpace(this.Database) == "" {
		this.Database = config.MYSQL_SCHEMA
	}
}
