package mysql

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

const (
	_defaultMaxOpenConnection     = 100
	_defaultMaxIdleConnection     = 10
	_defaultConnectionMaxLifetime = time.Hour
)

type Mysql struct {
	maxOpenConnection     int
	maxIdleConnection     int
	connectionMaxLifetime time.Duration

	DB *gorm.DB
}

func New(dns string, opts ...Option) (*Mysql, error) {
	//anhnv:anhnv!@#456@tcp(1.53.252.177:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local

	msql := &Mysql{
		maxOpenConnection:     _defaultMaxOpenConnection,
		maxIdleConnection:     _defaultMaxIdleConnection,
		connectionMaxLifetime: _defaultConnectionMaxLifetime,
	}

	for _, opt := range opts {
		opt(msql)
	}

	sqlDB, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(msql.maxOpenConnection)
	sqlDB.SetMaxIdleConns(msql.maxIdleConnection)
	sqlDB.SetConnMaxLifetime(msql.connectionMaxLifetime)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	msql.DB = gormDB

	return msql, err
}
