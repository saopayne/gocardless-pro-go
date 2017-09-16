package gocardless

import (
	"encoding/json"
	"fmt"
)

type MandatePdfService service

type MandatePdf struct {
	Url       string `json:"url,omitempty"`
	ExpiresAt string `json:"expires_at,omitempty"`
}

type MandatePdfCreateRequest struct {
	AccountNumber     string            `json:"account_number,omitempty"`
	AccountHolderName string            `json:"account_holder_name,omitempty"`
	BankCode          string            `json:"bank_code,omitempty"`
	Bic               string            `json:"bic,omitempty"`
	BranchCode        string            `json:"branch_code,omitempty"`
	CountryCode       string            `json:"country_code,omitempty"`
	Iban              string            `json:"iban,omitempty"`
	MandateReference  string            `json:"mandate_reference,omitempty"`
	Scheme            string            `json:"scheme,omitempty"`
	Metadata          map[string]string `json:"metadata,omitempty"`
	Links             map[string]string `json:"links,omitempty"`
}

// Create creates a new mandatePdf
func (s *MandatePdfService) CreateMandatePdf(mandatePdfReq *MandatePdfCreateRequest) (*MandatePdf, error) {
	u := fmt.Sprintf("/mandate_pdfs")
	mandatePdf := &MandatePdf{}
	rel := map[string]interface{}{
		"mandate_pdfs": mandatePdfReq,
	}
	linksString, _ := json.Marshal(rel)
	linkJson := string(linksString[:])
	fmt.Print(linkJson)
	err := s.client.Call("POST", u, rel, mandatePdf)

	return mandatePdf, err
}
