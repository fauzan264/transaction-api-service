package transaction

import (
	"time"

	"github.com/fauzan264/transaction-api-service/user"
	"github.com/google/uuid"
)

type Transaction struct {
	ID				uuid.UUID 			`gorm:"type:char(36);primary_key"`
	UserBalanceID	uuid.UUID 			`gorm:"type:char(36)"`
	Status			string				`gorm:"type:varchar(20)"`
	Amount			int					`gorm:"type:int"`
	CreatedAt		time.Time			`gorm:"type:timestamp"`

	UserBalance		user.UserBalance	`gorm:"foreignKey:UserBalanceID;references:ID"`
}

func (Transaction) TableName() string {
	return "transactions.transactions"
}