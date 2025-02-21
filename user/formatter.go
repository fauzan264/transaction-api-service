package user

type UserCreateFormatter struct {
	NumberBalance	string	`json:"nomor_rekening"`
}

func FormatUserCreate(user User) UserCreateFormatter {
	formatter := UserCreateFormatter{
		NumberBalance: user.UserBalance.Number,
	}

	return formatter
}