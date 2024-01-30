package output

type BaseModel struct {
    // Extensions:
    // x-order: "1"
	Success  bool `json:"success"`
	// Extensions:
	// x-order: "2"
	Message   string `json:"message"`
	// Extensions:
	// x-order: "0"
	Code   int `json:"code"`
}
