package dao

import (
	"fmt"
	"img-server/core"
	"img-server/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstant *gorm.DB

func init() {
}
func NewDb(c core.DbConfig) {
	log.Println("db connecting")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.DataBase)
	//dsn := "root:123456@tcp(127.0.0.1:3306)/ims?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connecting err:", err.Error())
		return
	}
	err = db.AutoMigrate(model.Menu{},model.User{})
	if err != nil {
		log.Println("db  atuo migrate err", err.Error())
		return
	}
	dbInstant =db
}
func GetDb() *gorm.DB {
	return dbInstant
}
