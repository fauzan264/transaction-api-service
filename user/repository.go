package user

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user User) (User, error)
	CheckNIK(nik string) bool
	CheckPhoneNumber(phone_number string) bool
	GetBalance(numberBalance string) (UserBalance, error)
	UpdateBalance(userBalance UserBalance) (UserBalance, error)
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
	err := r.db.Where("nik = ?", nik).First(&User{}).Error
	if err != nil {
		return false
	}

	return true
}

func (r *repository) CheckPhoneNumber(phone_number string) bool {
	err := r.db.Where("phone_number = ?", phone_number).First(&User{}).Error
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

func (r *repository) UpdateBalance(userBalance UserBalance) (UserBalance, error) {
	err := r.db.Save(&userBalance).Error
	if err != nil {
		return userBalance, err
	}

	return userBalance, nil
}