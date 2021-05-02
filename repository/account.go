package repository

import (
	"fmt"
	"log"

	"github.com/dharmjit/paytab-transfer/domain"
	"github.com/google/uuid"
)

// Repository Interface provides the set of basic methods against our domain object
type IAccountRepository interface {
	UpdateAccount(account domain.Account) error
	ListAccounts() ([]domain.Account, error)
	GetAccount(accountID uuid.UUID) (domain.Account, error)
}

// AccountRepository is an in memory Implementation of IAccountRepository type
type AccountRepository struct {
	accountMap map[uuid.UUID]domain.Account
}

func NewAccountRepository(accountMap map[uuid.UUID]domain.Account) *AccountRepository {
	return &AccountRepository{accountMap: accountMap}
}

func (ar *AccountRepository) UpdateAccount(account domain.Account) error {
	log.Printf("UpdateAccountRepo Called")
	ar.accountMap[account.ID] = account
	return nil
}

func (ar *AccountRepository) ListAccounts() ([]domain.Account, error) {
	log.Printf("ListAccountRepo Called")
	var accounts []domain.Account
	for _, acc := range ar.accountMap {
		accounts = append(accounts, acc)
	}
	return accounts, nil
}

func (ar *AccountRepository) GetAccount(accountID uuid.UUID) (domain.Account, error) {
	account, ok := ar.accountMap[accountID]
	if ok {
		return account, nil
	}
	return domain.Account{}, fmt.Errorf("account does not exist %v", accountID)
}
