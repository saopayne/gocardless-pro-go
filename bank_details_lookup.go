package main

import (
	"fmt"
)

type BankDetailsLookupService service

const (
	AUTOGIRO  = "autogiro"
	BACS      = "bacs"
	SEPA_CORE = "sepa_core"
)

type AvailableDebitScheme struct {
	Name string
}

type AvailableDebitSchemeList struct {
	list []AvailableDebitScheme
}

type BankDetailsLookupRequest struct {
	AccountNumber string `json:"account_number,omitempty"`
	BankCode      string `json:"bank_code,omitempty"`
	BranchCode    string `json:"branch_code,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	Iban          string `json:"iban,omitempty"`
}

type BankDetailsLookup struct {
	BankName    string                   `json:"bank_name,omitempty"`
	BIC         string                   `json:"bic,omitempty"`
	CustomerId  AvailableDebitSchemeList `json:"available_debit_schemes,omitempty"`
	ResponseUrl string                   `json:"responseurl,omitempty"`
	Metadata    Metadata                 `json:"metadata,omitempty"`
}

func (availableDebitSchemeList *AvailableDebitSchemeList) AddDebitScheme(debitScheme AvailableDebitScheme) []AvailableDebitScheme {
	availableDebitSchemeList.list = append(availableDebitSchemeList.list, debitScheme)
	return availableDebitSchemeList.list
}

// Performs a bank details lookup
// As part of the lookup a modulus check and reachability check are performed.
func (s *BankDetailsLookupService) Lookup(txn *BankDetailsLookupRequest) (*Response, error) {
	u := fmt.Sprintf("/bank_details_lookups")
	resp := &Response{}
	err := s.client.Call("POST", u, txn, resp)

	return resp, err
}
