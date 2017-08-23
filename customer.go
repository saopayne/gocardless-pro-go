package gocardless_pro_go

import (
	"fmt"
	"net/url"
)

type CustomerService service

type Customer struct {
	ID             			int            		`json:"id,omitempty"`
	CreatedAt      			string         		`json:"createdAt,omitempty"`
	UpdatedAt      			string         		`json:"updatedAt,omitempty"`
	AddressLine1			string				`json:"address_line_1,omitempty"`
	AddressLine2			string				`json:"address_line_2,omitempty"`
	AddressLine3			string				`json:"address_line_3,omitempty"`
	CompanyName				string				`json:"company_name,omitempty"`
	CountryCode				string				`json:"country_code,omitempty"`
	Email          			string         		`json:"email,omitempty"`
	FamilyName      		string         		`json:"family_name,omitempty"`
	GivenName       		string         		`json:"given_name,omitempty"`
	City					string				`json:"city,omitempty"`
	Language				string				`json:"language,omitempty"`
	PostalCode				string				`json:"postal_code,omitempty"`
	Region					string				`json:"region,omitempty"`
	SwedishIdentityNumber	string				`json:"swedish_identity_number,omitempty"`
	Metadata       			map[string]string   `json:"metadata,omitempty"`
}

// CustomerList is a list object for customers.
type CustomerList struct {
	Meta   ListMeta
	Values []Customer `json:"data"`
}

// Create creates a new customer
// For more details see https://developers.paystack.co/v1.0/reference#create-customer
func (s *CustomerService) Create(customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("/customer")
	cust := &Customer{}
	err := s.client.Call("POST", u, customer, cust)

	return cust, err
}

// Update updates a customer's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-customer
func (s *CustomerService) Update(customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("customer/%d", customer.ID)
	cust := &Customer{}
	err := s.client.Call("PUT", u, customer, cust)

	return cust, err
}

// Get returns the details of a customer.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-customer
func (s *CustomerService) Get(id int) (*Customer, error) {
	u := fmt.Sprintf("/customer/%d", id)
	cust := &Customer{}
	err := s.client.Call("GET", u, nil, cust)

	return cust, err
}

// List returns a list of customers.
// For more details see https://developers.paystack.co/v1.0/reference#list-customers
func (s *CustomerService) List() (*CustomerList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of customers
// For more details see https://developers.paystack.co/v1.0/reference#list-customers
func (s *CustomerService) ListN(count, offset int) (*CustomerList, error) {
	u := paginateURL("/customer", count, offset)
	cust := &CustomerList{}
	err := s.client.Call("GET", u, nil, cust)
	return cust, err
}

// SetRiskAction can be used to either whitelist or blacklist a customer
// For more details see https://developers.paystack.co/v1.0/reference#whiteblacklist-customer
func (s *CustomerService) SetRiskAction(customerCode, riskAction string) (*Customer, error) {
	params := url.Values{}
	params.Add("customer", customerCode)
	params.Add("risk_action", riskAction)
	cust := &Customer{}
	err := s.client.Call("POST", "/customer/set_risk_action", params, cust)

	return cust, err
}

// DeactivateAuthorization deactivates an authorization
// For more details see https://developers.paystack.co/v1.0/reference#deactivate-authorization
func (s *CustomerService) DeactivateAuthorization(authorizationCode string) (*Response, error) {
	params := url.Values{}
	params.Add("authorization_code", authorizationCode)

	resp := &Response{}
	err := s.client.Call("POST", "/customer/deactivate_authorization", params, resp)

	return resp, err
}

