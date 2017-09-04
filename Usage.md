#### This is comprehensive usage document for this library

##### 1. Customers

* Request
```go

client := NewClient(apiKey, nil)
cust := &Customer{
    FamilyName:   "Oyewale",
    GivenName:    "Ademola",
    Email:        "user123@gmail.com",
    PostalCode:   "E2 8DP",
    CountryCode:  "GB",
    City:         "Lagos",
    AddressLine1: "Just somewhere on Earth",
    AddressLine2: "Another place on Earth",
    AddressLine3: "Just the third address to justify things",
    Language:     "en",
}

// create the customer
client.LoggingEnabled = true
customer, err := client.Customer.Create(cust)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while creating a customer is :%s", err.Error())
}
fmt.Sprintf("The customer created is: %s ", string(customer.Email))

// Get customer by ID
customer, err = client.Customer.Get(customer.ID)
if err != nil {
    fmt.Sprintf("The error while getting a customer is :%s", err.Error())
}

fmt.Sprintf("The customer retrieved with ID: %d is : %s", customer.ID, customer.Email)
```
* Response
```json

    2017/09/02 20:50:14 Requesting GET api-sandbox.gocardless.com/customers/
    {"customers":{"id":"CU00029EH0FGEB","created_at":"2017-09-02T19:50:14.085Z","email":"user123@gmail.com","given_name":"Ademola","family_name":"Oyewale","company_name":null,"address_line1":"Just somewhere on Earth","address_line2":"Another place on Earth","address_line3":"Just the third address to justify things","city":"Lagos","region":null,"postal_code":"E2 8DP","country_code":"GB","language":"en","swedish_identity_number":null,"metadata":{}}}The call method is being called

    2017/09/02 20:50:15 Completed in 557.343235ms
    RESPONSE https://api-sandbox.gocardless.com/customers/
    {"customers":[{"id":"CU00029EH0FGEB","created_at":"2017-09-02T19:50:14.085Z","email":"user123@gmail.com","given_name":"Ademola","family_name":"Oyewale","company_name":null,"address_line1":"Just somewhere on Earth","address_line2":"Another place on Earth","address_line3":"Just the third address to justify things","city":"Lagos","region":null,"postal_code":"E2 8DP","country_code":"GB","language":"en","swedish_identity_number":null,"metadata":{}},{"id":"CU00029BQFAY9V","created_at":"2017-09-01T19:22:25.107Z","email":"user123@gmail.com","given_name":"Ademola","family_name":"Oyewale","company_name":null,"address_line1":"Just somewhere on Earth","address_line2":"Another place on Earth","address_line3":"Just the third address to justify things","city":"Lagos","region":null,"postal_code":"E2 8DP","country_code":"GB","language":"en","swedish_identity_number":null,"metadata":{}}],"meta":{"cursors":{"before":null,"after":null},"limit":50}}
```

##### 2. Creditors

* Request
```go
client := NewClient(apiKey, nil)

cust := &Creditor{
    Name:   "Oyewale",
    PostalCode:   "E2 8DP",
    CountryCode:  "GB",
    City:         "Lagos",
    AddressLine1: "Just somewhere on Earth",
    AddressLine2: "Another place on Earth",
    AddressLine3: "Just the third address to justify things",
}

// create the customer
client.LoggingEnabled = true
customer, err := client.Creditor.CreateCreditor(cust)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while creating a creditor is :%s", err.Error())
}
fmt.Sprintf("The creditor created is: %s ", string(customer.Name))

// Get customer by ID
customer, err = client.Creditor.GetCreditor(customer.Id)
if err != nil {
    fmt.Sprintf("The error while getting a creditor is :%s", err.Error())
}
fmt.Sprintf("The creditor retrieved with ID: %d is : %s", customer.Id, customer.Name)

creditorUpdated := &Creditor{
    Id:		customer.Id,
    Name:   	"Oyewale Sao",
    PostalCode:   "E2 8DP",
    CountryCode:  "GB",
    City:         "Lagos",
    AddressLine1: "Just somewhere on Earth",
    AddressLine2: "Another place on Earth",
    AddressLine3: "Just the third address to justify things",
}

// create the customer
client.LoggingEnabled = true
customer, err = client.Creditor.UpdateCreditor(creditorUpdated)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while updating a creditor is :%s", err.Error())
}

fmt.Sprintf("The creditor updated is: %s ", string(customer.Name))
```

##### 3. Customer Bank Account

* Making the Request

```go
client := NewClient(apiKey, nil)

acct := &CustomerBankAccountCreateRequest{
    BankCode:      "Oyewale",
    Currency:      "PND",
    BranchCode:    "LEI",
    AccountNumber: "03434",
    CountryCode:   "GB",
}

// create a customer bank account
client.LoggingEnabled = true
account, err := client.CustomerBankAccount.CreateCustomerBankAccount(acct)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while creating a customer bank account is :%s", err.Error())
}

fmt.Sprintf("The customer bank account created is: %s ", account.BankName)

// Get customer bank Account by ID
account, err = client.CustomerBankAccount.GetCustomerBankAccount(account.Id)
if err != nil {
    fmt.Sprintf("The error while getting a customer bank account is :%s", err.Error())
}

fmt.Sprintf("The customer bank account retrieved with ID: %d is : %s", account.Id, account.BankName)

custBankAccountUpdate := &CustomerBankAccount{
    BankName:          "Oyewale",
    AccountHolderName: "Ademola",
    CountryCode:       "GB",
}

account, err = client.CustomerBankAccount.UpdateCustomerBankAccount(custBankAccountUpdate, make(map[string]string))
if err != nil {
    // do something with error
    fmt.Sprintf("The error while updating a customer bank account is :%s", err.Error())
}

fmt.Sprintf("The customer bank account updated is: %s ", account.BankName)
```

