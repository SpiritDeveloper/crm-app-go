package docs

// Bad Request error Response | Validation error message or invalid json
// swagger:response error400
type Error400 struct {
	// in:body
	Success  bool `json:"success"`
	Message   string `json:"message"`
	Code   int `json:"code"`
}

// Internal server error Response | server is down or db constraint errors
// swagger:response error500
type Error500 struct {
	// in:body
	Success  bool `json:"success"`
	Message   string `json:"message"`
	Code   int `json:"code"`
	Payload struct{
		Error []string `json:"error"`
	} `json:"payload"`
}

// Not Found error Response | Page not found or record not found
// swagger:response error404
type Error404 struct {
	// in:body
	Success  bool `json:"success"`
	Message   string `json:"message"`
	Code   int `json:"code"`
	Payload struct{
		Error []string `json:"error"`
	} `json:"payload"`
}
