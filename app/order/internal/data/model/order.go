package model

import "time"

const TableNameOrder = "order"

type Order struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true"json:"id"`
	Text      string    `gorm:"column:text;not null"json:"text"`
	Status    int32     `gorm:"column:status;default:1"json:"status"`
	UserId    int32     `gorm:"column:user_id"json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at"json:"created_at"`
}

func (*Order) TableName() string {
	return TableNameOrder
}
