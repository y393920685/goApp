// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Brand is the golang structure for table brand.
type Brand struct {
	Id        uint        `json:"id"        orm:"id"         description:"品牌id"`
	Name      string      `json:"name"      orm:"name"       description:"品牌名称"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
