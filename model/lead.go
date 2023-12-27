package model

import (
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	BusinessUnitId      string    `gorm:"size:255;not null;" json:"businessUnitId"`
	Tpid     string    `gorm:"size:100;not null;unique" json:"tpid"`
	AccountId    string    `gorm:"size:100;not null;" json:"accountId"`
	Name    string    `gorm:"size:100;not null;" json:"name"`
	Lastname    string    `gorm:"size:100;not null;" json:"lastName"`
	Email    string    `gorm:"size:100;not null;" json:"email"`
	Phone    string    `gorm:"size:100;not null;" json:"phone"`
	Password    string    `gorm:"size:100;not null;" json:"password"`
	IsoCountry    string    `gorm:"size:100;not null;" json:"isoCountry"`
	SubAffiliate    string    `gorm:"size:100;not null;" json:"subAffiliate"`
	RegistrationUrl    string    `gorm:"size:100;not null;" json:"registrationUrl"`
	CampaignId    string    `gorm:"size:100;not null;" json:"campaignId"`
	Tag    string    `gorm:"size:100;not null;" json:"tag"`
	Tag1    string    `gorm:"size:100;not null;" json:"tag1"`
	AdditionalInfo1    string    `gorm:"size:100;not null;" json:"additionalInfo1"`
	AdditionalInfo2    string    `gorm:"size:100;not null;" json:"additionalInfo2"`
	AdditionalInfo3    string    `gorm:"size:100;not null;" json:"additionalInfo3"`
}