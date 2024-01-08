package repository

import (
	. "crm-app-go/model"

	"github.com/jinzhu/gorm"
)

type IConfigurationRepository interface {
	GetConfigurationByBrandId(BusinessUnitId string) (*Configuration, error)
}

type configurationRepository struct {
	DB *gorm.DB
}

func NewConfigurationRepository(db *gorm.DB) IConfigurationRepository {
	return &configurationRepository{db}
}

func (configurationRepository *configurationRepository) GetConfigurationByBrandId(BusinessUnitId string) (*Configuration, error) {
	var config Configuration
	result := configurationRepository.DB.Where("business_unit_id = ?", BusinessUnitId).First(&config)
	return &config, result.Error
}
