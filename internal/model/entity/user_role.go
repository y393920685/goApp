// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRole is the golang structure for table user_role.
type UserRole struct {
	Id        uint        `json:"id"        orm:"id"         description:"用户角色id"`
	UserId    uint        `json:"userId"    orm:"user_id"    description:"用户id"`
	RoleId    uint        `json:"roleId"    orm:"role_id"    description:"角色id"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
