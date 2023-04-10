package model

import "time"

type User struct {
	ID         string     `gorm:"column:id" json:"id"`
	Name       string     `gorm:"column:name" json:"name"`
	Email      string     `gorm:"column:email" json:"email"`
	Phone      string     `gorm:"column:phone" json:"phone"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime *time.Time `gorm:"column:update_time" json:"update_time"`
}

func (User) TableName() string {
	return "users"
}
