package gocardless_pro_go

type CreditorLink struct {
	Creditor       string  	 `json:"creditor,omitempty"`
}

type CustomerLink struct {
	Customer       string  	 `json:"customer,omitempty"`
}

type EventLink	struct {
	Mandate 						string		`json:"mandate,omitempty"`
	NewCustomerBankAccount			string		`json:"new_customer_bank_account,omitempty"`
	NewMandate						string		`json:"new_mandate,omitempty"`
	Organisation					string		`json:"organisation,omitempty"`
	ParentEvent						string		`json:"parent_event,omitempty"`
	Payment							string		`json:"payment,omitempty"`
	Payout							string		`json:"payout,omitempty"`
	PreviousCustomerBankAccount		string		`json:"previous_customer_bank_account,omitempty"`
	Refund							string		`json:"refund,omitempty"`
	Subscription					string		`json:"subscription,omitempty"`
}

type MandateLink struct {
	Creditor				string        	`json:"creditor,omitempty"`
	Customer				string          `json:"customer,omitempty"`
	CustomerBankAccount		string			`json:"customer_bank_account,omitempty"`
	NewMandate				string			`json:"new_mandate,omitempty"`
}

