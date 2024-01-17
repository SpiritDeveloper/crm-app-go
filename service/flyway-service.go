package service

import (
	. "crm-app-go/dto/input"
	. "crm-app-go/model"
	"crm-app-go/repository"
)

type IFlywayService interface{
	RegisterLeadInCrm(lead *RegisterLeadFlywayRequestBody) (*User, error)
	CrateTransactionInCrm(transaction *CrateTransactionFlyway) (*User, error)
}

type flywayService struct {}

var (
	flywayRepository repository.IFlywayRepository
)

type Item struct {}

func NewFlywayService(repository.IFlywayRepository) IFlywayService{
	return &flywayService{}
}


func (flywayService *flywayService) RegisterLeadInCrm(lead *RegisterLeadFlywayRequestBody) (*User, error) {
	item := User{}
	return &item, nil 
}


func (flywayService *flywayService) CrateTransactionInCrm(transaction *CrateTransactionFlyway) (*User, error) {
	item := User{}
	return &item, nil 
}
