package output
type RegisterTransactionSuccessModel struct {
	Payload struct {
		// Extensions:
    	// x-order: "0"
		Message string `json:"message"`
	} `json:"payload"`
}
