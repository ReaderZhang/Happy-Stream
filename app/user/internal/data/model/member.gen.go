// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMember = "member"

// Member mapped from table <member>
type Member struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username  string         `gorm:"column:username;not null" json:"username"`
	Phone     string         `gorm:"column:phone;not null" json:"phone"`
	Password  string         `gorm:"column:password;not null" json:"password"`
	Email     string         `gorm:"column:email" json:"email"`
	Avator    string         `gorm:"column:avator;default:https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2F0d%2F96%2F7b%2F0d967b771102473f963e34bd916e5dc0.jpeg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1647714492&t=0a9a941aede934f1f564b895897a0d52" json:"avator"`
	Recharge  float64        `gorm:"column:recharge;default:0.00" json:"recharge"`
	Balance   float64        `gorm:"column:balance;default:0.00" json:"balance"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Rid       int32          `gorm:"column:rid" json:"rid"`
}

// TableName Member's table name
func (*Member) TableName() string {
	return TableNameMember
}