package controller

import (
	. "crm-app-go/dto/input"
	"crm-app-go/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type IFlywayController interface {
	RegisterLeadInCrm(w http.ResponseWriter, r *http.Request)
	CrateTransactionInCrm(w http.ResponseWriter, r *http.Request)
}

type ApiError struct {
    Field string
    Msg   string
}

type flywayController struct{}

var (
	flywayService service.IFlywayService
)

func NewFlywayController(service service.IFlywayService) IFlywayController {
	flywayService = service
	return &flywayController{}
}

func (flywayController *flywayController) RegisterLeadInCrm(w http.ResponseWriter, r *http.Request) {
	var newRegister RegisterLeadFlywayRequestBody
	var errorsMessage []string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newRegister); err != nil {
		errorsMessage = append(errorsMessage, fmt.Sprintf("%v", "Invalid request payload"))
		respondWithArrayError(w, http.StatusBadRequest, errorsMessage)
		return
	}
	defer r.Body.Close()
    instance := validator.New()
    validate:= instance.Struct(newRegister)
    if validate != nil {
        errors := validate.(validator.ValidationErrors)
		for _, validationError := range errors {
			errorsMessage = append(errorsMessage, fmt.Sprintf("field '%s' failed validation: '%s'", validationError.Field(), validationError.Tag()))
        }
        respondWithArrayError(w, http.StatusBadRequest, errorsMessage)
        return
    }
	res := flywayService.RegisterLeadInCrm(&newRegister)
	respondWithJSON(w, http.StatusOK, res)
}

func (flywayController *flywayController) CrateTransactionInCrm(w http.ResponseWriter, r *http.Request) {
	var newTransaction CrateTransactionFlyway
	var errorsMessage []string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newTransaction); err != nil {
		errorsMessage = append(errorsMessage, fmt.Sprintf("%v", "Invalid request payload"))
		respondWithArrayError(w, http.StatusBadRequest, errorsMessage)
		return
	}
	defer r.Body.Close()
	instance := validator.New()
    validate:= instance.Struct(newTransaction)
    if validate != nil {
        errors := validate.(validator.ValidationErrors)
		for _, validationError := range errors {
			errorsMessage = append(errorsMessage, fmt.Sprintf("field '%s' failed validation: '%s'", validationError.Field(), validationError.Tag()))
        }
        respondWithArrayError(w, http.StatusBadRequest, errorsMessage)
        return
    }
	res := flywayService.CrateTransactionInCrm(&newTransaction)
	respondWithJSON(w, http.StatusOK, res)
}
