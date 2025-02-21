package user

type UserCreateFormatter struct {
	NumberBalance	string	`json:"nomor_rekening"`
}

type UserBalanceFormatter struct {
	Balance			float64 `json:"balance"`
}

func FormatUserCreate(user User) UserCreateFormatter {
	formatter := UserCreateFormatter{
		NumberBalance: user.UserBalance.Number,
	}

	return formatter
}

func FormatUserBalance(userBalance UserBalance) UserBalanceFormatter {
	formatter := UserBalanceFormatter{
		Balance: userBalance.Balance,
	}

	return formatter
}