package model

import (
	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model
	Token      string    `gorm:"size:255;not null;" json:"token"`
	Organization     string    `gorm:"size:100;not null;unique" json:"organization"`
}