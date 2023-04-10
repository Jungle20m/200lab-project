package storage

import (
	"gorm.io/gorm"
)

type Storage struct {
	db    *gorm.DB
	other string
}

func NewMysqlStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func NewFullStorage(db *gorm.DB, otherDB string) *Storage {
	return &Storage{
		db:    db,
		other: otherDB,
	}
}
