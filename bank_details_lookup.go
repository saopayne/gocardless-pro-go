package gocardless_pro_go

import (
	"fmt"
	"golang.org/x/crypto/nacl/box"
)

type BankDetailsLookupService service

const (
	AUTOGIRO 	= "autogiro"
	BACS 		= "bacs"
	SEPA_CORE 	= "sepa_core"
)

type AvailableDebitScheme struct {
	Name string
}

type AvailableDebitSchemeList struct {
	list []AvailableDebitScheme
}


type BankDetailsLookup struct {
	BankName        string  					`json:"bank_name,omitempty"`
	BIC            	string 						`json:"bic,omitempty"`
	CustomerId	    AvailableDebitSchemeList 	`json:"available_debit_schemes,omitempty"`
	ResponseUrl	  	string						`json:"responseurl,omitempty"`
	Metadata        Metadata 					`json:"metadata,omitempty"`
}

func (availableDebitSchemeList *AvailableDebitSchemeList) AddDebitScheme(debitScheme AvailableDebitScheme) []AvailableDebitScheme {
	availableDebitSchemeList.list = append(availableDebitSchemeList.list, debitScheme)
	return availableDebitSchemeList.list
}
