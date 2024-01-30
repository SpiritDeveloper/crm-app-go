package output

type ResponseNewCustomerDto struct {
	Tpid    string `json:"tpid"`
	Message string `json:"message"`
}

type ErrorNewCustomerDto struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}