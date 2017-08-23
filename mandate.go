package gocardless_pro_go


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
