package docs

import (
	. "crm-app-go/dto/input"
	"crm-app-go/model"
)

// swagger:route POST /api/v1/flyway/register-user Flyway postRegisterLead
// Register new user in crm flyway
// responses:
//   200: registerLeadResponse
//   400: error400
//   404: error404
//   500: error500

// Flyway Response Register new lead in crm flyway
// swagger:response registerLeadResponse
type RegisterLeadResponseWrapper struct {
	// in:body
	Success bool `json:"success"`
	Message string `json:"message"`
	Payload model.LeadFlayway `json:"payload"`
}

// swagger:parameters postRegisterLead
type RegisterLeadBodyWrapper struct {
	// Register new lead body request.
	// in:body
	Body RegisterLeadFlywayRequestBody
}

// swagger:route POST /api/v1/flyway/register-transactions Flyway postRegisterTransactionParam
// Register transaction in crm flyway
// responses:
//   200: createTransactionResponse
//   400: error400
//   404: error404
//   500: error500

// Flyway Response Register new transaction
// swagger:response createTransactionResponse
type RegisterTransactionResponseWrapper struct {
	// in:body
	Success bool `json:"success"`
	Message string `json:"message"`
	Payload model.LeadFlayway `json:"payload"`
}

// swagger:parameters postRegisterTransactionParam
type RegisterTransactionBodyWrapper struct {
	// Transaction Request Body.
	// in:body
	Body CrateTransactionFlyway
}


// swagger:route GET /api/v1/wallet/{ID}/transaction Flyway getTransactionParam
// Fetch transaction associated with given wallet id
// responses:
//   200: getTransactionResponse
//   400: error400
//   404: error404
//   500: error500

// swagger:parameters getTransactionParam
type GetTransactionParamsWrapper struct {
	// Wallet ID
	// In: path
	ID string `json:"id"`
}
