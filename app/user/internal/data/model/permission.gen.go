// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePermission = "permission"

// Permission mapped from table <permission>
type Permission struct {
	ID       int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ParentID string `gorm:"column:parent_id" json:"parent_id"`
	URL      string `gorm:"column:url" json:"url"`
}

// TableName Permission's table name
func (*Permission) TableName() string {
	return TableNamePermission
}
