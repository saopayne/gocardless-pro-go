package gocardless_pro_go


type PayoutService service

type Payout struct {
	Amount 			int64       		`json:"amount, omitempty"`
	DeductedFees	int64				`json:"deducted_fees,omitempty"`
	ArrivalDate		string				`json:"arrival_date,omitempty"`
	CreatedAt		string				`json:"created_at,omitempty"`
	ID				string				`json:"id,omitempty"`
	Reference		string				`json:"reference,omitempty"`
	Status			Status				`json:"status,omitempty"`
	Currency		Currency			`json:"currency,omitempty"`
	Links			[]PayoutLink		`json:"links,omitempty"`
	Metadata		map[string]string	`json:"metadata,omitempty"`
	PayoutType		PayoutType          `json:"payout_type,omitempty"`
}

type PayoutType	struct {
	PayoutType 	string		`json:"currency,omitempty"`
}

