// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserRoleDao is the data access object for the table user_role.
type UserRoleDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns UserRoleColumns // columns contains all the column names of Table for convenient usage.
}

// UserRoleColumns defines and stores column names for the table user_role.
type UserRoleColumns struct {
	Id        string // 用户角色id
	UserId    string // 用户id
	RoleId    string // 角色id
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// userRoleColumns holds the columns for the table user_role.
var userRoleColumns = UserRoleColumns{
	Id:        "id",
	UserId:    "user_id",
	RoleId:    "role_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUserRoleDao creates and returns a new DAO object for table data access.
func NewUserRoleDao() *UserRoleDao {
	return &UserRoleDao{
		group:   "default",
		table:   "user_role",
		columns: userRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserRoleDao) Columns() UserRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
