package dt

import "gin-example/common/dto"

type SysUserGetReq struct {
	dto.Pagination
	ID       int    `json:"id" form:"id"`
	UserName string `json:"userName" form:"userName"`
}
