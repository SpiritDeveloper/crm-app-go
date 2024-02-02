package output
type RegisterUserSuccessModel struct {
	Payload struct {
		// Extensions:
    	// x-order: "0"
		Url string `json:"url"`
		// Extensions:
    	// x-order: "1"
		AccountId string `json:"accountId"`
	} `json:"payload"`
}
