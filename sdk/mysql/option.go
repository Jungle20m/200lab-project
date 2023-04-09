package mysql

import "time"

type Option func(*Mysql)

func WithMaxOpenConnection(size int) Option {
	return func(msql *Mysql) {
		msql.maxOpenConnection = size
	}
}

func WithMaxIdleConnection(size int) Option {
	return func(msql *Mysql) {
		msql.maxIdleConnection = size
	}
}

func WithConnectionMaxLifetime(duration time.Duration) Option {
	return func(msql *Mysql) {
		msql.connectionMaxLifetime = duration
	}
}
