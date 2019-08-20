package config

import (
	"fmt"
	"strings"
)

const (
	MYSQL_HOST           = "127.0.0.1"
	MYSQL_PORT           = 3306
	MYSQL_USERNAME       = "root"
	MYSQL_PASSWORD       = "root"
	MYSQL_SCHEMA         = ""
	MYSQL_AUTO_COMMIT    = true
	MYSQL_MAX_OPEN_CONNS = 100
	MYSQL_MAX_IDEL_CONNS = 100
	MYSQL_CHARSET        = "utf8mb4"
	MYSQL_TIMEOUT        = 10
)

type MySQLConfig struct {
	Username          string `toml:"username"`
	Password          string `toml:"password"`
	Database          string `toml:"database"`
	Charset           string `toml:"charset"`
	Host              string `toml:"host"`
	Timeout           int    `toml:"timeout"`
	Port              int    `toml:"port"`
	MaxOpenConns      int    `toml:"max_open_conns"`
	MaxIdelConns      int    `toml:"max_idel_conns"`
	AllowOldPasswords int    `toml:"allow_old_passwords"`
	AutoCommit        bool   `toml:"auto_commit"`
}

/* 新建一个数据库执行器
Params:
    host: ip
    port: 端口
    username: 链接数据库用户名
    password: 链接数据库密码
    database: 要操作的数据库
*/
func NewMySQLConfig(
	host string,
	port int,
	username string,
	password string,
	database string,
	charset string,
	autoCommit bool,
	timeout int,
	maxOpenConns int,
	maxIdelConns int,
) *MySQLConfig {
	dbConfig := &MySQLConfig{
		Username:          username,
		Password:          password,
		Host:              host,
		Port:              port,
		Database:          database,
		Charset:           charset,
		MaxOpenConns:      maxOpenConns,
		MaxIdelConns:      maxIdelConns,
		Timeout:           timeout,
		AllowOldPasswords: 1,
		AutoCommit:        autoCommit,
	}

	return dbConfig
}

func (this *MySQLConfig) GetDataSource() string {
	dataSource := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=%v&allowOldPasswords=%v&timeout=%vs&autocommit=%v&parseTime=True&loc=Local",
		this.Username,
		this.Password,
		this.Host,
		this.Port,
		this.Database,
		this.Charset,
		this.AllowOldPasswords,
		this.Timeout,
		this.AutoCommit,
	)

	return dataSource
}

func (this *MySQLConfig) Check() error {
	if strings.TrimSpace(this.Database) == "" {
		return fmt.Errorf("数据库不能为空")
	}

	return nil
}

// 补充默认值
func (this *MySQLConfig) SupDefault() {
	if len(this.Username) == 0 {
		this.Username = MYSQL_USERNAME
	}
	if len(this.Password) == 0 {
		this.Password = MYSQL_PASSWORD
	}
	if len(this.Charset) == 0 {
		this.Charset = MYSQL_CHARSET
	}
	if len(this.Host) == 0 {
		this.Host = MYSQL_HOST
	}
	if this.Port < 1 {
		this.Port = MYSQL_PORT
	}
	if this.Timeout < 0 {
		this.Timeout = MYSQL_TIMEOUT
	}
	if this.MaxOpenConns < 1 {
		this.MaxOpenConns = MYSQL_MAX_OPEN_CONNS
	}
	if this.MaxIdelConns < 1 {
		this.MaxIdelConns = MYSQL_MAX_IDEL_CONNS
	}
	this.AllowOldPasswords = 1
	this.AutoCommit = MYSQL_AUTO_COMMIT
}
