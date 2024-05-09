package service

import (
	"github.com/marcos-travasso/simple-api/core/models"
	"github.com/marcos-travasso/simple-api/core/repository"
	"github.com/marcos-travasso/simple-api/core/repository/in_memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_GetAccount(t *testing.T) {
	t.Run("should return error when account was not found", func(t *testing.T) {
		repo := in_memory.NewRepository()
		s := NewService(repo)

		_, err := s.GetAccount("1234")

		assert.ErrorIs(t, err, repository.ErrAccountNotFound)
	})

	t.Run("should return stored account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		s := NewService(repo)

		storedAccount := models.Account{
			ID:      "123",
			Balance: 0,
		}
		err := repo.SaveAccount(storedAccount)
		require.NoError(t, err)

		retrievedAccount, err := s.GetAccount(storedAccount.ID)

		assert.Equal(t, storedAccount, retrievedAccount)
		assert.NoError(t, err)
	})
}

func TestService_CreateAccount(t *testing.T) {
	t.Run("should create an account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		s := NewService(repo)

		account := models.Account{
			ID:      "100",
			Balance: 10,
		}
		createdAccount, err := s.CreateAccount(account)
		assert.NoError(t, err)
		assert.Equal(t, account, createdAccount)
	})
}

func TestService_Deposit(t *testing.T) {
	t.Run("should deposit", func(t *testing.T) {
		repo := in_memory.NewRepository()
		s := NewService(repo)

		account := models.Account{
			ID:      "100",
			Balance: 10,
		}
		_, err := s.CreateAccount(account)
		require.NoError(t, err)

		updatedAccount, err := s.Deposit(account.ID, 10)
		assert.NoError(t, err)
		assert.Equal(t, 20, updatedAccount.Balance)
	})

	t.Run("should deposit in non existing account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		s := NewService(repo)

		_, err := s.Deposit("100", 10)
		assert.NoError(t, err)

		storedAccount, err := s.GetAccount("100")
		require.NoError(t, err)
		assert.Equal(t, 10, storedAccount.Balance)
	})
}

func TestService_Withdraw(t *testing.T) {
	t.Run("should withdraw", func(t *testing.T) {
		repo := in_memory.NewRepository()
		s := NewService(repo)

		account := models.Account{
			ID:      "100",
			Balance: 20,
		}
		_, err := s.CreateAccount(account)
		require.NoError(t, err)

		updatedAccount, err := s.Withdraw(account.ID, 5)
		assert.NoError(t, err)
		assert.Equal(t, 15, updatedAccount.Balance)
	})

	t.Run("should not withdraw from non existing account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		s := NewService(repo)

		_, err := s.Withdraw("200", 10)
		assert.ErrorIs(t, err, repository.ErrAccountNotFound)
	})
}

func TestService_Transfer(t *testing.T) {
	t.Run("should transfer between accounts", func(t *testing.T) {
		repo := in_memory.NewRepository()
		s := NewService(repo)

		origin := models.Account{ID: "100", Balance: 15}
		dest := models.Account{ID: "300", Balance: 0}
		origin, err := s.CreateAccount(origin)
		require.NoError(t, err)
		dest, err = s.CreateAccount(dest)
		require.NoError(t, err)

		transaction, err := s.Transfer(origin.ID, dest.ID, 15)
		assert.NoError(t, err)
		assert.Equal(t, 0, transaction.Origin.Balance)
		assert.Equal(t, 15, transaction.Destination.Balance)
	})

	t.Run("should transfer to not existing account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		s := NewService(repo)

		origin := models.Account{ID: "100", Balance: 15}
		origin, err := s.CreateAccount(origin)
		require.NoError(t, err)

		transaction, err := s.Transfer(origin.ID, "300", 15)
		assert.NoError(t, err)
		assert.Equal(t, 0, transaction.Origin.Balance)
		assert.Equal(t, 15, transaction.Destination.Balance)
	})

	t.Run("should not transfer from non existing account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		s := NewService(repo)

		_, err := s.Transfer("200", "300", 15)
		assert.ErrorIs(t, err, repository.ErrAccountNotFound)
	})
}

func TestService_Reset(t *testing.T) {
	t.Run("should clean accounts", func(t *testing.T) {
		repo := in_memory.NewRepository()
		s := NewService(repo)

		account := models.Account{
			ID:      "100",
			Balance: 0,
		}
		_, err := s.CreateAccount(account)
		require.NoError(t, err)

		s.Reset()

		_, err = s.GetAccount(account.ID)
		assert.ErrorIs(t, err, repository.ErrAccountNotFound)
	})
}
