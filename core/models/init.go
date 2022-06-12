package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine = initDb()

func initDb() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("初始化数据库失败")
		return nil
	}
	return engine
}
