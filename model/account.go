package model

type GetAccountRequest struct {
	Code  int    `json:"Code" validate:"required"`
	Name  string `json:"Name" validate:"required"`
	Type  int    `json:"Type" validate:"required"`
	Store int    `json:"Store"`
}
type AccountTransactionReq struct {
	TransactionType int     `json:"transactionType"`
	Safe            int     `json:"safe"`
	AccSerial       int     `json:"account"`
	Amount          float64 `json:"amount"`
	AccType         int     `json:"accountType"`
	Store           int     `json:"store"`
}

type Account struct {
	Serial       int
	AccountCode  int
	AccountName  string
	Trn          string
	RaseedBefore float64
}
