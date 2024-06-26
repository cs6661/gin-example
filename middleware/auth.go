package middleware

import (
	"gin-example/pkg/logger"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var getRegexp []regexp.Regexp
var postRegexp []regexp.Regexp
var putRegexp []regexp.Regexp
var deleteRegexp []regexp.Regexp
var allRegexp []regexp.Regexp

func AuthMiddleware(trusts []string) gin.HandlerFunc {
	logger.Logger.Info("加载路由白名单")
	for _, v := range trusts {
		logger.Logger.Debug(v)
		s := strings.Split(v, ":")
		if len(s) != 2 {
			continue
		}
		switch s[0] {
		case "GET":
			getRegexp = append(getRegexp, *regexp.MustCompile(s[1]))
		case "POST":
			postRegexp = append(postRegexp, *regexp.MustCompile(s[1]))
		case "PUT":
			putRegexp = append(putRegexp, *regexp.MustCompile(s[1]))
		case "DELETE":
			deleteRegexp = append(deleteRegexp, *regexp.MustCompile(s[1]))
		case "ALL":
			allRegexp = append(allRegexp, *regexp.MustCompile(s[1]))
		}
	}

	return func(c *gin.Context) {
		var userId int64
		token := c.GetHeader("Authorization")
		if len(token) > 8 {
			token = token[7:]
			//if tokenClaims, err := ParseToken(token); err == nil {
			//	if len(tokenClaims.Subject) > 0 {
			//		// 转int64
			//		userId, _ = strconv.ParseInt(tokenClaims.Subject, 10, 64)
			//		c.Set("userId", userId)
			//	}
			//	c.Set("openId", tokenClaims.OpenId)
			//}
		}

		if len(trusts) != 0 {
			method := c.Request.Method
			path := c.Request.URL.Path
			switch method {
			case "GET":
				for _, v := range getRegexp {
					if v.MatchString(path) {
						logger.Logger.Debug("白名单路由 GET: " + path)
						c.Next()
						return
					}
				}
			case "POST":
				for _, v := range postRegexp {
					if v.MatchString(path) {
						logger.Logger.Debug("白名单路由 POST: " + path)
						c.Next()
						return
					}
				}
			case "PUT":
				for _, v := range putRegexp {
					if v.MatchString(path) {
						logger.Logger.Debug("白名单路由 PUT: " + path)
						c.Next()
						return
					}
				}
			case "DELETE":
				for _, v := range deleteRegexp {
					if v.MatchString(path) {
						logger.Logger.Debug("白名单路由 DELETE: " + path)
						c.Next()
						return
					}
				}
			}
			for _, v := range allRegexp {
				if v.MatchString(path) {
					logger.Logger.Debug("白名单路由 ALL: " + path)
					c.Next()
					return
				}
			}
		}
		if userId == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			return
		}
		c.Next()
	}
}
