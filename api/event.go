package api

import "github.com/marcos-travasso/simple-api/core/models"

type EventRequest struct {
	Type        string `json:"type,omitempty"`
	Destination string `json:"destination,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Amount      int    `json:"amount,omitempty"`
}

type EventResponse struct {
	Origin      *models.Account `json:"origin,omitempty"`
	Destination *models.Account `json:"destination,omitempty"`
}

func handleEvent(e EventRequest) (EventResponse, error) {
	var r EventResponse
	switch e.Type {
	case "deposit":
		dest, err := GetService().Deposit(e.Destination, e.Amount)
		if err != nil {
			return r, err
		}
		r.Destination = &dest
	case "withdraw":
		origin, err := GetService().Withdraw(e.Origin, e.Amount)
		if err != nil {
			return r, err
		}
		r.Origin = &origin
	case "transfer":
		transaction, err := GetService().Transfer(e.Origin, e.Destination, e.Amount)
		if err != nil {
			return r, err
		}
		r.Origin = &transaction.Origin
		r.Destination = &transaction.Destination
	}

	return r, nil
}
