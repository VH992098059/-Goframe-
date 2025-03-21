// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo3/internal/dao/internal"
)

// internalActivityOrderDao is internal type for wrapping internal DAO implements.
type internalActivityOrderDao = *internal.ActivityOrderDao

// activityOrderDao is the data access object for table activity_order.
// You can define custom methods on it to extend its functionality as you wish.
type activityOrderDao struct {
	internalActivityOrderDao
}

var (
	// ActivityOrder is globally public accessible object for table activity_order operations.
	ActivityOrder = activityOrderDao{
		internal.NewActivityOrderDao(),
	}
)

// Fill with you ideas below.
