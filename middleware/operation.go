package middleware

import (
	"bytes"
	"database/sql"
	"gin-example/pkg/logger"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

type operation struct {
	DB *sql.DB
}

// Operation 操作日志中间件
func Operation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前用户ID 可通过jwt 中间件获取
		handlerName := strings.Split(c.HandlerName(), ".")[1]
		all := strings.ReplaceAll(handlerName, "Router", "")
		switch c.Request.Method {
		case "PUT":
			// Read the Body content
			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request.Body)
			}
			body := strings.Join(strings.Fields(string(bodyBytes)), "")
			// Restore the io.ReadCloser to its original state
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			// 获取修改对象ID
			c.Next()
			if c.Writer.Status() != 200 {
				c.Abort()
				return
			}
			// 增加记录到数据库
			logger.Logger.Info(body + " " + all)

		case "POST":
			// Read the Body content
			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request.Body)
			}
			body := strings.Join(strings.Fields(string(bodyBytes)), "")
			// Restore the io.ReadCloser to its original state
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			c.Next()
			if c.Writer.Status() != 200 {
				c.Abort()
				return
			}
			header := c.GetHeader("Content-Type")
			split := strings.Split(header, ";")
			if split[0] == "multipart/form-data" {
				//c.Abort()
				return
			}
			logger.Logger.Info(body)

		case "DELETE":
			c.Next()
			if c.Writer.Status() != 200 {
				c.Abort()
				return
			}
			rawQuery := c.Request.URL.RawQuery
			split := strings.Split(rawQuery, "=")
			logger.Logger.Info(split[1])

		}
	}
}
