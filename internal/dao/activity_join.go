// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo3/internal/dao/internal"
)

// internalActivityJoinDao is internal type for wrapping internal DAO implements.
type internalActivityJoinDao = *internal.ActivityJoinDao

// activityJoinDao is the data access object for table activity_join.
// You can define custom methods on it to extend its functionality as you wish.
type activityJoinDao struct {
	internalActivityJoinDao
}

var (
	// ActivityJoin is globally public accessible object for table activity_join operations.
	ActivityJoin = activityJoinDao{
		internal.NewActivityJoinDao(),
	}
)

// Fill with you ideas below.
