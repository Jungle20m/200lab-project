package component

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type AppContext struct {
	db  *sqlx.DB
	mux *chi.Mux
}

func NewAppContext(db *sqlx.DB, mux *chi.Mux) *AppContext {
	return &AppContext{db, mux}
}

func (ctx *AppContext) GetMysqlConnector() *sqlx.DB {
	return ctx.db
}

func (ctx *AppContext) GetRestHandler() *chi.Mux {
	return ctx.mux
}
