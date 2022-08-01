package storage

import "github.com/jmoiron/sqlx"

type storage struct {
	db *sqlx.DB
}

func NewStorage() *storage {
	return &storage{db: nil}
}

func (s *storage) SetMysqlConnector(mysqlConnector *sqlx.DB) {
	s.db = mysqlConnector
}
