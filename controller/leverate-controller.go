package controller

import (
	. "crm-app-go/dto/input"
	"crm-app-go/service"
	"encoding/json"
	"net/http"
)

type ILeverateController interface {
	SendLeadToCrm(w http.ResponseWriter, r *http.Request)
}

type leverateController struct{}

var (
	leverateService service.ILeverateService
)

func NewLeverateController(service service.ILeverateService) ILeverateController {
	leverateService = service
	return &leverateController{}
}

func (leverateController *leverateController) SendLeadToCrm(w http.ResponseWriter, r *http.Request) {

	var newCustomer NewCustomerDto
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newCustomer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	println("asdasd")
	res, err := leverateService.SendLeadToCrm(&newCustomer)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, res)
}
