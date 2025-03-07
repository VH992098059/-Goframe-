// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActiveOrder is the golang structure of table active_order for DAO operations like Where/Data.
type ActiveOrder struct {
	g.Meta       `orm:"table:active_order, do:true"`
	Id           interface{} // id
	Status       interface{} // 状态（1：待审核 2：待参加 3：已检票 4：已取消）
	ActivityId   interface{} // 活动id
	CommonUserId interface{} // 普通用户id
	CreateAt     *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 修改时间
	DeleteAt     *gtime.Time // 软删除
}
