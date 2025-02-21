package transaction

type TransactionInput struct {
	NumberBalance		string 	`json:"number_balance" binding:"required"`
	Amount				float64 `json:"amount" binding:"required"`
}