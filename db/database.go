package db

import (
	"crm-app-go/config"
	"github.com/jinzhu/gorm"
)

type IDatabaseEngine interface {
	// TODO read from config file
	GetDatabase(config config.Database) *gorm.DB
	RunMigration()
}