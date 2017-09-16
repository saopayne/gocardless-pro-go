package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"log"
	"net/http"
	"os"
)

type PaymentService service

type Payment struct {
	Amount         int64             `json:"amount, omitempty"`
	AmountRefunded int64             `json:"amount_refunded,omitempty"`
	ChargeDate     string            `json:"charge_date,omitempty"`
	CreatedAt      string            `json:"created_at,omitempty"`
	Description    string            `json:"description,omitempty"`
	ID             string            `json:"id,omitempty"`
	Reference      string            `json:"reference,omitempty"`
	Status         Status            `json:"status,omitempty"`
	Currency       Currency          `json:"currency,omitempty"`
	Links          map[string]string `json:"links,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
}

type Currency struct {
	Currency string `json:"currency,omitempty"`
}

type PaymentListRequest struct {
	CreatedAt    CreatedAt `json:"created_at,omitempty"`
	Limit        int       `json:"limit,omitempty"`
	Before       string    `json:"before,omitempty"`
	After        string    `json:"after,omitempty"`
	Creditor     string    `json:"creditor,omitempty"`
	Customer     string    `json:"customer,omitempty"`
	Status       string    `json:"status,omitempty"`
	Currency     string    `json:"currency,omitempty"`
	Mandate      string    `json:"mandate,omitempty"`
	Subscription string    `json:"subscription,omitempty"`
}

type PaymentList struct {
	Meta   ListMeta
	Values []Payment `json:"data"`
}

type PaymentCreateRequest struct {
	Metadata      map[string]string `json:"metadata,omitempty"`
	Reference     string            `json:"reference,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	AccountNumber string            `json:"account_number,omitempty"`
	Links         []string          `json:"links,omitempty"`
}

type PaymentCancelRequest struct {
	Identity string `json:"identity,omitempty"`
}

// Create creates a new payment
func (s *PaymentService) CreatePayment(paymentReq *PaymentCreateRequest) (*Payment, error) {
	u := fmt.Sprintf("/payments")
	payment := &Payment{}
	err := s.client.Call("POST", u, paymentReq, payment)

	return payment, err
}

// List returns a list of payments
func (s *PaymentService) ListPayments(req *PaymentListRequest) (*PaymentList, error) {
	reqd, err := http.NewRequest("GET", "/payments", nil)
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
	if req.Mandate != "" { params.Add("mandate", req.Mandate) }
	if req.Status != "" { params.Add("status", req.Status) }
	if req.Currency != "" {params.Add("currency", req.Currency)}
	if req.Customer != "" {params.Add("customer", req.Customer)}
	if req.Creditor != "" {params.Add("creditor", req.Creditor)}
	if req.Subscription != "" {params.Add("subscription", req.Subscription)}

	reqd.URL.RawQuery = params.Encode()
	path := reqd.URL.String()
	payments := &PaymentList{}
	err = s.client.Call("GET", path, nil, payments)

	return payments, err
}

func (s *PaymentService) GetPayment(id string) (*Payment, error) {
	u := fmt.Sprintf("/payments/%s", id)
	payment := &Payment{}
	err := s.client.Call("GET", u, nil, payment)

	return payment, err
}

func (s *PaymentService) UpdatePayment(updatedPayment *Payment, metadata map[string]string) (*Payment, error) {
	params := url.Values{}
	metadataString, _ := json.Marshal(metadata)
	params.Add("metadata", string(metadataString))
	u := fmt.Sprintf("/payments/%s", updatedPayment.ID)
	payment := &Payment{}
	err := s.client.Call("PUT", u, params, payment)

	return payment, err
}

func (s *PaymentService) CancelPayment(paymentToCancel *Payment, metadata map[string]string) (*Response, error) {
	params := url.Values{}
	metadataString, _ := json.Marshal(metadata)
	params.Add("metadata", string(metadataString))
	u := fmt.Sprintf("/payments/%s/actions/cancel", paymentToCancel.ID)
	resp := &Response{}
	err := s.client.Call("POST", u, params, resp)

	return resp, err
}

func (s *MandateService) RetryPayment(payment *Payment, metadata map[string]string) (*Response, error) {
	params := url.Values{}
	metadataString, _ := json.Marshal(metadata)
	params.Add("metadata", string(metadataString))
	u := fmt.Sprintf("/payments/%s/actions/retry", payment.ID)
	resp := &Response{}
	err := s.client.Call("POST", u, params, resp)

	return resp, err
}
