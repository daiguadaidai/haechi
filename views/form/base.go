package form

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/daiguadaidai/haechi/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	defaultStrPage    = "1"
	defaultStrPerPage = "100"
	defaultPage       = 1
	defaultPerPage    = 100
	maxPerPage        = 1000
)

// 分页器
func ParsePaginator(c *gin.Context) *utils.Paginator {
	// 解析第几页
	strPage := c.DefaultQuery("page", defaultStrPage)
	page, err := com.StrTo(strPage).Int64()
	if err != nil {
		page = defaultPage
	}
	if page <= 0 {
		page = defaultPage
	}

	// 解析每页展示条数
	strPerPage := c.DefaultQuery("per_page", defaultStrPerPage)
	perPage, err := com.StrTo(strPerPage).Int64()
	if err != nil {
		perPage = defaultPerPage
	}
	if perPage > maxPerPage {
		perPage = maxPerPage
	}
	return utils.NewPaginator(uint64((page-1)*perPage), uint64(perPage))
}

func NewForm(c *gin.Context, obj interface{}) (interface{}, error) {
	if err := c.ShouldBind(obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// 获取参数
func GetParam(c *gin.Context, key string) (string, error) {
	v := c.Param(key)
	if strings.TrimSpace(v) == "" {
		return "", fmt.Errorf("必须输入参数 %s 值")
	}
	return v, nil
}

// 获取参数
func GetParamInt64(c *gin.Context, key string) (int64, error) {
	v, err := GetParam(c, key)
	if err != nil {
		return 0, err
	}

	i, err := com.StrTo(v).Int64()
	if err != nil {
		return 0, err
	}
	return i, nil
}
