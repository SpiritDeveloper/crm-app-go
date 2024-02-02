package docs

import (
	. "crm-app-go/dto/input"
	. "crm-app-go/dto/output"
)

// swagger:route POST /api/v1/flyway/register-user Flyway postRegisterLead
// Register new user in crm flyway
// responses:
//   200: registerLeadResponseStatus200
//   400: error400
//   404: error404
//   500: error500

// swagger:parameters postRegisterLead
type RegisterLeadBodyWrapper struct {
	// Register new lead body request.
	// in:body
	Body RegisterLeadFlywayRequestBody
}

// Flyway Response With Status Code 200
// swagger:response registerLeadResponseStatus200
type Status200RegisterLeadResponseWrapper struct {
	// Register new lead body request.
	// in:body
	Response *RegisterUserSuccessModel `json:"payload"`
}

// swagger:route POST /api/v1/flyway/register-transactions Flyway postRegisterTransacgtions
// Register new user in crm flyway
// responses:
//   200: registerTransactionResponseStatus200
//   400: error400
//   404: error404
//   500: error500

// swagger:parameters postRegisterTransacgtions
type RegisterTransactionBodyWrapper struct {
	// Register new transaction in crmn flyway
	// in:body
	Body CrateTransactionFlyway
}

// Flyway Response With Status Code 200
// swagger:response registerTransactionResponseStatus200
type Status200RegisterTransactionResponseWrapper struct {
	// Register new lead body request.
	// in:body
	Response *RegisterTransactionSuccessModel `json:"payload"`
}
