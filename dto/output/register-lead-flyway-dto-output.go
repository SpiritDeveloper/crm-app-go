package output


type RegisterUserError400Model struct {
	*BaseModel
	Payload struct {
		// Extensions:
    	// x-order: "0"
		Error []*string `json:"error"`
	}  `json:"payload"`
}


type RegisterUserError404Model struct {
	*BaseModel
	Payload struct {
		// Extensions:
    	// x-order: "0"
		Error []*string `json:"error"`
	}  `json:"payload"`
}

type RegisterUserError500Model struct {
	*BaseModel
	Payload struct {
		// Extensions:
    	// x-order: "0"
		Error []*string `json:"error"`
	}  `json:"payload"`
}

type RegisterUserSuccessModel struct {
    *BaseModel
	Payload struct {
		// Extensions:
    	// x-order: "0"
		Code    int `json:"code"`
		// Extensions:
    	// x-order: "1"
		Url string `json:"url"`
		// Extensions:
    	// x-order: "2"
		AccountId string `json:"accountId"`
	} `json:"payload"`
}
