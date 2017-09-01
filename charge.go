package main

type ChargeService service

type Bank struct {
	Code          string `json:"bank,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
}

type Card struct {
	Number            string `json:"card_number,omitempty"`
	CVV               string `json:"card_cvc,omitempty"`
	ExpiryMonth       string `json:"expiry_month,omitempty"`
	ExpiryYear        string `json:"expiry_year,omitempty"`
	AddressLine1      string `json:"address_line1,omitempty"`
	AddressLine2      string `json:"address_line2,omitempty"`
	AddressLine3      string `json:"address_line3,omitempty"`
	AddressCountry    string `json:"address_country,omitempty"`
	AddressPostalCode string `json:"address_postal_code,omitempty"`
	Country           string `json:"country,omitempty"`
}

type ChargeRequest struct {
	Amount      float32  `json:"amount,omitempty"`
	AuthModel   string   `json:"authmodel,omitempty"`
	CardNumber  string   `json:"cardno,omitempty"`
	CVV         string   `json:"cvv,omitempty"`
	Currency    string   `json:"currency,omitempty"`
	CustomerId  Customer `json:"custid,omitempty"`
	ExpiryMonth string   `json:"expirymonth,omitempty"`
	ExpiryYear  string   `json:"expiryyear,omitempty"`
	Country     string   `json:"country,omitempty"`
	BVN         string   `json:"bvn,omitempty"`
	CardType    Card     `json:"cardtype,omitempty"`
	Pin         string   `json:"pin,omitempty"`
	Narration   string   `json:"narration,omitempty"`
	ResponseUrl string   `json:"responseurl,omitempty"`
	Metadata    Metadata `json:"metadata,omitempty"`
}

type ValidateRequest struct {
	OTP                      string `json:"otp,omitempty"`
	OTPTransactionIdentifier string `json:"otptransactionidentifier,omitempty"`
}

func (s *ChargeService) Create(req *ChargeRequest) (*Response, error) {
	resp := &Response{}
	err := s.client.Call("POST", "/pay", req, resp)

	return resp, err
}

func (s *ChargeService) Validate(req *ValidateRequest) (*Response, error) {
	resp := &Response{}
	err := s.client.Call("POST", "/validate", req, resp)

	return resp, err
}
