package service

import (
	"bytes"
	. "crm-app-go/dto/input"
	"crm-app-go/model"
	"crm-app-go/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type IFlywayService interface{
	RegisterLeadInCrm(lead *RegisterLeadFlywayRequestBody) Any
	CrateTransactionInCrm(transaction *CrateTransactionFlyway) Any
}

type flywayService struct {}

var (
	logFlywayRepository           repository.ILogRepository
	configurationFlywayRepository repository.IConfigurationRepository
)

func NewFlywayService(LogRepository repository.ILogRepository, ConfigurationRepository repository.IConfigurationRepository) IFlywayService{
	logFlywayRepository = LogRepository
	configurationFlywayRepository = ConfigurationRepository
	return &flywayService{}
}


func (flywayService *flywayService) RegisterLeadInCrm(lead *RegisterLeadFlywayRequestBody) Any {
	jsonBody, _ := json.Marshal(lead)

	configuration, err := configurationRepository.GetConfigurationByBrandId(strconv.Itoa(lead.BussinesUnitId))
	fmt.Println(configuration.BuOwnerId)
	if err != nil {

		message := "DB error: " + err.Error()

		newLog := &model.Log{
			Action:   "CREATE LEAD",
			Body:     string(jsonBody),
			Response: `{"status": false, "message":` + message + `}`,
			Success:  false,
		}
		logRepository.CreateLog(newLog)

		return nil, &ErrorNewCustomerDto{
			Status:  false,
			Message: message,
		}
	}

	token, err := leverateService.GenerateToken(configuration.Url, configuration.User, configuration.Password)

	if err != nil {

		message := "Token error: " + err.Error()

		newLog := &model.Log{
			Action:   "CREATE LEAD",
			Body:     string(jsonBody),
			Response: `{"status": false, "message":` + message + `}`,
			Success:  false,
		}
		logRepository.CreateLog(newLog)

		return nil, &ErrorNewCustomerDto{
			Status:  false,
			Message: message,
		}
	}

	payload := SendCustomerLeverate{
		TradingPlatformId: configuration.TradingPlatformId,
		IsoCurrency:       customer.IsoCurrency,
		FirstName:         customer.FirstName,
		LastName:          customer.LastName,
		Email:             customer.Email,
		Phone:             customer.Phone,
		Password:          customer.Password,
		IsoCountry:        customer.IsoCountry,
		SubAffiliate:      customer.SubAffiliate,
		RegistrationUrl:   customer.RegistrationUrl,
		Language:          "es",
		CampaignId:        customer.CampaignId,
		Tag:               customer.Tag,
		Tag1:              customer.Tag1,
		AdditionalInfo1:   customer.AdditionalInfo1,
		AdditionalInfo2:   customer.AdditionalInfo2,
		AdditionalInfo3:   customer.AdditionalInfo3,
		BuOwnerId:         configuration.BuOwnerId,
	}

	jsonData, _ := json.Marshal(payload)

	url := "" + configuration.Url + "/accounts/real"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		message := "Error request +" + err.Error()

		newLog := &model.Log{
			Action:   "CREATE LEAD",
			Body:     string(jsonData),
			Response: `{"status": false, "message":` + message + `}`,
			Success:  false,
		}
		logRepository.CreateLog(newLog)
		return nil, &ErrorNewCustomerDto{
			Status:  false,
			Message: message,
		}
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Api-Version", "3")

	response, err := http.DefaultClient.Do(req)

	if err != nil {

		message := "Error request " + err.Error()

		newLog := &model.Log{
			Action:   "CREATE LEAD",
			Body:     string(jsonData),
			Response: `{"status": false, "message":` + message + `}`,
			Success:  false,
		}
		logRepository.CreateLog(newLog)
		return nil, &ErrorNewCustomerDto{
			Status:  false,
			Message: message,
		}
	}
}
	

func (flywayService *flywayService) CrateTransactionInCrm(transaction *CrateTransactionFlyway) Any  {
	return Any{
		Success: true, 
		Message: "", 
		Payload: "",
	}
}
