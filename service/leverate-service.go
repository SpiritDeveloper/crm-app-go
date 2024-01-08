package service

import (
	"bytes"
	. "crm-app-go/dto/input"
	. "crm-app-go/dto/output"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"crm-app-go/model"
	"crm-app-go/repository"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ILeverateService interface {
	SendLeadToCrm(customer *NewCustomerDto) (*ResponseNewCustomerDto, *ErrorNewCustomerDto)
	GenerateToken(url string, user string, password string) (*LeverateGenerateTokenResponse, error)
}

type leverateService struct{}

var (
	logRepository           repository.ILogRepository
	configurationRepository repository.IConfigurationRepository
)

func NewLeverateService(LogRepository repository.ILogRepository, ConfigurationRepository repository.IConfigurationRepository) ILeverateService {
	logRepository = LogRepository
	configurationRepository = ConfigurationRepository
	return &leverateService{}
}

func (leverateService *leverateService) GenerateToken(url string, user string, password string) (*LeverateGenerateTokenResponse, error) {
	url_leverate := "" + url + "/token"

	payload := GenerateTokenDto{
		Username: user,
		Password: password,
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url_leverate, bytes.NewBuffer(jsonData))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var apiResponse LeverateGenerateTokenResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &LeverateGenerateTokenResponse{
		Token:     apiResponse.Token,
		ExpiresIn: apiResponse.ExpiresIn,
	}, nil
}

func (leverateService *leverateService) SendLeadToCrm(customer *NewCustomerDto) (*ResponseNewCustomerDto, *ErrorNewCustomerDto) {

	jsonBody, _ := json.Marshal(customer)

	configuration, err := configurationRepository.GetConfigurationByBrandId(customer.BussinesUnitId)
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

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		newLog := &model.Log{
			Action:   "CREATE LEAD",
			Body:     string(jsonData),
			Response: "",
			Success:  false,
		}
		logRepository.CreateLog(newLog)
		return nil, &ErrorNewCustomerDto{
			Status:  false,
			Message: "Response JSON not create",
		}
	}

	if response.StatusCode != http.StatusOK {
		return nil, &ErrorNewCustomerDto{
			Status:  false,
			Message: string(body),
		}
	}

	var apiResponse LeverateCreateCustomerResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, &ErrorNewCustomerDto{
			Status:  false,
			Message: "Customer already exist in crm",
		}
	}

	jsonResponse, err := json.Marshal(apiResponse)
	newLog := &model.Log{
		Action:   "CREATE LEAD",
		Body:     string(jsonData),
		Response: string(jsonResponse),
		Success:  true,
	}
	logRepository.CreateLog(newLog)

	return &ResponseNewCustomerDto{
		Tpid:    apiResponse.TpAccountName,
		Message: "Customer create successfully",
	}, nil

}
