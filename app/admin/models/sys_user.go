package models

import "time"

type SysUser struct {
	ID int `json:"id" form:"id"`
}

type User struct {
	ID         int       `json:"id" gorm:"column:id"`                   // 用户ID
	Phone      string    `json:"phone" gorm:"column:phone"`             // 用户电话
	Email      string    `json:"email" gorm:"column:email"`             // 用户邮箱
	Name       string    `json:"name" gorm:"column:name"`               // 用户姓名
	NickName   string    `json:"nick_name" gorm:"column:nick_name"`     // 用户昵称
	Sex        int       `json:"sex" gorm:"column:sex"`                 // 用户性别
	Age        int       `json:"age" gorm:"column:age"`                 // 用户年龄
	UserType   string    `json:"user_type" gorm:"column:user_type"`     // 用户类型
	Password   string    `json:"password" gorm:"column:password"`       // 用户密码
	Avatar     string    `json:"avatar" gorm:"column:avatar"`           // 用户头像
	Deleted    int       `json:"deleted" gorm:"column:deleted"`         // 逻辑删除字段
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 注册时间
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"` // 修改时间
}

func (m *User) TableName() string {
	return "user"
}
