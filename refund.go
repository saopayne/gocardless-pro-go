package gocardless_pro_go

type RefundService service


type Refund struct {
	ID							string				`json:"id,omitempty"`
	CreatedAt					string				`json:"created_at,omitempty"`
	Reference 					string				`json:"reference,omitempty"`
	Currency					string				`json:"currency,omitempty"`
	Amount						int64				`json:"amount,omitempty"`
	Links						[]RefundLink		`json:"links,omitempty"`
	Metadata					map[string]string	`json:"metadata,omitempty"`
}
