package gocardless_pro_go

// SubscriptionService handles operations related to the subscription
// For more details see https://developers.paystack.co/v1.0/reference#create-subscription
type SubscriptionService service

type Subscription struct {
	ID          int    `json:"id,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Integration int    `json:"integration,omitempty"`
	// inconsistent API response. Create returns Customer code, Fetch returns an object
	Customer  interface{} `json:"customer,omitempty"`
	Plan      Plan        `json:"plan,omitempty"`
	StartDate string      `json:"start,omitempty"`
	// inconsistent API response. Fetch returns string, List returns an object
	Authorization    interface{}   `json:"authorization,omitempty"`
	Invoices         []interface{} `json:"invoices,omitempty"`
	Status           string        `json:"status,omitempty"`
	Quantity         int           `json:"quantity,omitempty"`
	Amount           int           `json:"amount,omitempty"`
	SubscriptionCode string        `json:"subscription_code,omitempty"`
	EmailToken       string        `json:"email_token,omitempty"`
	EasyCronID       string        `json:"easy_cron_id,omitempty"`
	CronExpression   string        `json:"cron_expression,omitempty"`
	NextPaymentDate  string        `json:"next_payment_date,omitempty"`
	OpenInvoice      string        `json:"open_invoice,omitempty"`
}

