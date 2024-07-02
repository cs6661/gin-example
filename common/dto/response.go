package dto

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Code    MyCode      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func ResponseError(ctx *gin.Context, c MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.AbortWithStatusJSON(520, rd)
}

func HttpError(c *gin.Context, errorCode MyCode, message string) {
	c.AbortWithStatusJSON(520, &ResponseData{
		Code:    errorCode,
		Message: message,
	})
}
