package in_memory

import (
	"github.com/marcos-travasso/simple-api/core/models"
	"github.com/marcos-travasso/simple-api/core/repository"
)

type Repository struct {
	accounts map[string]models.Account
}

func NewRepository() *Repository {
	return &Repository{
		accounts: make(map[string]models.Account),
	}
}

func (r *Repository) SaveAccount(account models.Account) error {
	r.accounts[account.ID] = account
	return nil
}

func (r *Repository) FindAccount(id string) (models.Account, error) {
	account, exists := r.accounts[id]
	if !exists {
		return models.Account{}, repository.ErrAccountNotFound
	}

	return account, nil
}

func (r *Repository) DeleteAllAccounts() {
	r.accounts = make(map[string]models.Account)
}
