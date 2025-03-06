// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"        orm:"id"         description:"用户id"`
	Username  string      `json:"username"  orm:"username"   description:"用户名称"`
	Password  string      `json:"password"  orm:"password"   description:"用户密码"`
	RealName  string      `json:"realName"  orm:"real_name"  description:"真实姓名"`
	Age       uint        `json:"age"       orm:"age"        description:"年龄"`
	Gender    string      `json:"gender"    orm:"gender"     description:"性别 0:未知 1:男 2:女"`
	Phone     string      `json:"phone"     orm:"phone"      description:"联系电话"`
	Address   string      `json:"address"   orm:"address"    description:"地址"`
	Status    string      `json:"status"    orm:"status"     description:"用户状态 0:禁用 1:启用"`
	Avatar    string      `json:"avatar"    orm:"avatar"     description:"头像"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
