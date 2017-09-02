package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/mitchellh/mapstructure"
	"strings"
)

const (
	// library version
	version = "1.1"
	// defaultHTTPTimeout is the default timeout on the http client
	defaultHTTPTimeout = 40 * time.Second
	baseURL            = "https://api-sandbox.gocardless.com"
	// User agent used when communicating with the Gocardless API.
	userAgent            = "gocardless-webhook-service/" + version
	goCardlessApiVersion = "2015-07-06"
	acceptJsonType       = "application/json"
)

func main() {

	apiKey := "sandbox_o55p5OowBX59Rd8aDR7c_25LQdBTHRaACeVnqj0o"

	 //second param is an optional http client, allowing overriding of the HTTP client to use.
	 //This is useful if you're running in a Google AppEngine environment
	 //where the http.DefaultClient is not available.
	client := NewClient(apiKey, nil)

	acct := &CreditorBankAccountCreateRequest{
		BankCode:      "Oyewale",
		Currency:      "PND",
		BranchCode:    "LEI",
		AccountNumber: "03434",
		CountryCode:   "GB",
	}
	// create a creditor bank account
	client.LoggingEnabled = true
	account, err := client.CreditorBankAccount.CreateCreditorBankAccount(acct)
	if err != nil {
		// do something with error
		fmt.Sprintf("The error while creating a creditor bank account is :%s", err.Error())
	}
	fmt.Sprintf("The creditor bank account created is: %s ", account.BankName)

	// Get creditor bank Account by ID
	account, err = client.CreditorBankAccount.GetCreditorBankAccount(account.Id)
	if err != nil {
		fmt.Sprintf("The error while getting a creditor bank account is :%s", err.Error())
	}
	fmt.Sprintf("The creditor bank account retrieved with ID: %d is : %s", account.Id, account.BankName)

	//custBankAccountUpdate := &CreditorBankAccount{
	//	BankName:          "Oyewale",
	//	AccountHolderName: "Ademola",
	//	CountryCode:       "GB",
	//}
	//account, err = client.CreditorBankAccount.Update(custBankAccountUpdate, make(map[string]string))
	//if err != nil {
	//	// do something with error
	//	fmt.Sprintf("The error while updating a customer bank account is :%s", err.Error())
	//}
	//fmt.Sprintf("The customer bank account updated is: %s ", account.BankName)

}

type service struct {
	client *Client
}

// Client manages communication with the GoCardless API
type Client struct {
	common service      // Reuse a single struct instead of allocating one for each service on the heap.
	client *http.Client // HTTP client used to communicate with the API.
	// the API Key used to authenticate all GoCardless API requests
	key     string
	secret  string // don't know if this is necessary though
	baseURL *url.URL
	logger  Logger

	BankDetailsLookup   *BankDetailsLookupService
	Creditor            *CreditorService
	CreditorBankAccount *CreditorBankAccountService
	Customer            *CustomerService
	CustomerBankAccount *CustomerBankAccountService
	Event               *EventService
	Mandate             *MandateService
	MandatePdf          *MandatePdfService
	Payout              *PayoutService
	Payment             *PaymentService
	RedirectFlow        *RedirectFlowService
	Refund              *RefundService
	Subscription        *SubscriptionService
	LoggingEnabled      bool
	Logger              Logger
}

type Logger interface {
	Printf(format string, v ...interface{})
}

type Metadata map[string]interface{}

// Response represents arbitrary response data
type Response map[string]interface{}

// RequestValues aliased to url.Values as a workaround
type RequestValues url.Values

func (v RequestValues) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 3)
	for k, val := range v {
		m[k] = val[0]
	}
	return json.Marshal(m)
}

type ListMeta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

