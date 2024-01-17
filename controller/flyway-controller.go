package controller

import (
	. "crm-app-go/dto/input"
	"crm-app-go/service"
	"encoding/json"
	"log"
	"net/http"
)

type IFlywayController interface {
	registerLeadInCrm(w http.ResponseWriter, r *http.Request)
	createTransactionInCrm(w http.ResponseWriter, r *http.Request)
}

type flywayController struct{}

var (
	flywayService service.IFlywayService
)

func NewFlywayController(service service.IFlywayService) IFlywayController {
	flywayService = service
	return &flywayController{}
}

func (flywayController *flywayController) registerLeadInCrm(w http.ResponseWriter, r *http.Request) {
	var newRegister RegisterLeadFlywayRequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newRegister); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	res, err := flywayService.RegisterLeadInCrm(&newRegister)
	if err != nil {
		log.Printf("Not able to post User : %s", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, res)
}

func (flywayController *flywayController) createTransactionInCrm(w http.ResponseWriter, r *http.Request) {
	var newTransaction CrateTransactionFlyway
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newTransaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	res, err := flywayService.CrateTransactionInCrm(&newTransaction)
	if err != nil {
		log.Printf("Not able to post User : %s", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, res)
}
