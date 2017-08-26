package gocardless_pro_go

import (
	"fmt"
	"time"
	"net/url"
)

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


type SubscriptionListRequest struct {
	CreatedAt 			CreatedAt	`json:"created_at,omitempty"`
	Limit				int			`json:"limit,omitempty"`
	Before				string		`json:"before,omitempty"`
	After				string		`json:"after,omitempty"`
	Customer			string		`json:"customer,omitempty"`
	Mandate				string		`json:"mandate,omitempty"`
}

type SubscriptionList struct {
	Meta   ListMeta
	Values []Subscription `json:"data"`
}

type SubscriptionCreateRequest struct {
	Metadata			map[string]string		`json:"metadata,omitempty"`
	PaymentReference	string					`json:"payment_reference,omitempty"`
	Links				[]string				`json:"links,omitempty"`
	StartDate			string					`json:"start_date,omitempty"`
	Name				string					`json:"name,omitempty"`
	Month				string					`json:"month,omitempty"`
	Interval			int						`json:"interval,omitempty"`
	EndDate				string					`json:"end_date,omitempty"`
	DayOfMonth			string					`json:"day_of_month,omitempty"`
	IntervalUnit		IntervalUnit			`json:"interval_unit,omitempty"`
	Amount				string					`json:"amount,omitempty"`
	Currency			string					`json:"currency,omitempty"`
	AppFee				string					`json:"app_fee,omitempty"`
	Count				string					`json:"count,omitempty"`
	Scheme				string					`json:"scheme,omitempty"`
	AccountNumber		string					`json:"account_number,omitempty"`

}

type SubscriptionCancelRequest struct {
	Identity 	string  	`json:"identity,omitempty"`
}


// Create creates a new subscription
func (s *SubscriptionService) CreateSubscription(subscriptionReq *SubscriptionCreateRequest) (*Subscription, error) {
	u := fmt.Sprintf("/subscriptions")
	subscription := &Subscription{}
	err := s.client.Call("POST", u, subscriptionReq, subscription)

	return subscription, err
}

// List returns a list of subscriptions
func (s *SubscriptionService) ListSubscriptions(req *SubscriptionListRequest) (*SubscriptionList, error) {
	return s.ListNSubscriptions(10, 0, req)
}

func (s *SubscriptionService) ListNSubscriptions(count, offset int, req *SubscriptionListRequest) (*SubscriptionList, error) {
	params := url.Values{}
	params.Add("after", req.After)
	params.Add("before", req.Before)
	params.Add("created_at[gt]", req.CreatedAt.Gt)
	params.Add("created_at[gte]", req.CreatedAt.Gte)
	params.Add("created_at[lt]", req.CreatedAt.Lt)
	params.Add("created_at[lte]", req.CreatedAt.Lte)
	params.Add("limit", string(req.Limit))
	params.Add("mandate", req.Mandate)
	params.Add("customer", req.Customer)

	u := paginateURL("/subscriptions", count, offset)
	subscriptions := &SubscriptionList{}
	err := s.client.Call("GET", u, params, subscriptions)

	return subscriptions, err
}


func (s *PaymentService) GetPayment(id string) (*Payment, error) {
	u := fmt.Sprintf("/payments/%s", id)
	payment := &Payment{}
	err := s.client.Call("GET", u, nil, payment)

	return payment, err
}


func (s *PaymentService) UpdatePayment(updatedPayment *Payment, metadata map[string]string) (*Payment, error) {
	params := url.Values{}
	params.Add("metadata", string(metadata))
	u := fmt.Sprintf("/payments/%d", updatedPayment.ID)
	payment := &Payment{}
	err := s.client.Call("PUT", u, params, payment)

	return payment, err
}


func (s *PaymentService) CancelPayment(paymentToCancel *Payment, metadata map[string]string) (*Response, error) {
	params := url.Values{}
	params.Add("metadata", string(metadata))
	u := fmt.Sprintf("/payments/%s/actions/cancel", paymentToCancel.ID)
	resp := &Response{}
	err := s.client.Call("POST", u, params, resp)

	return resp, err
}

func (s *MandateService) RetryPayment(payment *Payment, metadata map[string]string) (*Response, error) {
	params := url.Values{}
	params.Add("metadata", string(metadata))
	u := fmt.Sprintf("/payments/%s/actions/retry", payment.ID)
	resp := &Response{}
	err := s.client.Call("POST", u, params, resp)

	return resp, err
}


