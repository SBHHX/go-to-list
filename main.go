package main

import (
	"go_to_list/db"
	"go_to_list/routers"
)

func main() {
	// 替换为实际数据库DSN
	dsn := "root:cyhdeqq,201496@tcp(127.0.0.1:3306)/todo_list_db?charset=utf8mb4&parseTime=True&loc=Local"
	if err := db.Connect(dsn); err != nil {
		panic(err)
	}
	r := routers.SetupRouter()
	r.Run(":8080")
}
