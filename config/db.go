package config

import (
	_activity "TodoApp/features/activity/data"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(ac AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ac.MYSQL_DBUSER, ac.MYSQL_DBPASS, ac.MYSQL_DBHOST, ac.MYSQL_DBPORT, ac.MYSQL_DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(_activity.Activity{})
}
