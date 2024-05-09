package service

import (
	"errors"
	"github.com/marcos-travasso/simple-api/core/models"
	"github.com/marcos-travasso/simple-api/core/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetAccount(id string) (models.Account, error) {
	return s.repo.FindAccount(id)
}

func (s *Service) CreateAccount(account models.Account) (models.Account, error) {
	err := s.repo.SaveAccount(account)
	if err != nil {
		return models.Account{}, err
	}
	return account, err
}

func (s *Service) Deposit(id string, amount int) (account models.Account, err error) {
	account, err = s.GetAccount(id)
	if errors.Is(err, repository.ErrAccountNotFound) {
		return s.CreateAccount(models.Account{ID: id, Balance: amount})
	}
	if err != nil {
		return
	}

	account.Balance += amount
	err = s.repo.SaveAccount(account)
	return
}

func (s *Service) Withdraw(id string, amount int) (account models.Account, err error) {
	account, err = s.GetAccount(id)
	if err != nil {
		return
	}

	account.Balance -= amount
	err = s.repo.SaveAccount(account)
	return
}

func (s *Service) Transfer(originId, destinationId string, amount int) (transaction models.Transaction, err error) {
	transaction.Origin, err = s.Withdraw(originId, amount)
	if err != nil {
		return
	}

	transaction.Destination, err = s.Deposit(destinationId, amount)
	return
}

func (s *Service) Reset() {
	s.repo.DeleteAllAccounts()
}
