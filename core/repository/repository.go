package repository

import (
	"errors"
	"github.com/marcos-travasso/simple-api/core/models"
)

var ErrAccountNotFound = errors.New("account not found")

type Repository interface {
	SaveAccount(account models.Account) error
	FindAccount(id string) (models.Account, error)
	DeleteAllAccounts()
}
