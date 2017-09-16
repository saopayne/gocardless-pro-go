package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type CustomerService service

type Customer struct {
	ID                    string            `json:"id,omitempty"`
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

type CustomerCreateRequest struct {
	Customers struct {
		Customers Customer `json:"customers"`
	}
}

type CustomerListRequest struct {
	CreatedAt CreatedAt `json:"created_at,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	Before    string    `json:"before,omitempty"`
	After     string    `json:"after,omitempty"`
}

type CustomerUpdateRequest struct {
	Name         string   `json:"name,omitempty"`
	Region       string   `json:"region,omitempty"`
	PostalCode   string   `json:"postal_code,omitempty"`
	City         string   `json:"city,omitempty"`
	AddressLine1 string   `json:"address_line1,omitempty"`
	AddressLine2 string   `json:"address_line2,omitempty"`
	AddressLine3 string   `json:"address_line3,omitempty"`
	CountryCode  string   `json:"country_code,omitempty"`
	Identity     string   `json:"identity,omitempty"`
	Links        []string `json:"links,omitempty"`
}

// CustomerList is a list object for customers.
type CustomerList struct {
	Meta   ListMeta
	Values []Customer `json:"data"`
}

type CustomerEnvelope struct {
	Customers *Customer `json:"customers"`
}

// Create creates a new customer
// For more details see https://developer.gocardless.com/api-reference/#customers-create-a-customer
func (s *CustomerService) Create(customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("/customers")
	cust := &Customer{}
	rel := map[string]interface{}{
		"customers": customer,
	}
	custJson, _ := json.Marshal(rel)
	customerObject := string(custJson[:])
	fmt.Sprintf("Making request with the params %s", customerObject)
	err := s.client.Call("POST", u, rel, cust)

	return cust, err
}

// Update a customer's properties.
// For more details see https://developer.gocardless.com/api-reference/#customers-update-a-customer
func (s *CustomerService) Update(customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("/customers/%s", customer.ID)
	cust := &Customer{}
	rel := map[string]interface{}{
		"customers": customer,
	}
	//custReq := &CustomerEnvelope{Customers: customer}
	custJson, _ := json.Marshal(rel)

	customerObject := string(custJson[:])
	fmt.Println(customerObject)
	err := s.client.Call("PUT", u, rel, cust)
	return cust, err
}

// Get returns the details of a customer.
// For more details https://developer.gocardless.com/api-reference/#customers-get-a-single-customer
func (s *CustomerService) Get(id string) (*Customer, error) {
	u := fmt.Sprintf("/customers/%s", id)
	cust := &Customer{}
	err := s.client.Call("GET", u, nil, cust)

	return cust, err
}

// ListN returns a list of customers
// Further documentation can be found here: https://developer.gocardless.com/api-reference/#customers-list-customers
func (s *CustomerService) ListAllCustomers(req *CustomerListRequest) (*CustomerList, error) {

	reqd, err := http.NewRequest("GET", "/customers", nil)

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
	if req.CreatedAt.Lte != "" {params.Add("created_at[lte]", req.CreatedAt.Lte)}
	if req.Limit > 0 {params.Add("limit", string(req.Limit))}

	reqd.URL.RawQuery = params.Encode()

	path := reqd.URL.String()

	cust := &CustomerList{}
	err = s.client.Call("GET", path, nil, cust)

	return cust, err
}
