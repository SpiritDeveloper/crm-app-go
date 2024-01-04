package input

type NewCustomerDto struct {
	BussinesUnitId   string `json:"bussinesUnitId" validate:"required"`
	BussinesUnitName string `json:"bussinesUnitName" validate:"required"`
	IsoCurrency      string `json:"isoCurrency" validate:"required"`
	FirstName        string `json:"firstName" validate:"required"`
	LastName         string `json:"lastName" validate:"required"`
	Email            string `json:"email" validate:"required"`
}
