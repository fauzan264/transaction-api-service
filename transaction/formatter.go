package transaction

type TransactionFormatter struct {
	Balance		float64	`json:"balance"`
}

func FormatTransactionBalance(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		Balance: transaction.UserBalance.Balance,
	}

	return formatter
}