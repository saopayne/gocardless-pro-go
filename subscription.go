package gocardless

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type SubscriptionService service

type Subscription struct {
	ID               string             `json:"id,omitempty"`
	CreatedAt        string             `json:"createdAt,omitempty"`
	EndDate          string             `json:"end_date,omitempty"`
	Amount           int                `json:"amount,omitempty"`
	Currency         string             `json:"currency,omitempty"`
	DayOfMonth       int                `json:"day_of_month,omitempty"`
	Interval         int                `json:"interval,omitempty"`
	StartDate        string             `json:"start,omitempty"`
	Status           Status             `json:"status,omitempty"`
	Links            []SubscriptionLink `json:"links,omitempty"`
	Metadata         map[string]string  `json:"metadata,omitempty"`
	Month            Month              `json:"month,omitempty"`
	IntervalUnit     IntervalUnit       `json:"interval_unit,omitempty"`
	UpcomingPayments []UpcomingPayment  `json:"upcoming_payments,omitempty"`
	Name             string             `json:"name,omitempty"`
	PaymentReference string             `json:"payment_reference,omitempty"`
}

type IntervalUnit struct {
	IntervalUnit string `json:"interval_unit,omitempty"`
}

type Month struct {
	Month string `json:"month,omitempty"`
}

type UpcomingPayment struct {
	Amount     int    `json:"amount,omitempty"`
	ChargeDate string `json:"charge_date,omitempty"`
}

type SubscriptionListRequest struct {
	CreatedAt CreatedAt `json:"created_at,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	Before    string    `json:"before,omitempty"`
	After     string    `json:"after,omitempty"`
	Customer  string    `json:"customer,omitempty"`
	Mandate   string    `json:"mandate,omitempty"`
}

type SubscriptionList struct {
	Meta   ListMeta
	Values []Subscription `json:"data"`
}

type SubscriptionCreateRequest struct {
	Metadata         map[string]string `json:"metadata,omitempty"`
	PaymentReference string            `json:"payment_reference,omitempty"`
	Links            []string          `json:"links,omitempty"`
	StartDate        string            `json:"start_date,omitempty"`
	Name             string            `json:"name,omitempty"`
	Month            string            `json:"month,omitempty"`
	Interval         int               `json:"interval,omitempty"`
	EndDate          string            `json:"end_date,omitempty"`
	DayOfMonth       string            `json:"day_of_month,omitempty"`
	IntervalUnit     string		      `json:"interval_unit,omitempty"`
	Amount           int64            `json:"amount,omitempty"`
	Currency         string            `json:"currency,omitempty"`
	AppFee           string            `json:"app_fee,omitempty"`
	Count            string            `json:"count,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	AccountNumber    string            `json:"account_number,omitempty"`
}


type SubscriptionCancelRequest struct {
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Create creates a new subscription
func (s *SubscriptionService) CreateSubscription(subscriptionReq *SubscriptionCreateRequest) (*Subscription, error) {
	u := fmt.Sprintf("/subscriptions")
	subscription := &Subscription{}
	rel := map[string]interface{}{
		"subscriptions": subscriptionReq,
	}
	err := s.client.Call("POST", u, rel, subscription)

	return subscription, err
}

// List returns a list of subscriptions
func (s *SubscriptionService) ListSubscriptions(req *SubscriptionListRequest) (*SubscriptionList, error) {
	reqd, err := http.NewRequest("GET", "/subscriptions", nil)
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
	if req.Customer != "" { params.Add("customer", req.Customer) }

	reqd.URL.RawQuery = params.Encode()
	path := reqd.URL.String()
	subscriptions := &SubscriptionList{}
	err = s.client.Call("GET", path, nil, subscriptions)

	return subscriptions, err
}

func (s *SubscriptionService) GetSubscription(id string) (*Subscription, error) {
	u := fmt.Sprintf("/subscriptions/%s", id)
	sub := &Subscription{}
	err := s.client.Call("GET", u, nil, sub)

	return sub, err
}

func (s *SubscriptionService) UpdateSubscription(updatedSubscription *Subscription, name string, reference string, metadata map[string]string) (*Subscription, error) {
	metadataString, _ := json.Marshal(metadata)
	metaJson := string(metadataString[:])
	metadataMap := make(map[string]string)
	metadataMap["metadata"] = string(metaJson[:])
	metadataMap["name"] = name
	metadataMap["payment_reference"] = reference
	rel := map[string]interface{}{
		"mandates": metadataMap,
	}

	u := fmt.Sprintf("/subscriptions/%s", updatedSubscription.ID)
	sub := &Subscription{}
	err := s.client.Call("PUT", u, rel, sub)

	return sub, err
}

func (s *SubscriptionService) CancelSubscription(subscriptionToCancel *Subscription, metadata map[string]string) (*Response, error) {
	metadataString, _ := json.Marshal(metadata)
	metaJson := string(metadataString[:])
	metadataMap := make(map[string]string)
	metadataMap["metadata"] = string(metaJson[:])
	rel := map[string]interface{}{
		"subscriptions": metadataMap,
	}
	u := fmt.Sprintf("/subscriptions/%s/actions/cancel", subscriptionToCancel.ID)
	resp := &Response{}
	err := s.client.Call("POST", u, rel, resp)

	return resp, err
}
