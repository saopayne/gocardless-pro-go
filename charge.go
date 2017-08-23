package gocardless_pro_go

// ChargeService handles operations related to bulk charges
// For more details see https://www.flutterwave.com/documentation/card-payments/
type ChargeService service

type Bank struct {
	Code          string `json:"bank,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
}

type Card struct {
	Number            string `json:"card_number,omitempty"`
	CVV               string `json:"card_cvc,omitempty"`
	ExpiryMonth      string `json:"expiry_month,omitempty"`
	ExpiryYear        string `json:"expiry_year,omitempty"`
	AddressLine1      string `json:"address_line1,omitempty"`
	AddressLine2      string `json:"address_line2,omitempty"`
	AddressLine3      string `json:"address_line3,omitempty"`
	AddressCountry    string `json:"address_country,omitempty"`
	AddressPostalCode string `json:"address_postal_code,omitempty"`
	Country           string `json:"country,omitempty"`
}

// ChargeRequest represents a Flutterwave charge request
type ChargeRequest struct {
	Amount            float32  	`json:"amount,omitempty"`
	AuthModel		  string 	`json:"authmodel,omitempty"`
	CardNumber        string 	`json:"cardno,omitempty"`
	CVV               string 	`json:"cvv,omitempty"`
	Currency		  string 	`json:"currency,omitempty"`
	CustomerId		  Customer 	`json:"custid,omitempty"`
	ExpiryMonth       string 	`json:"expirymonth,omitempty"`
	ExpiryYear        string 	`json:"expiryyear,omitempty"`
	Country           string 	`json:"country,omitempty"`
	MerchantId        Merchant  `json:"merchantid,omitempty"`
	BVN				  string 	`json:"bvn,omitempty"`
	CardType		  Card		`json:"cardtype,omitempty"`
	Pin               string 	`json:"pin,omitempty"`
	Narration		  string	`json:"narration,omitempty"`
	ResponseUrl		  string	`json:"responseurl,omitempty"`
	Metadata          Metadata 	`json:"metadata,omitempty"`
}

// For calls from the previous endpoint that return a responsecode:02, they need to be validated with this call.
type ValidateRequest struct {
	OTP							string		`json:"otp,omitempty"`
	OTPTransactionIdentifier 	string 		`json:"otptransactionidentifier,omitempty"`
	MerchantId					Merchant	`json:"merchantid,omitempty"`
}

// Create submits a charge request using card details or bank details or authorization code
// For more details see https://www.flutterwave.com/documentation/card-payments/
func (s *ChargeService) Create(req *ChargeRequest) (*Response, error) {
	resp := &Response{}
	err := s.client.Call("POST", "/pay", req, resp)

	return resp, err
}

// For calls from the previous endpoint that return a responsecode:02, they need to be validated with this call.
// Transactions that need validation are those where the authmodel is PIN, BVN or RANDOM_DEBIT.
// In the case of RANDOM_DEBIT, OTP would be an amount like 1.20, which the customer should provide in full like that.
// http://staging1flutterwave.co:8080/pwc/rest/card/mvva/pay/validate
func (s *ChargeService) Validate(req *ValidateRequest) (*Response, error) {
	resp := &Response{}
	err := s.client.Call("POST", "/validate", req, resp)

	return resp, err
}

