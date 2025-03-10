package transaction

import (
	"errors"
	"time"

	"github.com/fauzan264/transaction-api-service/user"
	"github.com/google/uuid"
)

type Service interface {
	WithdrawTransaction(input TransactionInput) (Transaction, error)
	SavingTransaction(input TransactionInput) (Transaction, error)
}

type service struct {
	repository Repository
	userRepository user.Repository
}

func NewService(repository Repository, userRepository user.Repository) *service {
	return &service{repository, userRepository}
}

func (s *service) WithdrawTransaction(input TransactionInput) (Transaction, error) {
	getBalance, err := s.userRepository.GetBalance(input.NumberBalance)
	if err != nil {
		return Transaction{}, errors.New("Number balance not found")
	}

	if getBalance.Balance < input.Amount {
		return Transaction{}, errors.New("Your balance is insufficient for this transaction")
	}

	getBalance.Balance = getBalance.Balance - input.Amount
	getBalance.UpdatedAt = time.Now()
	_, err = s.userRepository.UpdateBalance(getBalance)
	if err != nil {
		return Transaction{}, err
	}

	transaction := Transaction{
		ID: uuid.New(),
		UserBalanceID: getBalance.ID,
		Status: "success",
		Amount: int(input.Amount),
		CreatedAt: time.Now(),
		UserBalance: getBalance,
	}

	createTransaction, err := s.repository.CreateTransaction(transaction)
	if err != nil {
		return createTransaction, err
	}
	
	return createTransaction, nil
}

func (s *service) SavingTransaction(input TransactionInput) (Transaction, error) {
	getBalance, err := s.userRepository.GetBalance(input.NumberBalance)
	if err != nil {
		return Transaction{}, errors.New("Number balance not found")
	}

	if input.Amount <= 0 {
		return Transaction{}, errors.New("Saving amount cannot be less than 0")
	}

	getBalance.Balance = getBalance.Balance + input.Amount
	getBalance.UpdatedAt = time.Now()
	_, err = s.userRepository.UpdateBalance(getBalance)
	if err != nil {
		return Transaction{}, err
	}

	transaction := Transaction{
		ID: uuid.New(),
		UserBalanceID: getBalance.ID,
		Status: "success",
		Amount: int(input.Amount),
		CreatedAt: time.Now(),
		UserBalance: getBalance,
	}

	createTransaction, err := s.repository.CreateTransaction(transaction)
	if err != nil {
		return createTransaction, err
	}

	return createTransaction, nil
}