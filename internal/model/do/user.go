// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        interface{} // 用户id
	Username  interface{} // 用户名称
	Password  interface{} // 用户密码
	RealName  interface{} // 真实姓名
	Age       interface{} // 年龄
	Gender    interface{} // 性别 0:未知 1:男 2:女
	Phone     interface{} // 联系电话
	Address   interface{} // 地址
	Status    interface{} // 用户状态 0:禁用 1:启用
	Avatar    interface{} // 头像
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
