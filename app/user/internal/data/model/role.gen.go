// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRole = "role"

// Role mapped from table <role>
type Role struct {
	ID       int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoleName string `gorm:"column:role_name" json:"role_name"`
	Pid      int32  `gorm:"column:pid;default:1" json:"pid"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}