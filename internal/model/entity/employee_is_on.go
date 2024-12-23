// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeIsOn is the golang structure for table employee_is_on.
type EmployeeIsOn struct {
	EmployeeId string      `json:"employeeId" orm:"employee_id" description:"职工编号"`
	UserId     string      `json:"userId"     orm:"user_id"     description:"用户编号"`
	Name       string      `json:"name"       orm:"name"        description:"姓名"`
	IsOn       string      `json:"isOn"       orm:"is_on"       description:"是否在职"`
	SchoolId   string      `json:"schoolId"   orm:"school_id"   description:"学校编号"`
	CreateAt   *gtime.Time `json:"createAt"   orm:"create_at"   description:"创建时间"`
	UpdateAt   *gtime.Time `json:"updateAt"   orm:"update_at"   description:"修改时间"`
	DeleteAt   *gtime.Time `json:"deleteAt"   orm:"delete_at"   description:"软删除"`
}
