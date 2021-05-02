package domain

import (
	"github.com/google/uuid"
)

// Account is the domain object for this service
type Account struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Balance float64   `json:"balance,string,omitempty"`
}
