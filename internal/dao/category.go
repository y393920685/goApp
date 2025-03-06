// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"goApp/internal/dao/internal"
)

// internalCategoryDao is an internal type for wrapping the internal DAO implementation.
type internalCategoryDao = *internal.CategoryDao

// categoryDao is the data access object for the table category.
// You can define custom methods on it to extend its functionality as needed.
type categoryDao struct {
	internalCategoryDao
}

var (
	// Category is a globally accessible object for table category operations.
	Category = categoryDao{
		internal.NewCategoryDao(),
	}
)

// Add your custom methods and functionality below.
