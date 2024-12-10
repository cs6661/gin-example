package middleware

import (
	"bytes"
	"context"
	"fmt"
	"gin-example/pkg/constant"
	"gin-example/pkg/logger"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.New().String()
		c.Set(constant.TraceId, traceId)
		logs := logger.Logger.Named(traceId)
		// 创建新的 context.Context，将 traceID 放入 context 中
		ctx := context.WithValue(c.Request.Context(), constant.TraceId, traceId)

		// 将带有 traceID 的 context.Context 传递给后续的处理函数
		c.Request = c.Request.WithContext(ctx)
		var url = fmt.Sprintf("%s %s", c.Request.Method, c.Request.RequestURI)
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			if c.Request.Header.Get("Content-Type") == "application/json" {
				var body []byte
				var buf bytes.Buffer
				tee := io.TeeReader(c.Request.Body, &buf)
				body, _ = io.ReadAll(tee)
				c.Request.Body = io.NopCloser(&buf)
				str := string(body)
				logs.Info(fmt.Sprintf(`request: %s
				 %s`, url, str))
			} else {
				logs.Info(fmt.Sprintf("request: %s", url))
			}

		} else {
			logs.Info(fmt.Sprintf("request: %s", url))
		}

		start := time.Now()
		c.Next()
		tc := time.Since(start)
		logs.Info(fmt.Sprintf("耗时: %s [ %v ]", url, tc))
	}
}
