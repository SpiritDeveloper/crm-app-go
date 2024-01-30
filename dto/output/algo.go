package output

//type ResponseNewCustomerDto struct {
//	Tpid    string `json:"tpid"`
//	Message string `json:"message"`
//}
//
//type ErrorNewCustomerDto struct {
//	Message string `json:"message"`
//	Status  bool   `json:"status"`
//}

type LeverateCreateCustomerResponse struct {
	AccountId         string `json:"accountId"`
	TpAccountName     string `json:"tpAccountName"`
	TpAccountPassword string `json:"tpAccountPassword"`
}

type LeverateGenerateTokenResponse struct {
	Token     string `json:"Token"`
	ExpiresIn int    `json:"ExpiresIn"`
}