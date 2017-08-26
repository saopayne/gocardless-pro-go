package gocardless_pro_go

import (
	"fmt"
	"net/url"
)

type RefundService service

type Refund struct {
	ID        string                `json:"id,omitempty"`
	CreatedAt string                `json:"created_at,omitempty"`
	Reference string                `json:"reference,omitempty"`
	Currency  string                `json:"currency,omitempty"`
	Amount    int64                 `json:"amount,omitempty"`
	Links     []RefundLink          `json:"links,omitempty"`
	Metadata  map[string]string     `json:"metadata,omitempty"`
}

type RefundListRequest struct {
	CreatedAt CreatedAt     `json:"created_at,omitempty"`
	Limit     int           `json:"limit,omitempty"`
	Before    string        `json:"before,omitempty"`
	After     string        `json:"after,omitempty"`
	Payment   string        `json:"payment,omitempty"`
}

type RefundList struct {
	Meta   ListMeta
	Values []Refund `json:"data"`
}

type RefundCreateRequest struct {
	Metadata                map[string]string       `json:"metadata,omitempty"`
	Reference               string                  `json:"reference,omitempty"`
	Amount                  string                  `json:"amount,omitempty"`
	TotalAmountConfirmation string                  `json:"total_amount_confirmation,omitempty"`
	Links                   []string                `json:"links,omitempty"`
}

// Create creates a new refund
func (s *RefundService) CreateRefund(refundReq *RefundCreateRequest) (*Refund, error) {
	u := fmt.Sprintf("/refunds")
	refund := &Refund{}
	err := s.client.Call("POST", u, refundReq, refund)

	return refund, err
}

// List returns a list of refunds
func (s *RefundService) ListRefunds(req *RefundListRequest) (*RefundList, error) {
	return s.ListNRefunds(10, 0, req)
}

func (s *RefundService) ListNRefunds(count, offset int, req *RefundListRequest) (*RefundList, error) {
	params := url.Values{}
	params.Add("after", req.After)
	params.Add("before", req.Before)
	params.Add("created_at[gt]", req.CreatedAt.Gt)
	params.Add("created_at[gte]", req.CreatedAt.Gte)
	params.Add("created_at[lt]", req.CreatedAt.Lt)
	params.Add("created_at[lte]", req.CreatedAt.Lte)
	params.Add("limit", string(req.Limit))
	params.Add("payment", req.Payment)

	u := paginateURL("/refunds", count, offset)
	refunds := &RefundList{}
	err := s.client.Call("GET", u, params, refunds)

	return refunds, err
}

func (s *RefundService) GetRefund(id string) (*Refund, error) {
	u := fmt.Sprintf("/refunds/%s", id)
	refund := &Refund{}
	err := s.client.Call("GET", u, nil, payment)

	return refund, err
}

func (s *RefundService) UpdateRefund(updatedRefund *Refund, metadata map[string]string) (*Refund, error) {
	params := url.Values{}
	params.Add("metadata", string(metadata))
	u := fmt.Sprintf("/refunds/%d", updatedRefund.ID)
	refund := &Refund{}
	err := s.client.Call("PUT", u, params, refund)

	return refund, err
}
