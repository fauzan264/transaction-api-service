package user

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user User) (User, error)
	CheckNIK(nik string) bool
	CheckNoHP(phone_number string) bool
	GetBalance(numberBalance string) (UserBalance, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) CheckNIK(nik string) bool {
	err := r.db.Model(&User{}).Where("nik = ?", nik).Error
	if err != nil {
		return false
	}

	return true
}

func (r *repository) CheckNoHP(phone_number string) bool {
	err := r.db.Model(&User{}).Where("phone_number = ?", phone_number).Error
	if err != nil {
		return false
	}

	return true
}

func (r *repository) GetBalance(numberBalance string) (UserBalance, error) {
	var userBalance UserBalance
	result := r.db.Where("number = ?", numberBalance).Find(&userBalance)

	if result.RowsAffected == 0 {
		return userBalance, errors.New("Number balance not found")
	}

	if result.Error != nil {
		return userBalance, result.Error
	}

	return userBalance, nil
}