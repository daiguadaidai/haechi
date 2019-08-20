package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Total   int         `json:"total"`
}

// 返回错误
func returnError(c *gin.Context, code int, err error) {
	resp := &response{Data: make([]interface{}, 0)}
	resp.Status = false
	resp.Message = err.Error()
	c.JSON(code, resp)
}

// 成功并放回数据
func returnSuccess(c *gin.Context, obj interface{}) {
	resp := &response{Data: obj}
	resp.Status = true
	resp.Message = "success"
	c.JSON(http.StatusOK, resp)
}

// 成功并放回数据
func returnList(c *gin.Context, obj interface{}, total int) {
	resp := &response{Data: obj, Total: total}
	resp.Status = true
	resp.Message = "success"
	c.JSON(http.StatusOK, resp)
}

// 返回文件信息
func returnFile(c *gin.Context, fileName string, path string) {
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(path)
}
