package docs

import (
	"crm-app-go/dto/input"
)

type FlywayRequestBody struct {
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
}

// swagger:route POST /api/v1/flyway/register-user-crm Flyway postUserParam
// Register new lead for flyway
// responses:
//   200: userResponse
//   400: error400
//   404: error404
//   500: error500

// Flyway Response
// swagger:response userResponse
type RegisterResponseWrapper struct {
	// in:body
	Body input.RegisterLeadFlyway
}

// swagger:parameters postUserParam
type RegisterParamsWrapper struct {
	// Flyway Request Body.
	// in:body
	Body FlywayRequestBody
}

