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
}

type leverateService struct{}

var (
	logRepository repository.ILogRepository
)

func NewLeverateService(repository repository.ILogRepository) ILeverateService {
	logRepository = repository
	return &leverateService{}
}

type leverateCreateCustomerResponse struct {
	AccountId         string `json:"accountId"`
	TpAccountName     string `json:"tpAccountName"`
	TpAccountPassword string `json:"tpAccountPassword"`
}

func (leverateService *leverateService) SendLeadToCrm(customer *NewCustomerDto) (*ResponseNewCustomerDto, *ErrorNewCustomerDto) {

	payload := NewCustomerDto{
		BussinesUnitId:   customer.BussinesUnitId,
		IsoCurrency:      customer.IsoCurrency,
		BussinesUnitName: customer.BussinesUnitName,
		FirstName:        customer.FirstName,
		LastName:         customer.LastName,
		Email:            customer.Email,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	if err := validate.Struct(customer); err != nil {
		fmt.Println("Validation error:", err)

		newLog := &model.Log{
			Action:   "CREATE LEAD",
			Body:     string(jsonData),
			Response: `{"status": false, "message": "Error Body"}`,
			Success:  false,
		}
		logRepository.CreateLog(newLog)

		return nil, &ErrorNewCustomerDto{
			Status:  false,
			Message: "Validation error: " + err.Error(),
		}
	}

	url := "https://0373-148-244-126-218.ngrok-free.app/test"

	contentType := "application/json"

	response, err := http.Post(url, contentType, bytes.NewBuffer(jsonData))

	if err != nil {
		newLog := &model.Log{
			Action:   "CREATE LEAD",
			Body:     string(jsonData),
			Response: `{"status": false, "message": "Error request"}`,
			Success:  false,
		}
		logRepository.CreateLog(newLog)
		return nil, &ErrorNewCustomerDto{
			Status:  false,
			Message: "Error to create customer in leverate",
		}
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {

		newLog := &model.Log{
			Action:   "CREATE LEAD",
			Body:     string(jsonData),
			Response: "",//string(response.Body),
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
			Message: "Customer already exist in crm",
		}
	}

	var apiResponse leverateCreateCustomerResponse

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, &ErrorNewCustomerDto{
			Status:  false,
			Message: "Customer already exist in crm",
		}
	}

	return &ResponseNewCustomerDto{
		Tpid:    apiResponse.AccountId,
		Message: "Customer create successfully",
	}, nil

}
