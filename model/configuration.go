package model

import (
	"github.com/jinzhu/gorm"
)

type Configuration struct {
	gorm.Model
	BusinessUnitId      string    `gorm:"size:255;not null;" json:"businessUnitId"`
	BusinessUnitName      string    `gorm:"size:255;not null;" json:"businessUnitName"`
	TradingPlatformId     string    `gorm:"size:100;not null;unique" json:"tradingPlatformId"`
	BuOwnerId    string    `gorm:"size:100;not null;" json:"buOwnerId"`
	Organization    string    `gorm:"size:100;not null;" json:"organization"`
	Url    string    `gorm:"size:100;not null;" json:"url"`
	User string    `gorm:"size:100;not null;" json:"user"`
	Password string    `gorm:"size:100;not null;" json:"password"`
}