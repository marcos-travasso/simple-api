package api

import (
	"encoding/json"
	"github.com/marcos-travasso/simple-api/core/models"
	"github.com/marcos-travasso/simple-api/core/repository/in_memory"
	core "github.com/marcos-travasso/simple-api/core/service"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_GetBalanceHandler(t *testing.T) {
	t.Run("should return not found for non existing account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		InjectService(core.NewService(repo))

		req := httptest.NewRequest("GET", "/balance?account_id=1234", nil)
		w := httptest.NewRecorder()

		getBalanceHandler(w, req)

		assertBody(t, w, http.StatusNotFound, "0")
	})

	t.Run("should return balance for existing account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		InjectService(core.NewService(repo))

		req := httptest.NewRequest("GET", "/balance?account_id=100", nil)
		w := httptest.NewRecorder()

		_, err := GetService().CreateAccount(models.Account{ID: "100", Balance: 20})
		require.NoError(t, err)

		getBalanceHandler(w, req)

		assertBody(t, w, http.StatusOK, "20")
	})
}

func Test_PostEventHandler(t *testing.T) {
	t.Run("should deposit balance for non existing account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		InjectService(core.NewService(repo))

		payload := EventRequest{
			Type:        "deposit",
			Destination: "100",
			Amount:      10,
		}
		payloadJSON, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/event", strings.NewReader(string(payloadJSON)))
		w := httptest.NewRecorder()

		postEventHandler(w, req)

		assertBody(t, w, http.StatusCreated, "{\"destination\":{\"id\":\"100\",\"balance\":10}}")
	})

	t.Run("should deposit balance into existing account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		InjectService(core.NewService(repo))

		_, err := GetService().CreateAccount(models.Account{ID: "100", Balance: 10})
		require.NoError(t, err)

		payload := EventRequest{
			Type:        "deposit",
			Destination: "100",
			Amount:      10,
		}
		payloadJSON, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/event", strings.NewReader(string(payloadJSON)))
		w := httptest.NewRecorder()

		postEventHandler(w, req)

		assertBody(t, w, http.StatusCreated, "{\"destination\":{\"id\":\"100\",\"balance\":20}}")
	})

	t.Run("should not withdraw from non existing account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		InjectService(core.NewService(repo))

		payload := EventRequest{
			Type:        "withdraw",
			Destination: "200",
			Amount:      10,
		}
		payloadJSON, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/event", strings.NewReader(string(payloadJSON)))
		w := httptest.NewRecorder()

		postEventHandler(w, req)

		assertBody(t, w, http.StatusNotFound, "0")
	})

	t.Run("should withdraw from existing account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		InjectService(core.NewService(repo))

		_, err := GetService().CreateAccount(models.Account{ID: "100", Balance: 20})
		require.NoError(t, err)

		payload := EventRequest{
			Type:   "withdraw",
			Origin: "100",
			Amount: 5,
		}
		payloadJSON, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/event", strings.NewReader(string(payloadJSON)))
		w := httptest.NewRecorder()

		postEventHandler(w, req)

		assertBody(t, w, http.StatusCreated, "{\"origin\":{\"id\":\"100\",\"balance\":15}}")
	})

	t.Run("should transfer from existing account", func(t *testing.T) {
		repo := in_memory.NewRepository()
		InjectService(core.NewService(repo))

		_, err := GetService().CreateAccount(models.Account{ID: "100", Balance: 15})
		require.NoError(t, err)

		payload := EventRequest{
			Type:        "transfer",
			Origin:      "100",
			Destination: "300",
			Amount:      15,
		}
		payloadJSON, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/event", strings.NewReader(string(payloadJSON)))
		w := httptest.NewRecorder()

		postEventHandler(w, req)

		assertBody(t, w, http.StatusCreated, "{\"origin\":{\"id\":\"100\",\"balance\":0},\"destination\":{\"id\":\"300\",\"balance\":15}}")
	})
}

func Test_ResetHandler(t *testing.T) {
	t.Run("should reset state", func(t *testing.T) {
		repo := in_memory.NewRepository()
		InjectService(core.NewService(repo))

		req := httptest.NewRequest("POST", "/reset", nil)
		w := httptest.NewRecorder()

		resetHandler(w, req)

		assertBody(t, w, http.StatusOK, "OK")
	})
}
