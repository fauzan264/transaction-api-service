package user

import (
	"errors"
	"time"

	"github.com/fauzan264/transaction-api-service/helper"
	"github.com/google/uuid"
)

type Service interface {
	RegisterUser(input RegisterUserinput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return & service{repository}
}

func (s *service) RegisterUser(input RegisterUserinput) (User, error) {
	if input.Name == "" || input.NIK == "" || input.NoHP == "" {
		return User{}, errors.New("name, nik, and no_hp cannot be empty")
	}

	if s.repository.CheckNIK(input.NIK) {
		return User{}, errors.New("The provided NIK is already in use. Please use a different NIK.")
	}

	if s.repository.CheckNoHP(input.NoHP) {
		return User{}, errors.New("The provided Phone Number is already in use. Please use a different Phone Number.")
	}
	
	user_id := uuid.New()
	user := User{
		ID: user_id,
		Name: input.Name,
		NIK: input.NIK,
		NoHp: input.NoHP,
		CreatedAt: time.Now(),
	}

	accountNumber := helper.GenerateAccountNumber()
	userBalance := UserBalance{
		ID: uuid.New(),
		UserID: user_id,
		Number: accountNumber,
		Balance: 0,
		UpdatedAt: time.Now(),
	}

	user.UserBalance = userBalance

	newUser, err := s.repository.CreateUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}