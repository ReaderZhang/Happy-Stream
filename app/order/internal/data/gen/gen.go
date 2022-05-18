package main

import (
	"fmt"
	"github.com/qqz/Happy-Stream/app/order/internal/data/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/Happy-Stream"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
	}
	db.AutoMigrate(model.Order{})
}
