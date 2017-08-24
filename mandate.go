package gocardless_pro_go

import (
	"fmt"
	"net/url"
)

type MandateService service

type Mandate struct {
	ID							string				`json:"id,omitempty"`
	CreatedAt					string				`json:"created_at,omitempty"`
	Reference 					string				`json:"reference,omitempty"`
	Scheme						string				`json:"scheme,omitempty"`
	Status						Status				`json:"status,omitempty"`
	PaymentsRequireApproval		bool				`json:"payments_require_approval,omitempty"`
	NextPossibleChargeDate		string				`json:"next_possible_charge_date,omitempty"`
	Links						[]MandateLink		`json:"links,omitempty"`
	Metadata					map[string]string	`json:"metadata,omitempty"`
}

type Status struct {
	Status 	string	`json:"status,omitempty"`
}

type MandateListRequest struct {
	CreatedAt 			CreatedAt	`json:"created_at,omitempty"`
	Limit				int			`json:"limit,omitempty"`
	Before				string		`json:"before,omitempty"`
	After				string		`json:"after,omitempty"`
	Creditor			string		`json:"creditor,omitempty"`
	Customer			string		`json:"customer,omitempty"`
	CustomerBankAccount string		`json:"customer_bank_account,omitempty"`
	Reference			string		`json:"reference,omitempty"`
	Status				string		`json:"status,omitempty"`
}

// CustomerBankAccountList is a list object for customer bank accounts.
type MandateList struct {
	Meta   ListMeta
	Values []CustomerBankAccount `json:"data"`
}

type MandateCreateRequest struct {
	Metadata			map[string]string		`json:"metadata,omitempty"`
	Reference			string					`json:"reference,omitempty"`
	Scheme				string					`json:"scheme,omitempty"`
	AccountNumber		string					`json:"account_number,omitempty"`
	Links				[]string				`json:"links,omitempty"`

}

type MandateCancelRequest struct {
	Identity 	string  	`json:"identity,omitempty"`
}


// Create creates a new mandate
func (s *MandateService) CreateMandate(mandateReq *MandateCreateRequest) (*Mandate, error) {
	u := fmt.Sprintf("/mandates")
	mandate := &Mandate{}
	err := s.client.Call("POST", u, mandateReq, mandate)

	return mandate, err
}

// List returns a list of mandates
func (s *MandateService) ListMandates(req *MandateListRequest) (*MandateList, error) {
	return s.ListNMandates(10, 0, req)
}

func (s *MandateService) ListNMandates(count, offset int, req *MandateListRequest) (*MandateList, error) {
	params := url.Values{}
	params.Add("after", req.After)
	params.Add("before", req.Before)
	params.Add("created_at[gt]", req.CreatedAt.Gt)
	params.Add("created_at[gte]", req.CreatedAt.Gte)
	params.Add("created_at[lt]", req.CreatedAt.Lt)
	params.Add("created_at[lte]", req.CreatedAt.Lte)
	params.Add("limit", string(req.Limit))
	params.Add("reference", req.Reference)
	params.Add("status", req.Status)
	params.Add("customer_bank_account", req.CustomerBankAccount)
	params.Add("customer", req.Customer)
	params.Add("creditor", req.Creditor)
	u := paginateURL("/mandates", count, offset)
	mandates := &MandateList{}
	err := s.client.Call("GET", u, params, mandates)

	return mandates, err
}


func (s *MandateService) GetMandate(id string) (*Mandate, error) {
	u := fmt.Sprintf("/customer_bank_accounts/%s", id)
	mandate := &Mandate{}
	err := s.client.Call("GET", u, nil, mandate)

	return mandate, err
}


// Update updates a customer's properties.
// For more details see https://developer.gocardless.com/api-reference/#customer-bank-accounts-update-a-customer-bank-account
func (s *CustomerBankAccountService) UpdateMandate(customerBankAccount *CustomerBankAccount) (*CustomerBankAccount, error) {
	u := fmt.Sprintf("customer_bank_accounts/%d", customerBankAccount.Id)
	account := &CustomerBankAccount{}
	err := s.client.Call("PUT", u, customerBankAccount, account)

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

