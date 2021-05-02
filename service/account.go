package service

import (
	"errors"

	"github.com/dharmjit/paytab-transfer/domain"
	"github.com/dharmjit/paytab-transfer/repository"
	"github.com/google/uuid"
)

// AccountService provides service methods as in clean code architecture
// and inject repository interface as dependency
type AccountService struct {
	accountRepo repository.IAccountRepository
}

// NewNewAccountService intializes the AccountService
func NewAccountService(accountRepo repository.IAccountRepository) *AccountService {
	return &AccountService{accountRepo: accountRepo}
}

// TransferFunds service method is for transferring funds from one account to another
// business validation like account existence and Debit account balances are done here
func (as *AccountService) TransferFunds(fromAccount uuid.UUID, toAccount uuid.UUID, amount float64) error {
	debitAccount, err := as.accountRepo.GetAccount(fromAccount)
	if err != nil {
		return err
	}
	creditAccount, err := as.accountRepo.GetAccount(toAccount)
	if err != nil {
		return err
	}
	if debitAccount.Balance < amount {
		return errors.New("not sufficient balance")
	}
	debitAccount.Balance -= amount
	creditAccount.Balance += amount
	as.accountRepo.UpdateAccount(debitAccount)
	as.accountRepo.UpdateAccount(creditAccount)
	return nil
}

// ListAccounts provides the complte list of all account
func (as *AccountService) ListAccounts() ([]domain.Account, error) {
	return as.accountRepo.ListAccounts()
}

// GetAccount is to see details of an account
func (as *AccountService) GetAccount(accountID uuid.UUID) (domain.Account, error) {
	return as.accountRepo.GetAccount(accountID)
}
