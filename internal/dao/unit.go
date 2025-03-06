// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"goApp/internal/dao/internal"
)

// internalUnitDao is an internal type for wrapping the internal DAO implementation.
type internalUnitDao = *internal.UnitDao

// unitDao is the data access object for the table unit.
// You can define custom methods on it to extend its functionality as needed.
type unitDao struct {
	internalUnitDao
}

var (
	// Unit is a globally accessible object for table unit operations.
	Unit = unitDao{
		internal.NewUnitDao(),
	}
)

// Add your custom methods and functionality below.
