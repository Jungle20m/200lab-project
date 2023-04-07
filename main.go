package main

import (
	"200lab/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID         int        `gorm:"column:id"`
	Name       string     `gorm:"column:name"`
	Email      string     `gorm:"column:email"`
	Phone      string     `gorm:"column:phone"`
	CreateTime *time.Time `gorm:"column:create_time"`
	UpdateTime *time.Time `gorm:"column:update_time"`
}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf)

	db, err := gorm.Open(mysql.Open(conf.Mysql.Dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	var user User

	db.First(&user)

	fmt.Println(user)
}
