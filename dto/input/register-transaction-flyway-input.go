package input

type CrateTransactionFlyway struct {
	Email string    `validate:"gt=0,omitempty,notblank" json:"email"`
	Amount int    `validate:"gt=0,omitempty,notblank" json:"amount"`
	PspName int  `validate:"gt=0,omitempty,notblank" json:"pspName"`
	IdTransaction int  `validate:"gt=0,omitempty,notblank" json:"idTransaction"`
}
