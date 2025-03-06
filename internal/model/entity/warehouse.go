// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Warehouse is the golang structure for table warehouse.
type Warehouse struct {
	Id        uint        `json:"id"        orm:"id"         description:"仓库id"`
	Name      string      `json:"name"      orm:"name"       description:"仓库名称"`
	Contact   string      `json:"contact"   orm:"contact"    description:"仓库联系人"`
	Phone     string      `json:"phone"     orm:"phone"      description:"仓库联系电话"`
	Status    string      `json:"status"    orm:"status"     description:"仓库状态 0:禁用 1:启用"`
	Active    string      `json:"active"    orm:"active"     description:"仓库默认 0:否 1:是"`
	Address   string      `json:"address"   orm:"address"    description:"仓库地址"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
