package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Log struct {
	gorm.Model
	Action   string    `gorm:"size:255;not null;" json:"action"`
	Body     string    `gorm:"type:json" json:"body"`
	Response string    `gorm:"type:json" json:"response"`
	Success  bool      `json:"success"`
	CreateOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createOn"`
}
