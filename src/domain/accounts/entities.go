package accounts

import "maisdindin/src/domain/budgets"

// AccountID is the account's UUID
type AccountID string

// AccountMethod describes types of Account
type AccountMethod int

// Valid account types
const (
	Checking AccountMethod = iota
	Savings
	Cash
	CreditCard
	LineOfCredit
	Assets
	Liability
)

// Account is used for abstracting user's sources of money
type Account struct {
	ID          AccountID        `json:"id"`
	BudgetID    budgets.BudgetID `json:"budget_id"`
	Method      AccountMethod    `json:"type"`
	Description string           `json:"description"`
	Balance     float64          `json:"balance"`
}

// NewAccount returns a Account reference
func NewAccount(id AccountID, budgetID budgets.BudgetID, method AccountMethod, description string, balance float64) *Account {
	return &Account{
		ID:          id,
		BudgetID:    budgetID,
		Method:      method,
		Description: description,
		Balance:     balance,
	}
}
