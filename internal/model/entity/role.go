// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        uint        `json:"id"        orm:"id"         description:"角色id"`
	Name      string      `json:"name"      orm:"name"       description:"角色名称"`
	Value     string      `json:"value"     orm:"value"      description:"角色值"`
	Status    string      `json:"status"    orm:"status"     description:"角色状态 0:禁用 1:启用"`
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
