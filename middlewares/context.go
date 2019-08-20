package middlewares

import (
	"fmt"
	"github.com/daiguadaidai/haechi/contexts"
	"github.com/gin-gonic/gin"
)

// ServerContextKey is the key in gin.Context for ServerContext.
const HttpContextKey = "http-context"

// WithServerContext injects the ServerContext instance into gin.Context.
func WithServerContext(httpCtx *contexts.HttpContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(HttpContextKey, httpCtx)
		ctx.Next()
	}
}

// GetServerContext fetches the ServerContext instance from middlewares.
func GetServerContext(ctx *gin.Context) (*contexts.HttpContext, error) {
	v, exists := ctx.Get(HttpContextKey)
	if !exists {
		return nil, fmt.Errorf("no server context found in gin.context")
	}
	vv, ok := v.(*contexts.HttpContext)
	if !ok {
		return nil, fmt.Errorf("invalid worker pool instance: %v", v)
	}
	return vv, nil
}
