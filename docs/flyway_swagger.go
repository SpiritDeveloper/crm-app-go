package docs

import (
	. "crm-app-go/dto/input"
	. "crm-app-go/dto/output"
)

// swagger:route POST /api/v1/flyway/register-user Flyway postRegisterLead
// Register new user in crm flyway
// responses:
//   200: registerLeadResponseStatus200
//   400: registerLeadResponseStatus400
//   404: registerLeadResponseStatus404
//   500: registerLeadResponseStatus500

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

// Flyway Response With Status Code 400
// swagger:response registerLeadResponseStatus400
type Status400RegisterLeadResponseWrapper struct {
	// Register new lead body request.
	// in:body
	Response *RegisterUserError400Model `json:"payload"`
}

// Flyway Response With Status Code 404
// swagger:response registerLeadResponseStatus404
type Status404RegisterLeadResponseWrapper struct {
	// Register new lead body request.
	// in:body
	Response *RegisterUserError404Model `json:"payload"`
}

// Flyway Response With Status Code 500
// swagger:response registerLeadResponseStatus500
type Status500RegisterLeadResponseWrapper struct {
	// Register new lead body request.
	// in:body
	Response *RegisterUserError500Model `json:"payload"`
}



// swagger:route POST /api/v1/flyway/register-transactions Flyway postRegisterTransacgtions
// Register new user in crm flyway
// responses:
//   200: registerTransactionResponseStatus200
//   400: registerTransactionResponseStatus400
//   404: registerTransactionResponseStatus404
//   500: registerTransactionResponseStatus500

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

// Flyway Response With Status Code 400
// swagger:response registerTransactionResponseStatus400
type Status400RegisterTransactionResponseWrapper struct {
	// Response error when a controlled error exists.
	// in:body
	Response *RegisterTransactionError400Model `json:"payload"`
}

// Flyway Response With Status Code 404
// swagger:response registerTransactionResponseStatus404
type Status404RegisterTransactionResponseWrapper struct {
	// Response when not found data.
	// in:body
	Response *RegisterTransactionError404Model `json:"payload"`
}

// Flyway Response With Status Code 500
// swagger:response registerTransactionResponseStatus500
type Status500RegisterTransactionResponseWrapper struct {
	// Response when exist have internal error in server.
	// in:body
	Response *RegisterTransactionError500Model `json:"payload"`
}
