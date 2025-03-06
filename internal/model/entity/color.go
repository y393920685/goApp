// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Color is the golang structure for table color.
type Color struct {
	Id        uint        `json:"id"        orm:"id"         description:"颜色id"`
	Name      string      `json:"name"      orm:"name"       description:"颜色名称"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
