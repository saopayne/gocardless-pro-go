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
	AddressLine1			string				`json:"address_line1,omitempty"`
	AddressLine2			string				`json:"address_line2,omitempty"`
	AddressLine3			string				`json:"address_line3,omitempty"`
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

type CustomerCreateRequest struct {
	ID             			int            		`json:"id,omitempty"`
	AddressLine1			string				`json:"address_line1,omitempty"`
	AddressLine2			string				`json:"address_line2,omitempty"`
	AddressLine3			string				`json:"address_line3,omitempty"`
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
}


type CustomerListRequest struct {
	CreatedAt 	CreatedAt		`json:"created_at,omitempty"`
	Limit		int				`json:"limit,omitempty"`
	Before		string			`json:"before,omitempty"`
	After		string			`json:"after,omitempty"`
}

type CustomerUpdateRequest struct {
	Name      			string  					`json:"name,omitempty"`
	Region          	string 						`json:"region,omitempty"`
	PostalCode			string 						`json:"postal_code,omitempty"`
	City				string 						`json:"city,omitempty"`
	AddressLine1		string						`json:"address_line1,omitempty"`
	AddressLine2		string						`json:"address_line2,omitempty"`
	AddressLine3		string						`json:"address_line3,omitempty"`
	CountryCode			string						`json:"country_code,omitempty"`
	Identity			string						`json:"identity,omitempty"`
	Links				[]string					`json:"links,omitempty"`
}

// CustomerList is a list object for customers.
type CustomerList struct {
	Meta   ListMeta
	Values []Customer `json:"data"`
}

// Create creates a new customer
// For more details see https://developer.gocardless.com/api-reference/#customers-create-a-customer
func (s *CustomerService) Create(customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("/customers")
	cust := &Customer{}
	err := s.client.Call("POST", u, customer, cust)

	return cust, err
}

// Update updates a customer's properties.
// For more details see https://developer.gocardless.com/api-reference/#customers-update-a-customer
func (s *CustomerService) Update(customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("customers/%d", customer.ID)
	cust := &Customer{}
	err := s.client.Call("PUT", u, customer, cust)

	return cust, err
}

// Get returns the details of a customer.
// For more details https://developer.gocardless.com/api-reference/#customers-get-a-single-customer
func (s *CustomerService) Get(id int) (*Customer, error) {
	u := fmt.Sprintf("/customers/%d", id)
	cust := &Customer{}
	err := s.client.Call("GET", u, nil, cust)

	return cust, err
}

func (s *CustomerService) List(req *CustomerListRequest) (*CustomerList, error) {
	return s.ListN(100,10, req)
}

// ListN returns a list of customers
func (s *CustomerService) ListN(count, offset int,req *CustomerListRequest) (*CustomerList, error) {
	params := url.Values{}
	params.Add("after", req.After)
	params.Add("before", req.Before)
	params.Add("created_at[gt]", req.CreatedAt.Gt)
	params.Add("created_at[gte]", req.CreatedAt.Gte)
	params.Add("created_at[lt]", req.CreatedAt.Lt)
	params.Add("created_at[lte]", req.CreatedAt.Lte)
	params.Add("limit", string(req.Limit))
	u := paginateURL("/customers", count, offset)
	cust := &CustomerList{}
	err := s.client.Call("GET", u, params, cust)
	return cust, err
}
