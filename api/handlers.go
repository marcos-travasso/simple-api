package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/marcos-travasso/simple-api/core/repository"
	"net/http"
)

func getBalanceHandler(w http.ResponseWriter, r *http.Request) {
	accountId := r.URL.Query().Get("account_id")

	account, err := GetService().GetAccount(accountId)
	if errors.Is(err, repository.ErrAccountNotFound) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%d", 0)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%d", account.Balance)
}

func postEventHandler(w http.ResponseWriter, r *http.Request) {
	var eventRequest EventRequest
	json.NewDecoder(r.Body).Decode(&eventRequest)

	response, err := handleEvent(eventRequest)
	if errors.Is(err, repository.ErrAccountNotFound) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%d", 0)
		return
	}

	responseJSON, _ := json.Marshal(response)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(responseJSON))
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	GetService().Reset()

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
