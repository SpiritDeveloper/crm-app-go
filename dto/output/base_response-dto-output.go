package output

type BaseModel struct {
    // Extensions:
    // x-order: "0"
	Success  bool `json:"success"`
	// Extensions:
	// x-order: "1"
	Message   string `json:"message"`
}
