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
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newRegister); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Create a new validator instance
    instance := validator.New()

    // Validate the User struct
    validate:= instance.Struct(newRegister)
    if validate != nil {
        // Validation failed, handle the error
        errors := validate.(validator.ValidationErrors)
		for _, validationError := range errors {
            fmt.Println(validationError.Field(), validationError.Error())
        }
        respondWithError(w, http.StatusBadRequest, "")
        return
    }

	res := flywayService.RegisterLeadInCrm(&newRegister)
	respondWithJSON(w, http.StatusOK, res)
}

func (flywayController *flywayController) CrateTransactionInCrm(w http.ResponseWriter, r *http.Request) {
	var newTransaction CrateTransactionFlyway
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newTransaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	res := flywayService.CrateTransactionInCrm(&newTransaction)
	respondWithJSON(w, http.StatusOK, res)
}
