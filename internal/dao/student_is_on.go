// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo3/internal/dao/internal"
)

// internalStudentIsOnDao is internal type for wrapping internal DAO implements.
type internalStudentIsOnDao = *internal.StudentIsOnDao

// studentIsOnDao is the data access object for table student_is_on.
// You can define custom methods on it to extend its functionality as you wish.
type studentIsOnDao struct {
	internalStudentIsOnDao
}

var (
	// StudentIsOn is globally public accessible object for table student_is_on operations.
	StudentIsOn = studentIsOnDao{
		internal.NewStudentIsOnDao(),
	}
)

// Fill with you ideas below.
