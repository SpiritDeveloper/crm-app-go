package repository

import (
	"github.com/jinzhu/gorm"
)

type IFlywayRepository interface {}

type flywayRepository struct {
	DB *gorm.DB
}

func NewFlywayRepository(db *gorm.DB) IFlywayRepository {
	return &flywayRepository{db}
}

