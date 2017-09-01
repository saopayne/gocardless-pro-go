package main

import (
	"fmt"
	"net/url"
)

type RedirectFlowService service

type RedirectFlow struct {
	ID                 string        `json:"id,omitempty"`
	Description        string        `json:"description,omitempty"`
	CreatedAt          string        `json:"created_at,omitempty"`
	Scheme             string        `json:"scheme,omitempty"`
	RedirectUrl        string        `json:"redirect_url,omitempty"`
	SessionToken       string        `json:"session_token,omitempty"`
	SuccessRedirectUrl string        `json:"success_redirect_url,omitempty"`
	Links              []MandateLink `json:"links,omitempty"`
}

type RedirectFlowCreateRequest struct {
	PrefilledCustomer  PrefilledCustomer `json:"prefilled_customer,omitempty"`
	Description        string            `json:"description,omitempty"`
	Scheme             string            `json:"scheme,omitempty"`
	RedirectUrl        string            `json:"redirect_url,omitempty"`
	SessionToken       string            `json:"session_token,omitempty"`
	SuccessRedirectUrl string            `json:"success_redirect_url,omitempty"`
	Links              []string          `json:"links,omitempty"`
}

type PrefilledCustomer struct {
	ID                    int               `json:"id,omitempty"`
	CreatedAt             string            `json:"createdAt,omitempty"`
	UpdatedAt             string            `json:"updatedAt,omitempty"`
	AddressLine1          string            `json:"address_line1,omitempty"`
	AddressLine2          string            `json:"address_line2,omitempty"`
	AddressLine3          string            `json:"address_line3,omitempty"`
	CompanyName           string            `json:"company_name,omitempty"`
	CountryCode           string            `json:"country_code,omitempty"`
	Email                 string            `json:"email,omitempty"`
	FamilyName            string            `json:"family_name,omitempty"`
	GivenName             string            `json:"given_name,omitempty"`
	City                  string            `json:"city,omitempty"`
	Language              string            `json:"language,omitempty"`
	PostalCode            string            `json:"postal_code,omitempty"`
	Region                string            `json:"region,omitempty"`
	SwedishIdentityNumber string            `json:"swedish_identity_number,omitempty"`
	Metadata              map[string]string `json:"metadata,omitempty"`
}

func (s *RedirectFlowService) Create(redirectFlow *RedirectFlow) (*RedirectFlow, error) {
	u := fmt.Sprintf("/redirect_flows")
	rFlow := &RedirectFlow{}
	err := s.client.Call("POST", u, redirectFlow, rFlow)

	return rFlow, err
}

func (s *RedirectFlowService) GetRedirectFlow(id string) (*RedirectFlow, error) {
	u := fmt.Sprintf("/redirect_flows/%s", id)
	rFlow := &RedirectFlow{}
	err := s.client.Call("GET", u, nil, rFlow)

	return rFlow, err
}

func (s *PaymentService) CompleteRedirectFlow(rFlowToComplete *RedirectFlow, sessionToken string) (*Response, error) {
	params := url.Values{}
	params.Add("session_token", sessionToken)
	u := fmt.Sprintf("/redirect_flows/%s/actions/complete", rFlowToComplete.ID)
	resp := &Response{}
	err := s.client.Call("POST", u, params, resp)

	return resp, err
}
