package output


type RegisterTransactionError400Model struct {
	*BaseModel
	Payload struct {
		// Extensions:
    	// x-order: "0"
		Error []*string `json:"error"`
	} `json:"payload"`
}


type RegisterTransactionError404Model struct {
	*BaseModel
	Payload struct {
		// Extensions:
    	// x-order: "0"
		Error []*string `json:"error"`
	} `json:"payload"`
}

type RegisterTransactionError500Model struct {
	*BaseModel
	Payload struct {
		// Extensions:
    	// x-order: "0"
		Error []*string `json:"error"`
	} `json:"payload"`
}

type RegisterTransactionSuccessModel struct {
    *BaseModel
	Payload struct {
		// Extensions:
    	// x-order: "0"
		Code    int `json:"code"`
		// Extensions:
    	// x-order: "1"
		Message string `json:"message"`
		// Extensions:
    	// x-order: "2"
		Data struct{
			// Extensions:
			// x-order: "0"
			Code    int `json:"code"`
			// Extensions:
			// x-order: "1"
			Message string `json:"message"`
			// Extensions:
			// x-order: "2"
		} `json:"data"`
	} `json:"payload"`
}
