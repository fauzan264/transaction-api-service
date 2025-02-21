package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID			uuid.UUID	`gorm:"type:char(36);primary_key"`
	Name		string		`gorm:"type:varchar(100)"`
	NIK			string		`gorm:"type:varchar(16)"`
	NoHp		string		`gorm:"type:varchar(13)"`
	CreatedAt	time.Time	`gorm:"type:timestamp"`
	UpdatedAt	time.Time	`gorm:"type:timestamp"`
	
	UserBalance	UserBalance `gorm:"foreignKey:UserID"`
}

type UserBalance struct {
	ID			uuid.UUID	`gorm:"type:char(36);primary_key"`
	UserID		uuid.UUID	`gorm:"type:char(36)"`
	Number		string		`gorm:"type:varchar(20)"`
	Balance		float64		`gorm:"type:decimal(12,2);default:0"`
	UpdatedAt	time.Time	`gorm:"type:timestamp"`
}

func (User) TableName() string {
	return "users.users"
}

func (UserBalance) TableName() string {
	return "users.balances"
}