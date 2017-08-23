package gocardless_pro_go


type SubscriptionService service


type Subscription struct {
	ID          		string   			`json:"id,omitempty"`
	CreatedAt   		string 				`json:"createdAt,omitempty"`
	EndDate   			string 				`json:"end_date,omitempty"`
	Amount      		int    				`json:"amount,omitempty"`
	Currency			string				`json:"currency,omitempty"`
	DayOfMonth			int					`json:"day_of_month,omitempty"`
	Interval			int					`json:"interval,omitempty"`
	StartDate 			string      		`json:"start,omitempty"`
	Status          	Status      		`json:"status,omitempty"`
	Links				[]SubscriptionLink	`json:"links,omitempty"`
	Metadata			map[string]string	`json:"metadata,omitempty"`
	Month				Month				`json:"month,omitempty"`
	IntervalUnit		IntervalUnit		`json:"interval_unit,omitempty"`
	UpcomingPayments	[]UpcomingPayment	`json:"upcoming_payments,omitempty"`
	Name				string				`json:"name,omitempty"`
	PaymentReference	string				`json:"payment_reference,omitempty"`
}

type IntervalUnit struct {
	IntervalUnit	string 		`json:"interval_unit,omitempty"`
}

type Month	struct {
	Month	string 	`json:"month,omitempty"`
}

type UpcomingPayment struct {
	Amount		int			`json:"amount,omitempty"`
	ChargeDate	string		`json:"charge_date,omitempty"`
}