##### 4. Creditor Bank Account

* Making request

```go
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
```

##### 5. Bank Details Lookup

* Making the request
```go
client := NewClient(apiKey, nil)

bankLookup := &BankDetailsLookupRequest{
    AccountNumber: "55779911",
    BranchCode: "200000",
    CountryCode:       "GB",
}
account, err := client.BankDetailsLookup.Lookup(bankLookup)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while looking up details of a bank is :%s", err.Error())
}
fmt.Sprintf("The customer bank account updated is: %s ", account)
```

* Response
```json

    2017/09/02 22:03:42 Requesting POST api-sandbox.gocardless.com/bank_details_lookups
    2017/09/02 22:03:46 Completed in 4.495158108s
    RESPONSE https://api-sandbox.gocardless.com/bank_details_lookups
    {"bank_details_lookups":{"bank_name":"BARCLAYS BANK PLC","available_debit_schemes":["bacs"],"bic":"BUKBGB22XXX"}}

```

##### 6. Events

* Making Request

```go
client := NewClient(apiKey, nil)
event, err := client.Event.GetEvent("EV123")
if err != nil {
    fmt.Sprintf("The error while getting an event is :%s", err.Error())
}
fmt.Sprintf("The event retrieved with ID: %d is : %s", event.ID, event.Details)
```

##### 7. Mandate
* Making Request

```go
//second param is an optional http client, allowing overriding of the HTTP client to use.
//This is useful if you're running in a Google AppEngine environment
//where the http.DefaultClient is not available.
client := NewClient(apiKey, nil)
linksMap := map[string]string{
    "customer_bank_account": "XXXX",
}
linksMapString, _ := json.Marshal(linksMap)
linksJson := string(linksMapString[:])
rel := map[string]string{
    "links": linksJson,
}
linksString, _ := json.Marshal(rel)
linkJson := string(linksString[:])
linkJson,_ = strconv.Unquote(linkJson)

mandateReq := &MandateCreateRequest{
    Scheme: "bacs",
    Links: linkJson,
    CustomerBankAccount: "MD123",
}
// create a mandate
client.LoggingEnabled = true
mandate, err := client.Mandate.CreateMandate(mandateReq)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while creating a mandate is :%s", err.Error())
}
fmt.Sprintf("The mandate created is: %s ", mandate.Scheme)

// Get mandate by ID
mandate, err = client.Mandate.GetMandate("MD123")
if err != nil {
    fmt.Sprintf("The error while getting a mandate is :%s", err.Error())
}
fmt.Sprintf("The event retrieved with ID: %d is : %s", mandate.ID, mandate.Scheme)

mandateUpdateReq := &Mandate{
    Reference: "New reference",
    Scheme: "bacs",
}

// update a mandate
client.LoggingEnabled = true
mandateToUpdate, err := client.Mandate.UpdateMandate(mandateUpdateReq, make(map[string]string))
if err != nil {
    // do something with error
    fmt.Sprintf("The error while updating a mandate is :%s", err.Error())
}
fmt.Sprintf("The mandate updated is: %s ", mandateToUpdate.Reference)

// cancel a mandate
mandateCancelReq := &Mandate{
    Reference: "New reference",
    Scheme: "bacs",
    ID: "MD123",
}
client.LoggingEnabled = true
mandateToCancel, err := client.Mandate.CancelMandate(mandateCancelReq, make(map[string]string))
if err != nil {
    // do something with error
    fmt.Sprintf("The error while canceling a mandate is :%s", err.Error())
}
fmt.Sprintf("The mandate canceled returned the response: %s ", mandateToCancel)

// reinstate a mandate
mandateReinReq := &Mandate{
    Reference: "New reference",
    Scheme: "bacs",
    ID: "MD123",
}
client.LoggingEnabled = true
mandateToRein, err := client.Mandate.ReinstateMandate(mandateReinReq, make(map[string]string))
if err != nil {
    // do something with error
    fmt.Sprintf("The error while reinstating a mandate is :%s", err.Error())
}
fmt.Sprintf("The mandate reinstated returned a response: %s ", mandateToRein)

```

##### 8. Mandate  PDFs

* Making the request

```go
client := NewClient(apiKey, nil)
linkMap := make(map[string]string)
linkMap["mandate"] = "MD123"

mandateReq := &MandatePdfCreateRequest{
    Links: linkMap,
}
// create a mandate
client.LoggingEnabled = true
mandate, err := client.MandatePdf.CreateMandatePdf(mandateReq)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while creating a mandate pdf is :%s", err.Error())
}
fmt.Sprintf("The mandate pdf created is: %s ", mandate.Url)
```

##### 9. Payouts

* Making the request

```go
// Get Payout by ID
payout, err := client.Payout.GetPayout("PO123")
if err != nil {
    fmt.Sprintf("The error while getting a payout is :%s", err.Error())
}
fmt.Sprintf("The payout retrieved with ID: %d is : %s", payout.ID, payout.Reference)
```