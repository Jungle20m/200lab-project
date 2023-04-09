package common

import (
	"200lab/config"
	"gorm.io/gorm"
)

type IAppContext interface {
	GetConfig() *config.Config
	GetDB() *gorm.DB
}

type appContext struct {
	Config *config.Config
	DB     *gorm.DB
}

func NewAppContext(config *config.Config, db *gorm.DB) *appContext {
	return &appContext{Config: config, DB: db}
}

func (ctx *appContext) GetConfig() *config.Config {
	return ctx.Config
}

func (ctx *appContext) GetDB() *gorm.DB {
	return ctx.DB
}
