package input

type NewCustomerDto struct {
	BussinesUnitId   int    `json:"bussinesUnitId"`
	BussinesUnitName string `json:"bussinesUnitName"`
	IsoCurrency      string `json:"isoCurrency"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
}
