package input

type RegisterLeadFlyway struct {
	BussinesUnitId int    `json:"bussinesUnitId" validate:"required"`
	Name string `json:"name" validate:"required"`
	Last_name string `json:"last_name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone_number string `json:"phone_number" validate:"required"`
	Country_iso string `json:"country_iso" validate:"required"`
	SubAffiliate string `json:"subAffiliate" validate:"required"`
	CampaignId string `json:"campaignId" validate:"required"`
	Tag string `json:"tag" validate:"required"`
	Tag1 string `json:"tag1" validate:"required"`
	AdditionalInfo1 string `json:"additionalInfo1" validate:"required"`
	AdditionalInfo2 string `json:"additionalInfo2" validate:"required"`
	AdditionalInfo3 string `json:"additionalInfo3" validate:"required"`
}
