package gocardless_pro_go

import "fmt"

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

type RefundListRequest struct {
	CreatedAt 			CreatedAt	`json:"created_at,omitempty"`
	Limit				int			`json:"limit,omitempty"`
	Before				string		`json:"before,omitempty"`
	After				string		`json:"after,omitempty"`
	Creditor			string		`json:"creditor,omitempty"`
	Customer			string		`json:"customer,omitempty"`
	Status				string		`json:"status,omitempty"`
	Currency			string		`json:"currency,omitempty"`
	Mandate				string		`json:"mandate,omitempty"`
	Subscription		string		`json:"subscription,omitempty"`
}

type RefundList struct {
	Meta   ListMeta
	Values []Refund `json:"data"`
}

type RefundCreateRequest struct {
	Metadata				map[string]string		`json:"metadata,omitempty"`
	Reference				string					`json:"reference,omitempty"`
	Amount					string					`json:"amount,omitempty"`
	TotalAmountConfirmation	string					`json:"total_amount_confirmation,omitempty"`
	Links					[]string				`json:"links,omitempty"`
}



// Create creates a new refund
func (s *RefundService) CreateRefund(refundReq *RefundCreateRequest) (*Refund, error) {
	u := fmt.Sprintf("/refunds")
	refund := &Refund{}
	err := s.client.Call("POST", u, refundReq, refund)

	return refund, err
}

// List returns a list of refund
func (s *PaymentService) ListPayments(req *PaymentListRequest) (*PaymentList, error) {
	return s.ListNPayments(10, 0, req)
}

func (s *PaymentService) ListNPayments(count, offset int, req *PaymentListRequest) (*PaymentList, error) {
	params := url.Values{}
	params.Add("after", req.After)
	params.Add("before", req.Before)
	params.Add("created_at[gt]", req.CreatedAt.Gt)
	params.Add("created_at[gte]", req.CreatedAt.Gte)
	params.Add("created_at[lt]", req.CreatedAt.Lt)
	params.Add("created_at[lte]", req.CreatedAt.Lte)
	params.Add("limit", string(req.Limit))
	params.Add("status", req.Status)
	params.Add("mandate", req.Mandate)
	params.Add("customer", req.Customer)
	params.Add("creditor", req.Creditor)
	params.Add("currency", req.Currency)
	params.Add("subscription", req.Subscription)

	u := paginateURL("/payments", count, offset)
	payments := &PaymentList{}
	err := s.client.Call("GET", u, params, payments)

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
	params.Add("metadata", string(metadata))
	u := fmt.Sprintf("payments/%d", updatedPayment.ID)
	payment := &Payment{}
	err := s.client.Call("PUT", u, params, payment)

	return payment, err
}
