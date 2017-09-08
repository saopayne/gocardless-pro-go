package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type RefundService service

type Refund struct {
	ID        string            `json:"id,omitempty"`
	CreatedAt string            `json:"created_at,omitempty"`
	Reference string            `json:"reference,omitempty"`
	Currency  string            `json:"currency,omitempty"`
	Amount    int64             `json:"amount,omitempty"`
	Links     []RefundLink      `json:"links,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type RefundListRequest struct {
	CreatedAt CreatedAt `json:"created_at,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	Before    string    `json:"before,omitempty"`
	After     string    `json:"after,omitempty"`
	Payment   string    `json:"payment,omitempty"`
}

type RefundList struct {
	Meta   ListMeta
	Values []Refund `json:"data"`
}

type RefundCreateRequest struct {
	Metadata                map[string]string `json:"metadata,omitempty"`
	Reference               string            `json:"reference,omitempty"`
	Amount                  int64             `json:"amount,omitempty"`
	TotalAmountConfirmation string            `json:"total_amount_confirmation,omitempty"`
	Links                   map[string]string `json:"links,omitempty"`

}

//Creates a new refund object.
//This fails with:
//refund_payment_invalid_state error if the linked payment isn’t either confirmed or paid_out.
//total_amount_confirmation_invalid if the confirmation amount doesn’t match the total amount refunded for the payment. This safeguard is there to prevent two processes from creating refunds without awareness of each other.
//number_of_refunds_exceeded if five or more refunds have already been created against the payment.
func (s *RefundService) CreateRefund(refundReq *RefundCreateRequest) (*Refund, error) {
	u := fmt.Sprintf("/refunds")
	refund := &Refund{}

	linksString, _ := json.Marshal(refundReq.Links)
	linksJson := string(linksString[:])
	linksMap := make(map[string]string)

	linksMap["links"] = string(linksJson[:])
	refundReq.Links = linksMap
	fmt.Println(refundReq)
	refundsMap := map[string]interface{}{
		"refunds": refundReq,
	}
	err := s.client.Call("POST", u, refundsMap, refund)

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

//Retrieves all details for a single refund
//Relative endpoint: GET /refunds/RF123
func (s *RefundService) GetRefund(id string) (*Refund, error) {
	u := fmt.Sprintf("/refunds/%s", id)
	refund := &Refund{}
	err := s.client.Call("GET", u, nil, refund)

	return refund, err
}

//Updates a refund object.
//Relative endpoint: PUT /refunds/RF123
func (s *RefundService) UpdateRefund(updatedRefund *Refund, metadata map[string]string) (*Refund, error) {
	refund := &Refund{}

	metadataString, _ := json.Marshal(metadata)
	metaJson := string(metadataString[:])
	metadataMap := make(map[string]string)
	metadataMap["metadata"] = string(metaJson[:])
	rel := map[string]interface{}{
		"refunds": metadataMap,
	}
	u := fmt.Sprintf("/refunds/%s", updatedRefund.ID)

	err := s.client.Call("PUT", u, rel, refund)

	return refund, err
}
