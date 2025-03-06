// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"goApp/internal/dao/internal"
)

// internalBankDao is an internal type for wrapping the internal DAO implementation.
type internalBankDao = *internal.BankDao

// bankDao is the data access object for the table bank.
// You can define custom methods on it to extend its functionality as needed.
type bankDao struct {
	internalBankDao
}

var (
	// Bank is a globally accessible object for table bank operations.
	Bank = bankDao{
		internal.NewBankDao(),
	}
)

// Add your custom methods and functionality below.
