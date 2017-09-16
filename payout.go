package gocardless

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type PayoutService service

type Payout struct {
	Amount       int64             `json:"amount, omitempty"`
	DeductedFees int64             `json:"deducted_fees,omitempty"`
	ArrivalDate  string            `json:"arrival_date,omitempty"`
	CreatedAt    string            `json:"created_at,omitempty"`
	ID           string            `json:"id,omitempty"`
	Reference    string            `json:"reference,omitempty"`
	Status       Status            `json:"status,omitempty"`
	Currency     Currency          `json:"currency,omitempty"`
	Links        []PayoutLink      `json:"links,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	PayoutType   PayoutType        `json:"payout_type,omitempty"`
}

type PayoutType struct {
	PayoutType string `json:"currency,omitempty"`
}

type PayoutListRequest struct {
	CreatedAt           CreatedAt `json:"created_at,omitempty"`
	Limit               int       `json:"limit,omitempty"`
	Before              string    `json:"before,omitempty"`
	After               string    `json:"after,omitempty"`
	Creditor            string    `json:"creditor,omitempty"`
	CreditorBankAccount string    `json:"creditor_bank_account,omitempty"`
	Status              string    `json:"status,omitempty"`
	Currency            string    `json:"currency,omitempty"`
	PayoutType          string    `json:"payout_type,omitempty"`
}

type PayoutList struct {
	Meta   ListMeta
	Values []Payout `json:"data"`
}

// List returns a list of payments
func (s *PayoutService) ListPayouts(req *PayoutListRequest) (*PayoutList, error) {
	reqd, err := http.NewRequest("GET", "/payouts", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	params := reqd.URL.Query()
	if req.After != "" { params.Add("after", req.After) }
	if req.Before != "" { params.Add("before", req.Before) }
	if req.CreatedAt.Gt != "" { params.Add("created_at[gt]", req.CreatedAt.Gt) }
	if req.CreatedAt.Gte != "" { params.Add("created_at[gte]", req.CreatedAt.Gte) }
	if req.CreatedAt.Lt != "" { params.Add("created_at[lt]", req.CreatedAt.Lt) }
	if req.CreatedAt.Lte != "" { params.Add("created_at[lte]", req.CreatedAt.Lte) }
	if req.Limit > 0 { params.Add("limit", string(req.Limit)) }
	if req.Status != "" { params.Add("status", req.Status) }
	if req.Currency != "" {params.Add("currency", req.Currency)}
	if req.CreditorBankAccount != "" {params.Add("creditor_bank_account", req.CreditorBankAccount)}
	if req.Creditor != "" {params.Add("creditor", req.Creditor)}
	if req.PayoutType != "" {params.Add("payout_type", req.PayoutType)}

	reqd.URL.RawQuery = params.Encode()
	path := reqd.URL.String()
	payoutList := &PayoutList{}
	err = s.client.Call("GET", path, nil, payoutList)

	return payoutList, err
}

func (s *PayoutService) GetPayout(id string) (*Payout, error) {
	u := fmt.Sprintf("/payouts/%s", id)
	payout := &Payout{}
	err := s.client.Call("GET", u, nil, payout)

	return payout, err
}
