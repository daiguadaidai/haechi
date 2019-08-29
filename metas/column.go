package metas

import (
	"fmt"
	"github.com/daiguadaidai/haechi/utils"
	"strings"
)

const (
	STR_TYPE_DECIMAL    string = "DECIMAL"
	STR_TYPE_TINY       string = "TINYINT"
	STR_TYPE_SHORT      string = "SMALLINT"
	STR_TYPE_LONG       string = "INT"
	STR_TYPE_FLOAT      string = "FLOAT"
	STR_TYPE_DOUTBLE    string = "DOUTBLE"
	STR_TYPE_NULL       string = "NULL"
	STR_TYPE_TIMESTAMP  string = "TIMESTAMP"
	STR_TYPE_BIGINT     string = "BIGINT"
	STR_TYPE_MEDIUMINT  string = "MEDIUMINT"
	STR_TYPE_DATE       string = "DATE"
	STR_TYPE_DURATION   string = "TIME"
	STR_TYPE_DATETIME   string = "DATETIME"
	STR_TYPE_YEAR       string = "YEAR"
	STR_TYPE_VARCHAR    string = "VARCHAR"
	STR_TYPE_BIT        string = "BIT"
	STR_TYPE_JSON       string = "JSON"
	STR_TYPE_ENUM       string = "ENUM"
	STR_TYPE_SET        string = "SET"
	STR_TYPE_TINYBLOB   string = "TINYBLOB"
	STR_TYPE_MEDIUMBLOB string = "MEDIUMBLOB"
	STR_TYPE_LONGBLOB   string = "LONGBLOB"
	STR_TYPE_BLOB       string = "BLOB"
	STR_TYPE_VAR_STRING string = "VARSTRING"
	STR_TYPE_STRING     string = "STRING"
	STR_TYPE_GEOMETRY   string = "GEOMETRY"
	STR_TYPE_TINYTEXT   string = "TINYTEXT"
	STR_TYPE_TEXT       string = "TEXT"
	STR_TYPE_MEDIUMTEXT string = "MEDIUMTEXT"
	STR_TYPE_LONGTEXT   string = "LONGTEXT"
)

var strTypeMap map[string]struct{} = map[string]struct{}{
	STR_TYPE_DECIMAL:    struct{}{},
	STR_TYPE_TINY:       struct{}{},
	STR_TYPE_SHORT:      struct{}{},
	STR_TYPE_LONG:       struct{}{},
	STR_TYPE_FLOAT:      struct{}{},
	STR_TYPE_DOUTBLE:    struct{}{},
	STR_TYPE_NULL:       struct{}{},
	STR_TYPE_TIMESTAMP:  struct{}{},
	STR_TYPE_BIGINT:     struct{}{},
	STR_TYPE_MEDIUMINT:  struct{}{},
	STR_TYPE_DATE:       struct{}{},
	STR_TYPE_DURATION:   struct{}{},
	STR_TYPE_DATETIME:   struct{}{},
	STR_TYPE_YEAR:       struct{}{},
	STR_TYPE_VARCHAR:    struct{}{},
	STR_TYPE_BIT:        struct{}{},
	STR_TYPE_JSON:       struct{}{},
	STR_TYPE_ENUM:       struct{}{},
	STR_TYPE_SET:        struct{}{},
	STR_TYPE_TINYBLOB:   struct{}{},
	STR_TYPE_MEDIUMBLOB: struct{}{},
	STR_TYPE_LONGBLOB:   struct{}{},
	STR_TYPE_BLOB:       struct{}{},
	STR_TYPE_VAR_STRING: struct{}{},
	STR_TYPE_STRING:     struct{}{},
	STR_TYPE_GEOMETRY:   struct{}{},
	STR_TYPE_TINYTEXT:   struct{}{},
	STR_TYPE_TEXT:       struct{}{},
	STR_TYPE_MEDIUMTEXT: struct{}{},
	STR_TYPE_LONGTEXT:   struct{}{},
}

func TypeExists(t string) (string, bool) {
	upType := strings.ToUpper(t)
	if _, ok := strTypeMap[upType]; ok {
		return upType, ok
	}

	return upType, false
}

// 判断类型是否有小数
func TypeHaveDecimal(t string) bool {
	switch strings.ToUpper(t) {
	case STR_TYPE_FLOAT, STR_TYPE_DOUTBLE, STR_TYPE_DECIMAL:
		return true
	}
	return false
}

// 判断类型是否有长度
func TypeHaveLen(t string) bool {
	switch strings.ToUpper(t) {
	case STR_TYPE_DECIMAL, STR_TYPE_TINY, STR_TYPE_SHORT, STR_TYPE_LONG, STR_TYPE_FLOAT, STR_TYPE_DOUTBLE, STR_TYPE_NULL,
		STR_TYPE_TIMESTAMP, STR_TYPE_BIGINT, STR_TYPE_MEDIUMINT, STR_TYPE_DATE, STR_TYPE_DURATION, STR_TYPE_DATETIME, STR_TYPE_YEAR,
		STR_TYPE_VARCHAR, STR_TYPE_BIT, STR_TYPE_VAR_STRING, STR_TYPE_STRING:
		return true
	}
	return false
}

// 判断类型是否有values, ENUM('a', 'b', 'c'), SET('a', 'b', 'c')
func TypeHaveValues(t string) bool {
	switch strings.ToUpper(t) {
	case STR_TYPE_ENUM, STR_TYPE_SET:
		return true
	}
	return false
}

