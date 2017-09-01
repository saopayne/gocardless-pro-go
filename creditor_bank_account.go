package main

import (
	"fmt"
	"net/url"
)

type CreditorBankAccountService service

// CreditorBankAccountList is a list object for creditbank accounts.
type CreditorBankAccountList struct {
	Meta   ListMeta
	Values []CreditorBankAccount `json:"data"`
}

type CreditorBankAccountCreateRequest struct {
	AccountNumber             string            `json:"account_number,omitempty"`
	BankCode                  string            `json:"bank_code,omitempty"`
	BranchCode                string            `json:"branch_code,omitempty"`
	CountryCode               string            `json:"country_code,omitempty"`
	Iban                      string            `json:"iban,omitempty"`
	AccountHolderName         string            `json:"account_holder_name,omitempty"`
	Currency                  string            `json:"currency,omitempty"`
	Links                     []CreditorLink    `json:"links,omitempty"`
	Metadata                  map[string]string `json:"metadata,omitempty"`
	SetAsDefaultPayoutAccount bool              `json:"set_as_default_payout_account,omitempty"`
}

type CreditorBankAccountDisableRequest struct {
	Identity string `json:"identity,omitempty"`
}

type CreditorBankAccountListRequest struct {
	CreatedAt CreatedAt `json:"created_at,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	Before    string    `json:"before,omitempty"`
	After     string    `json:"after,omitempty"`
	Creditor  string    `json:"creditor,omitempty"`
	Enabled   string    `json:"enabled,omitempty"`
}

type CreditorBankAccount struct {
	Id                  string            `json:"id,omitempty"`
	BankName            string            `json:"bank_name,omitempty"`
	CountryCode         string            `json:"country_code,omitempty"`
	CreatedAt           string            `json:"created_at,omitempty"`
	Currency            string            `json:"currency,omitempty"`
	AccountHolderName   string            `json:"account_holder_name,omitempty"`
	AccountNumberEnding string            `json:"account_number_ending,omitempty"`
	Enabled             bool              `json:"enabled,omitempty"`
	Links               []CreditorLink    `json:"links,omitempty"`
	Metadata            map[string]string `json:"metadata,omitempty"`
	ResponseUrl         string            `json:"responseurl,omitempty"`
}

// Create creates a new credit bank account
func (s *CreditorBankAccountService) CreateCreditorBankAccount(bankAccount *CreditorBankAccountCreateRequest) (*CreditorBankAccount, error) {
	u := fmt.Sprintf("/creditor_bank_accounts")
	account := &CreditorBankAccount{}
	err := s.client.Call("POST", u, bankAccount, account)

	return account, err
}

// List returns a list of credit bank accounts.
// https://developer.gocardless.com/api-reference/#creditor-bank-accounts-list-creditor-bank-accounts
func (s *CreditorBankAccountService) ListCreditorBankAccount(req *CreditorBankAccountListRequest) (*CreditorBankAccountList, error) {
	return s.ListNCreditorBankAccount(10, 0, req)
}

// ListN Returns a cursor-paginated list of your creditor bank accounts.
// https://developer.gocardless.com/api-reference/#creditor-bank-accounts-list-creditor-bank-accounts
func (s *CreditorBankAccountService) ListNCreditorBankAccount(count, offset int, req *CreditorBankAccountListRequest) (*CreditorBankAccountList, error) {
	params := url.Values{}
	params.Add("after", req.After)
	params.Add("before", req.Before)
	params.Add("created_at[gt]", req.CreatedAt.Gt)
	params.Add("created_at[gte]", req.CreatedAt.Gte)
	params.Add("created_at[lt]", req.CreatedAt.Lt)
	params.Add("created_at[lte]", req.CreatedAt.Lte)
	params.Add("limit", string(req.Limit))
	u := paginateURL("/creditor_bank_accounts", count, offset)
	bankAccounts := &CreditorBankAccountList{}
	err := s.client.Call("GET", u, params, bankAccounts)

	return bankAccounts, err
}

// Retrieves the details of an existing creditor bank account.
// https://developer.gocardless.com/api-reference/#creditor-bank-accounts-get-a-single-creditor-bank-account
func (s *CreditorBankAccountService) GetCreditorBankAccount(id string) (*CreditorBankAccount, error) {
	u := fmt.Sprintf("/creditor_bank_accounts/%s", id)
	txn := &CreditorBankAccount{}
	err := s.client.Call("GET", u, nil, txn)

	return txn, err
}

// Immediately disables the bank account, no money can be paid out to a disabled account.
// https://developer.gocardless.com/api-reference/#creditor-bank-accounts-disable-a-creditor-bank-account
func (s *CreditorBankAccountService) DisableCreditorBankAccount(bankAccount *CreditorBankAccountDisableRequest) (*Response, error) {
	u := fmt.Sprintf("/creditor_bank_accounts/%s/actions/disable", bankAccount.Identity)
	resp := &Response{}
	err := s.client.Call("POST", u, bankAccount, resp)

	return resp, err
}
