package input

type Any struct {
    Success interface{} `json:"success" validate:"required"`
	Message interface{} `json:"message" validate:"required"`
	Payload interface{} `json:"payload" validate:"required"`
}

type NewCustomerDto struct {
	BussinesUnitId   string `json:"bussinesUnitId" validate:"required"`
	BussinesUnitName string `json:"bussinesUnitName" validate:"required"`
	IsoCurrency      string `json:"isoCurrency" validate:"required"`
	FirstName        string `json:"firstName" validate:"required"`
	LastName         string `json:"lastName" validate:"required"`
	Email            string `json:"email" validate:"required"`
	Password         string `json:"password"`
	Phone            string `json:"phone" validate:"required"`
	IsoCountry       string `json:"isoCountry" validate:"required"`
	SubAffiliate     string `json:"subAffiliate" validate:"required"`
	RegistrationUrl  string `json:"registrationUrl" validate:"required"`
	CampaignId       string `json:"campaignId" validate:"required"`
	Tag              string `json:"tag"`
	Tag1             string `json:"tag1"`
	AdditionalInfo1  string `json:"additionalInfo1"`
	AdditionalInfo2  string `json:"additionalInfo2"`
	AdditionalInfo3  string `json:"additionalInfo3"`
}

type GenerateTokenDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SendCustomerLeverate struct {
	TradingPlatformId string `json:"tradingPlatformId"`
	IsoCurrency       string `json:"isoCurrency"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Password          string `json:"password"`
	IsoCountry        string `json:"isoCountry"`
	SubAffiliate      string `json:"subAffiliate"`
	RegistrationUrl   string `json:"registrationUrl"`
	Language          string `json:"language"`
	CampaignId        string `json:"campaignId"`
	Tag               string `json:"tag"`
	Tag1              string `json:"tag1"`
	AdditionalInfo1   string `json:"additionalInfo1"`
	AdditionalInfo2   string `json:"additionalInfo2"`
	AdditionalInfo3   string `json:"additionalInfo3"`
	BuOwnerId         string `json:"buOwnerId"`
}