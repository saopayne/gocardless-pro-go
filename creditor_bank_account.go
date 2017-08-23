package gocardless_pro_go

type CreditorBankAccountService service

type CreditorBankAccount struct {
	Id        			string  					`json:"id,omitempty"`
	BankName      		string  					`json:"bank_name,omitempty"`
	CountryCode			string						`json:"country_code,omitempty"`
	CreatedAt			string						`json:"created_at,omitempty"`
	Currency			string 						`json:"currency,omitempty"`
	AccountHolderName	string						`json:"account_holder_name,omitempty"`
	AccountNumberEnding	string						`json:"account_number_ending,omitempty"`
	Enabled				bool						`json:"enabled,omitempty"`
	Links				[]CreditorLink				`json:"links,omitempty"`
	Metadata			map[string]string			`json:"metadata,omitempty"`
	ResponseUrl	  		string						`json:"responseurl,omitempty"`
}