// NewClient creates a new GoCardless API client with the given API key
// and HTTP client, allowing overriding of the HTTP client to use.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func NewClient(key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultHTTPTimeout}
	}

	u, _ := url.Parse(baseURL)
	c := &Client{
		client:         httpClient,
		key:            key,
		baseURL:        u,
		LoggingEnabled: true,
		Logger:         log.New(os.Stderr, "", log.LstdFlags),
	}

	c.common.client = c

	c.BankDetailsLookup = (*BankDetailsLookupService)(&c.common)
	c.Customer = (*CustomerService)(&c.common)
	c.CustomerBankAccount = (*CustomerBankAccountService)(&c.common)
	c.Creditor = (*CreditorService)(&c.common)
	c.CreditorBankAccount = (*CreditorBankAccountService)(&c.common)
	c.BankDetailsLookup = (*BankDetailsLookupService)(&c.common)
	c.Event = (*EventService)(&c.common)
	c.Mandate = (*MandateService)(&c.common)
	c.MandatePdf = (*MandatePdfService)(&c.common)
	c.Payout = (*PayoutService)(&c.common)
	c.Payment = (*PaymentService)(&c.common)
	c.RedirectFlow = (*RedirectFlowService)(&c.common)
	c.Refund = (*RefundService)(&c.common)
	c.Subscription = (*SubscriptionService)(&c.common)

	return c
}

func (c *Client) Call(method string, path string, body, v interface{}) error {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return err
		}
	}
	u, _ := c.baseURL.Parse(path)
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		if c.LoggingEnabled {
			c.Logger.Printf("Cannot create GoCardless request: %v\n", err)
		}
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Authorization", "Bearer "+c.key)
	// this header sets the api version
	req.Header.Set("GoCardless-Version", goCardlessApiVersion)
	req.Header.Set("Accept", acceptJsonType)

	if ua := req.Header.Get("User-Agent"); ua == "" {
		req.Header.Set("User-Agent", userAgent)
	} else {
		req.Header.Set("User-Agent", userAgent+" "+ua)
	}

	if c.LoggingEnabled {
		c.Logger.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
	}

	start := time.Now()

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if c.LoggingEnabled {
		c.Logger.Printf("Completed in %v\n", time.Since(start))
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if c.LoggingEnabled {
			c.Logger.Printf("Request to GoCardless failed: %v\n", err)
		}
		return err
	}

	var respMap Response
	json.Unmarshal(respBody, &respMap)
	fmt.Printf("RESPONSE %s \n %+v", u, string(respBody[:]))

	if strings.Contains(resp.Status, "20") && resp.StatusCode >= 400 {
		if c.LoggingEnabled {
			c.Logger.Printf("GoCardless error: %v\n", err)
		}
		return responseToError(resp, respMap)
	}
	return checkResponse(respMap, v)
}

func (c *Client) ResolveCardBIN(bin int) (*Response, error) {
	u := fmt.Sprintf("/decision/bin/%d", bin)
	resp := &Response{}
	err := c.Call("GET", u, nil, resp)

	return resp, err
}

func (c *Client) CheckBalance(bin int) (*Response, error) {
	resp := &Response{}
	err := c.Call("GET", "balance", nil, resp)
	return resp, err
}

func (c *Client) GetSessionTimeout() (*Response, error) {
	resp := &Response{}
	err := c.Call("GET", "/integration/payment_session_timeout", nil, resp)
	return resp, err
}

func (c *Client) UpdateSessionTimeout(timeout int) (*Response, error) {
	data := url.Values{}
	data.Add("timeout", string(timeout))
	resp := &Response{}
	u := "/integration/payment_session_timeout"
	err := c.Call("PUT", u, data, resp)
	return resp, err
}

// INTERNALS
func paginateURL(path string, count, offset int) string {
	return fmt.Sprintf("%s?perPage=%d&page=%d", path, count, offset)
}

func mapstruct(data interface{}, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err = decoder.Decode(data); err != nil {
		return err
	}
	return nil
}

func getTestKey() string {
	key := os.Getenv("GOCARDLESS-KEY")

	if len(key) == 0 {
		panic("GOCARDLESS environment variable is not set\n")
	}

	return key
}

func responseToError(resp *http.Response, respMap Response) error {
	err := &Error{
		HTTPStatusCode: resp.StatusCode,
		Message:        respMap["message"].(string),
		URL:            resp.Request.URL,
	}
	if errorDetails, ok := respMap["errors"]; ok {
		err.Errors = errorDetails.(map[string][]interface{})
	}
	return err
}

func checkResponse(respMap Response, v interface{}) error {
	if data, ok := respMap["data"]; ok {
		switch t := respMap["data"].(type) {
		case map[string]interface{}:
			return mapstruct(data, v)
		default:
			_ = t
			return mapstruct(respMap, v)
		}
	}
	// response data does not contain data node, return anyways
	return mapstruct(respMap, v)
}
