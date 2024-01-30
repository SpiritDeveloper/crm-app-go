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
		Message string `json:"message"`
	} `json:"payload"`
}
