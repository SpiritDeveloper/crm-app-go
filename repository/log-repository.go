package repository

import (
	. "crm-app-go/model"

	"github.com/jinzhu/gorm"
)

type ILogRepository interface {
	CreateLog(log *Log) (*Log, error)
}

type logRepository struct {
	DB *gorm.DB
}

func NewLogRepository(db *gorm.DB) ILogRepository {
	return &logRepository{db}
}

func (logRepository *logRepository) CreateLog(log *Log) (*Log, error) {
	result := logRepository.DB.Create(log)
	return log, result.Error
}