func TypeHaveDefaultValue(t string) bool {
	switch strings.ToUpper(t) {
	case STR_TYPE_TINYBLOB, STR_TYPE_BLOB, STR_TYPE_MEDIUMBLOB, STR_TYPE_LONGBLOB, STR_TYPE_TINYTEXT, STR_TYPE_TEXT, STR_TYPE_MEDIUMTEXT,
		STR_TYPE_LONGTEXT, STR_TYPE_JSON, STR_TYPE_GEOMETRY:
		return false
	}
	return true
}

type Column struct {
	Name               string   `json:"name" form:"name"`
	Type               string   `json:"type" form:"type"`
	TypeLen            int      `json:"type_len" form:"type_len"`
	TypeDecimal        int      `json:"type_decimal" form:"type_decimal"`
	TypeValues         []string `json:"type_values" form:"type_values"` // ENUM('a', 'b', 'c'), SET('a', 'b', 'c')
	IsNull             bool     `json:"is_null" form:"is_null"`
	HaveDefaultValue   bool     `json:"have_default_value" form:"have_default_value"` // 是否有默认值
	DefaultNull        bool     `json:"default_null" form:"default_null"`             // 默认值是否为NULL
	DefaultValue       string   `json:"default_value" form:"default_value"`           // 默认值
	DefaultValueIsFunc bool     `json:"default_value_is_func" form:"default_value_is_func"`
	OnUpdateValue      string   `json:"on_update_value" form:"on_update_value"`
	Charset            string   `json:"charset" form:"charset"`
	Collate            string   `json:"collate" form:"collate"`
	Comment            string   `json:"comment" form:"comment"`
}

func NewColumn(name string) *Column {
	return &Column{
		Name: name,
	}
}

func (this *Column) SetType(t string, length, decimal int, values []string) error {
	t, ok := TypeExists(t)
	if !ok {
		return fmt.Errorf("设置字段类型失败. 类型:%s 不合法", t)
	}

	this.Type = t
	this.TypeLen = length
	this.TypeDecimal = decimal
	this.TypeValues = values

	return nil
}

func (this *Column) SetDefaultValue(haveValue, defaultNull bool, value string, isFunc bool) {
	this.HaveDefaultValue = haveValue
	this.DefaultNull = defaultNull
	this.DefaultValue = value
	this.DefaultValueIsFunc = isFunc
}

func (this *Column) SetOnUpdateValue(value string) {
	this.OnUpdateValue = value
}

func (this *Column) SetCharset(charset string) {
	this.Charset = charset
}

func (this *Column) SetCollate(collate string) {
	this.Collate = collate
}

func (this *Column) SetComment(comment string) {
	this.Comment = comment
}

// 获取整个字段定义的字符串
func (this *Column) GetMetaStr() (string, error) {
	items := make([]string, 0, 2)
	// 字段名字
	items = append(items, fmt.Sprintf("`%s`", this.Name))

	// 字段类型
	t, err := this.GetTypeStr()
	if err != nil {
		return "", err
	}
	items = append(items, t)

	// character set utf8mb4
	if this.Charset != "" {
		items = append(items, fmt.Sprintf("CHARACTER SET %s", this.Charset))
	}

	// collate utf8mb4_unicode_ci
	if this.Collate != "" {
		items = append(items, fmt.Sprintf("COLLATE %s", this.Collate))
	}

	// 默认值
	defaultValue, ok := this.GetDefaultValue()
	if ok {
		items = append(items, fmt.Sprintf("DEFAULT %s", defaultValue))
	}

	if strings.TrimSpace(this.OnUpdateValue) != "" {
		items = append(items, fmt.Sprintf("ON UPDATE %s", this.OnUpdateValue))
	}

	items = append(items, fmt.Sprintf("COMMENT %#v", this.GetComment()))

	return strings.Join(items, " "), nil
}

// 获取字段类型
func (this *Column) GetTypeStr() (string, error) {
	t, ok := TypeExists(this.Type)
	if !ok {
		return "", fmt.Errorf("未识别到的字段类型: %s", this.Type)
	}

	if TypeHaveDecimal(t) {
		if this.TypeLen > 0 && this.TypeDecimal > 0 {
			return fmt.Sprintf("%s(%d, %d)", t, this.TypeLen, this.TypeDecimal), nil
		} else if this.TypeLen > 0 {
			return fmt.Sprintf("%s(%d)", t, this.TypeLen), nil
		}
	} else if TypeHaveLen(t) {
		if this.TypeLen > 0 {
			return fmt.Sprintf("%s(%d)", t, this.TypeLen), nil
		}
	} else if TypeHaveValues(t) {
		if this.TypeValues == nil || len(this.TypeValues) == 0 {
			return "", fmt.Errorf("类型: %[1]s. 需要有值. 如: %[1]s('a', 'b', 'c')", this.Type)
		}
		valuesStr := strings.Join(utils.EnumStrs(this.TypeValues), ", ")
		return fmt.Sprintf("%s(%s)", this.Type, valuesStr), nil
	}

	return t, nil
}

// 获取默认值
func (this *Column) GetDefaultValue() (string, bool) {
	if !TypeHaveDefaultValue(this.Type) {
		return "", false
	}

	if !this.HaveDefaultValue {
		return "", false
	}

	if this.DefaultNull {
		return "NULL", true
	}

	if this.DefaultValueIsFunc {
		return fmt.Sprintf("%v", this.DefaultValue), true
	}

	return fmt.Sprintf("%#v", this.DefaultValue), true
}

// 获取注释
func (this *Column) GetComment() string {
	if this.Comment == "" {
		return fmt.Sprint("请填写注释")
	}

	return this.Comment
}
