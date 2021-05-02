package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dharmjit/paytab-transfer/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// TransferReq is struct type to hold request object recieved in transferfunds API
type TransferReq struct {
	FromAccount uuid.UUID `json:"from_account,omitempty"`
	ToAccount   uuid.UUID `json:"to_account,omitempty"`
	Amount      float64   `json:"amount,omitempty"`
}

func TransferFunds(s *service.AccountService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("TransferFundsHandler Called")
		var tr TransferReq
		err := json.NewDecoder(r.Body).Decode(&tr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if tr.FromAccount == uuid.Nil || tr.ToAccount == uuid.Nil || tr.Amount <= 0 {
			http.Error(w, "Request Validation Failed", http.StatusBadRequest)
			return
		}
		err = s.TransferFunds(tr.FromAccount, tr.ToAccount, tr.Amount)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func ListAccounts(s *service.AccountService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("ListAccountsHandler Called")
		errMessage := "Error in Transfer"
		accounts, err := s.ListAccounts()

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(accounts); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
		}
	})
}

func GetAccount(s *service.AccountService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GetAccountHandler Called")
		params := mux.Vars(r)
		accountID := params["accountID"]
		accountIDUUID, err := uuid.Parse(accountID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("string Parse to UUID failed"))
			return
		}
		account, err := s.GetAccount(accountIDUUID)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(account); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failed"))
		}
	})
}
