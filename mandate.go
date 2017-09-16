package gocardless

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type MandateService service

type Mandate struct {
	ID                      string            `json:"id,omitempty"`
	CreatedAt               string            `json:"created_at,omitempty"`
	Reference               string            `json:"reference,omitempty"`
	Scheme                  string            `json:"scheme,omitempty"`
	Status                  Status            `json:"status,omitempty"`
	PaymentsRequireApproval bool              `json:"payments_require_approval,omitempty"`
	NextPossibleChargeDate  string            `json:"next_possible_charge_date,omitempty"`
	Links                   []MandateLink     `json:"links,omitempty"`
	Metadata                map[string]string `json:"metadata,omitempty"`
}

type Status struct {
	Status string `json:"status,omitempty"`
}

type MandateListRequest struct {
	CreatedAt           CreatedAt `json:"created_at,omitempty"`
	Limit               int       `json:"limit,omitempty"`
	Before              string    `json:"before,omitempty"`
	After               string    `json:"after,omitempty"`
	Creditor            string    `json:"creditor,omitempty"`
	Customer            string    `json:"customer,omitempty"`
	CustomerBankAccount string    `json:"customer_bank_account,omitempty"`
	Reference           string    `json:"reference,omitempty"`
	Status              string    `json:"status,omitempty"`
}

// CustomerBankAccountList is a list object for customer bank accounts.
type MandateList struct {
	Meta   ListMeta
	Values []CustomerBankAccount `json:"data"`
}

type MandateCreateRequest struct {
	Metadata            map[string]string `json:"metadata,omitempty"`
	Reference           string            `json:"reference,omitempty"`
	Scheme              string            `json:"scheme,omitempty"`
	AccountNumber       string            `json:"account_number,omitempty"`
	CustomerBankAccount string            `json:"customer_bank_account,omitempty"`
	Links               string            `json:"links,omitempty"`
}

type MandateCancelRequest struct {
	Identity string `json:"identity,omitempty"`
}

// Create creates a new mandate
func (s *MandateService) CreateMandate(mandateReq *MandateCreateRequest) (*Mandate, error) {
	u := fmt.Sprintf("/mandates")
	mandate := &Mandate{}
	rel := map[string]interface{}{
		"mandates": mandateReq,
	}

	err := s.client.Call("POST", u, rel, mandate)

	return mandate, err
}

// List returns a list of mandates
func (s *MandateService) ListNMandates(req *MandateListRequest) (*MandateList, error) {
	reqd, err := http.NewRequest("GET", "/mandates", nil)
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
	if req.Reference != "" { params.Add("reference", req.Reference) }
	if req.Status != "" { params.Add("status", req.Status) }
	if req.CustomerBankAccount != "" {params.Add("customer_bank_account", req.CustomerBankAccount)}
	if req.Customer != "" {params.Add("customer", req.Customer)}
	if req.Creditor != "" {params.Add("creditor", req.Creditor)}

	reqd.URL.RawQuery = params.Encode()
	path := reqd.URL.String()
	mandates := &MandateList{}
	err = s.client.Call("GET", path, nil, mandates)

	return mandates, err
}

func (s *MandateService) GetMandate(id string) (*Mandate, error) {
	u := fmt.Sprintf("/mandates/%s", id)
	mandate := &Mandate{}
	err := s.client.Call("GET", u, nil, mandate)

	return mandate, err
}

func (s *MandateService) UpdateMandate(updatedMandate *Mandate, metadata map[string]string) (*Mandate, error) {
	mandate := &Mandate{}
	metadataString, _ := json.Marshal(metadata)
	metaJson := string(metadataString[:])
	metadataMap := make(map[string]string)
	metadataMap["metadata"] = string(metaJson[:])
	rel := map[string]interface{}{
		"mandates": metadataMap,
	}
	u := fmt.Sprintf("mandates/%s", updatedMandate.ID)

	err := s.client.Call("PUT", u, rel, mandate)

	return mandate, err
}

func (s *MandateService) CancelMandate(mandateToCancel *Mandate, metadata map[string]string) (*Response, error) {
	resp := &Response{}
	metadataString, _ := json.Marshal(metadata)
	metaJson := string(metadataString[:])
	metadataMap := make(map[string]string)
	metadataMap["metadata"] = string(metaJson[:])
	rel := map[string]interface{}{
		"mandates": metadataMap,
	}
	u := fmt.Sprintf("/mandates/%s/actions/cancel", mandateToCancel.ID)
	err := s.client.Call("POST", u, rel, resp)

	return resp, err
}

func (s *MandateService) ReinstateMandate(mandateToReinstate *Mandate, metadata map[string]string) (*Response, error) {
	resp := &Response{}
	metadataString, _ := json.Marshal(metadata)
	metaJson := string(metadataString[:])
	metadataMap := make(map[string]string)
	metadataMap["metadata"] = string(metaJson[:])
	rel := map[string]interface{}{
		"mandates": metadataMap,
	}
	u := fmt.Sprintf("/mandates/%s/actions/reinstate", mandateToReinstate.ID)

	err := s.client.Call("POST", u, rel, resp)

	return resp, err
}
