package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "/home/qqz/VSCodeWorkspaces/Happy-Stream/app/order/internal/data/gen",
	})

	db, _ := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/Happy-Stream"))

	g.UseDB(db)

	g.GenerateAllTable()

	g.Execute()
}
