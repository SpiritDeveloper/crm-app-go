package input

type CrateTransactionFlyway struct {
	Email int    `json:"email" validate:"required"`
	Amount int    `json:"amount" validate:"required"`
	PspName int    `json:"pspName" validate:"required"`
	IdTransaction int    `json:"idTransaction" validate:"required"`
}
