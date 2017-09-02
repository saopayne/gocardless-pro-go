package main

import (
	"fmt"
	"net/url"
	"encoding/json"
)

type CustomerBankAccountService service

type CustomerBankAccount struct {
	Id                  string            `json:"id,omitempty"`
	BankName            string            `json:"bank_name,omitempty"`
	CountryCode         string            `json:"country_code,omitempty"`
	CreatedAt           string            `json:"created_at,omitempty"`
	Currency            string            `json:"currency,omitempty"`
	AccountHolderName   string            `json:"account_holder_name,omitempty"`
	AccountNumberEnding string            `json:"account_number_ending,omitempty"`
	Enabled             bool              `json:"enabled,omitempty"`
	Links               []CustomerLink    `json:"links,omitempty"`
	Metadata            map[string]string `json:"metadata,omitempty"`
	ResponseUrl         string            `json:"responseurl,omitempty"`
}

type CustomerBankListRequest struct {
	CreatedAt CreatedAt `json:"created_at,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	Before    string    `json:"before,omitempty"`
	After     string    `json:"after,omitempty"`
}

// CustomerBankAccountList is a list object for customer bank accounts.
type CustomerBankAccountList struct {
	Meta   ListMeta
	Values []CustomerBankAccount `json:"data"`
}

type CustomerBankAccountCreateRequest struct {
	Iban              string            `json:"iban,omitempty"`
	BankCode          string            `json:"bank_code,omitempty"`
	BranchCode        string            `json:"branch_code,omitempty"`
	CountryCode       string            `json:"country_code,omitempty"`
	Currency          string            `json:"currency,omitempty"`
	AccountHolderName string            `json:"account_holder_name,omitempty"`
	AccountNumber     string            `json:"account_number,omitempty"`
	Links             []CustomerLink    `json:"links,omitempty"`
	Metadata          map[string]string `json:"metadata,omitempty"`
	ResponseUrl       string            `json:"responseurl,omitempty"`
}

type CustomerBankAccountDisableRequest struct {
	Identity string `json:"identity,omitempty"`
}

// Create creates a new customer bank account
func (s *CustomerBankAccountService) CreateCustomerBankAccount(bankAccount *CustomerBankAccountCreateRequest) (*CustomerBankAccount, error) {
		// do something with error
	fmt.Sprintf("The error while creating a customer bank account is :%s", "About to create")

	u := fmt.Sprintf("/customer_bank_accounts")
	account := &CustomerBankAccount{}
	rel := map[string]interface{}{
		"customer_bank_accounts": bankAccount,
	}
	custJson, _ := json.Marshal(rel)
	customerObject := string(custJson[:])
	fmt.Sprintf("Making the request with params %s", customerObject)
	if rel == nil {
		fmt.Println("Rubbish call")
	}
	err := s.client.Call("POST", u, rel, account)
	if err != nil {
		fmt.Println("Successful call")
	}
	return account, err
}

// List returns a list of customer bank accounts.
// https://developer.gocardless.com/api-reference/#customer-bank-accounts-list-customer-bank-accounts
func (s *CustomerBankAccountService) ListCustomerBankAccount(req *CustomerBankListRequest) (*CustomerBankAccountList, error) {
	return s.ListNCustomerBankAccount(10, 0, req)
}

// ListN Returns a cursor-paginated list of your customer bank accounts.
// https://developer.gocardless.com/api-reference/#customer-bank-accounts-list-customer-bank-accounts
func (s *CustomerBankAccountService) ListNCustomerBankAccount(count, offset int, req *CustomerBankListRequest) (*CustomerBankAccountList, error) {
	params := url.Values{}
	params.Add("after", req.After)
	params.Add("before", req.Before)
	params.Add("created_at[gt]", req.CreatedAt.Gt)
	params.Add("created_at[gte]", req.CreatedAt.Gte)
	params.Add("created_at[lt]", req.CreatedAt.Lt)
	params.Add("created_at[lte]", req.CreatedAt.Lte)
	params.Add("limit", string(req.Limit))
	u := paginateURL("/customer_bank_accounts", count, offset)
	bankAccounts := &CustomerBankAccountList{}
	err := s.client.Call("GET", u, params, bankAccounts)

	return bankAccounts, err
}

// Retrieves the details of an existing customer bank account.
// https://developer.gocardless.com/api-reference/#customer-bank-accounts-get-a-single-customer-bank-account
func (s *CustomerBankAccountService) GetCustomerBankAccount(id string) (*CustomerBankAccount, error) {
	u := fmt.Sprintf("/customer_bank_accounts/%s", id)
	account := &CustomerBankAccount{}
	err := s.client.Call("GET", u, nil, account)

	return account, err
}

// Update updates a customer's properties.
// For more details see https://developer.gocardless.com/api-reference/#customer-bank-accounts-update-a-customer-bank-account
func (s *CustomerBankAccountService) UpdateCustomerBankAccount(customerBankAccount *CustomerBankAccount, metadata map[string]string) (*CustomerBankAccount, error) {
	u := fmt.Sprintf("customer_bank_accounts/%s", customerBankAccount.Id)
	params := url.Values{}
	metadataString, _ := json.Marshal(metadata)
	customerBankAccount.Metadata = metadata
	rel := map[string]interface{}{
		"customer_bank_accounts": customerBankAccount,
	}
	params.Add("metadata", string(metadataString))
	account := &CustomerBankAccount{}
	err := s.client.Call("PUT", u, rel, account)

	return account, err
}

// Immediately disables the bank account, no money can be paid out to a disabled account.
// https://developer.gocardless.com/api-reference/#customer-bank-accounts-disable-a-customer-bank-account
func (s *CustomerBankAccountService) DisableCustomerBankAccount(bankAccount *CustomerBankAccount) (*Response, error) {
	u := fmt.Sprintf("/customer_bank_accounts/%s/actions/disable", bankAccount.Id)
	resp := &Response{}
	err := s.client.Call("POST", u, bankAccount, resp)

	return resp, err
}
