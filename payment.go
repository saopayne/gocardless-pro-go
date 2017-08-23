package gocardless_pro_go

type PaymentService service


type Payment struct {
	Amount 			int64       		`json:"amount, omitempty"`
	AmountRefunded	int64				`json:"amount_refunded,omitempty"`
	ChargeDate		string				`json:"charge_date,omitempty"`
	CreatedAt		string				`json:"created_at,omitempty"`
	Description 	string				`json:"description,omitempty"`
	ID				string				`json:"id,omitempty"`
	Reference		string				`json:"reference,omitempty"`
	Status			Status				`json:"status,omitempty"`
	Currency		Currency			`json:"currency,omitempty"`
	Links			[]PaymentLink		`json:"links,omitempty"`
	Metadata		map[string]string	`json:"metadata,omitempty"`
}

type  Currency	struct {
	Currency 	string		`json:"currency,omitempty"`
}

