// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"goApp/internal/dao/internal"
)

// internalProductDao is an internal type for wrapping the internal DAO implementation.
type internalProductDao = *internal.ProductDao

// productDao is the data access object for the table product.
// You can define custom methods on it to extend its functionality as needed.
type productDao struct {
	internalProductDao
}

var (
	// Product is a globally accessible object for table product operations.
	Product = productDao{
		internal.NewProductDao(),
	}
)

// Add your custom methods and functionality below.
