package service

import (
	. "crm-app-go/dto/input"
	"crm-app-go/repository"
	"encoding/json"
	"fmt"
	"strconv"
)

type IFlywayService interface{
	RegisterLeadInCrm(lead *RegisterLeadFlywayRequestBody) Any
	CrateTransactionInCrm(transaction *CrateTransactionFlyway) Any
}

type flywayService struct {}


func NewFlywayService(LogRepository repository.ILogRepository, ConfigurationRepository repository.IConfigurationRepository) IFlywayService{
	logRepository = LogRepository
	configurationRepository = ConfigurationRepository
	return &flywayService{}
}


func (flywayService *flywayService) RegisterLeadInCrm(lead *RegisterLeadFlywayRequestBody) Any {
	jsonBody, _ := json.Marshal(lead)
	fmt.Println(string(jsonBody))

	configuration, err := configurationRepository.GetConfigurationByBrandId(strconv.Itoa(lead.BussinesUnitId))
	fmt.Println(configuration.BuOwnerId)
	fmt.Println(err)

	return Any{
		Success: true, 
		Message: "", 
		Payload: "",
	}
}
	

func (flywayService *flywayService) CrateTransactionInCrm(transaction *CrateTransactionFlyway) Any  {
	return Any{
		Success: true, 
		Message: "", 
		Payload: "",
	}
}
