#### This is comprehensive usage document for this library

---
##### 1. Customers
---
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

// listing all customers
customerReq	:= &CustomerListRequest{
    Limit: 100,
}
// list all customers
_, err := client.Customer.ListAllCustomers(customerReq)
if err != nil {
    fmt.Sprintf("The error while getting list of customers  is :%s", err.Error())
}

```
* Response
```json

    2017/09/02 20:50:14 Requesting GET api-sandbox.gocardless.com/customers/
    {"customers":{"id":"CU00029EH0FGEB","created_at":"2017-09-02T19:50:14.085Z","email":"user123@gmail.com","given_name":"Ademola","family_name":"Oyewale","company_name":null,"address_line1":"Just somewhere on Earth","address_line2":"Another place on Earth","address_line3":"Just the third address to justify things","city":"Lagos","region":null,"postal_code":"E2 8DP","country_code":"GB","language":"en","swedish_identity_number":null,"metadata":{}}}The call method is being called

    2017/09/02 20:50:15 Completed in 557.343235ms
    RESPONSE https://api-sandbox.gocardless.com/customers/
    {"customers":[{"id":"CU00029EH0FGEB","created_at":"2017-09-02T19:50:14.085Z","email":"user123@gmail.com","given_name":"Ademola","family_name":"Oyewale","company_name":null,"address_line1":"Just somewhere on Earth","address_line2":"Another place on Earth","address_line3":"Just the third address to justify things","city":"Lagos","region":null,"postal_code":"E2 8DP","country_code":"GB","language":"en","swedish_identity_number":null,"metadata":{}},{"id":"CU00029BQFAY9V","created_at":"2017-09-01T19:22:25.107Z","email":"user123@gmail.com","given_name":"Ademola","family_name":"Oyewale","company_name":null,"address_line1":"Just somewhere on Earth","address_line2":"Another place on Earth","address_line3":"Just the third address to justify things","city":"Lagos","region":null,"postal_code":"E2 8DP","country_code":"GB","language":"en","swedish_identity_number":null,"metadata":{}}],"meta":{"cursors":{"before":null,"after":null},"limit":50}}
```
---
##### 2. Creditors
---

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

// create the creditor
client.LoggingEnabled = true
customer, err := client.Creditor.CreateCreditor(cust)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while creating a creditor is :%s", err.Error())
}
fmt.Sprintf("The creditor created is: %s ", string(customer.Name))

// Get creditor by ID
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

// updatex the creditor
client.LoggingEnabled = true
creditor, err = client.Creditor.UpdateCreditor(creditorUpdated)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while updating a creditor is :%s", err.Error())
}

fmt.Sprintf("The creditor updated is: %s ", string(creditor.Name))

//list all creditors
creditorListReq	:= &CreditorListRequest{
    Limit: 100,
}
// list all creditors
_, err := client.Creditor.ListCreditors(creditorListReq)
if err != nil {
    fmt.Sprintf("The error while getting list of creditors is :%s", err.Error())
}

```

* Response
```json
{ "creditors":[
    {"id":"CR00004YMS7RA5","created_at":"2017-09-01T12:45:10.932Z","name":"Tracchis","address_line1":null,"address_line2":null,"address_line3":null,"city":null,"region":null,"postal_code":null,"country_code":"GB",
    "logo_url":null,"scheme_identifiers":[{"name":"GoCardless Ltd","scheme":"bacs","reference":"275069","minimum_advance_notice":3,"currency":"GBP","address_line1":"338-346 Goswell Road","address_line2":null,"address_line3":null,
    "city":"London","region":null,"postal_code":"EC1V 7LQ","country_code":"GB","email":"help@gocardless.com","phone_number":"+44 20 7183 8674","can_specify_mandate_reference":false}],"verification_status":"successful",
    "links":{"default_gbp_payout_account":"BA0002378VB942","default_eur_payout_account":"BA0002378WJXWD","default_sek_payout_account":"BA0002378X955C"}}
    ],
    "meta":{"cursors":{"before":null,"after":null},"limit":1}
}
```

---
##### 3. Customer Bank Account
---
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

//list all customers bank accounts
customerBankAccountListReq	:= &CustomerBankListRequest{
    Limit: 100,
}
// list all customers bank accounts
_, err := client.CustomerBankAccount.ListCustomerBankAccounts(customerBankAccountListReq)
if err != nil {
    fmt.Sprintf("The error while getting list of customers bank accounts is :%s", err.Error())
}
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

//listing all creditors bank accounts
creditorBankAccountListReq	:= &CreditorBankAccountListRequest{
    Limit: 100,
}
// list all creditors bank accounts
_, err := client.CreditorBankAccount.ListCreditorBankAccounts(creditorBankAccountListReq)
if err != nil {
    fmt.Sprintf("The error while getting list of creditors bank accounts is :%s", err.Error())
}
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

//Getting a list of events
eventsListReq := &EventListRequest{
    Limit: 100,
}
// list all events
_, err := client.Event.ListEvents(eventsListReq)
if err != nil {
    fmt.Sprintf("The error while getting list of events  is :%s", err.Error())
}
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

// list all mandates
mandateListReq := &MandateListRequest{
    Limit: 100,
}
// list all mandates
_, err := client.Mandate.ListNMandates(mandateListReq)
if err != nil {
    fmt.Sprintf("The error while getting list of mandates  is :%s", err.Error())
}
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

##### 10. Redirect Flows

* Making the request

```go
client := NewClient(apiKey, nil)
rFlowCreateReq := &RedirectFlowCreateRequest{
    Description: "Wine vines",
    SessionToken: "SESS_wSs0uGYMISxzqOBq",
    PrefilledCustomer: PrefilledCustomer{
        GivenName: "Ademola",
        FamilyName: "Oyewale",
    },
    SuccessRedirectUrl: "https://wewee.ngrok.io/",
}
// create a redirecflow
client.LoggingEnabled = true
redirectFlow, err := client.RedirectFlow.Create(rFlowCreateReq)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while creating a redirect flow pdf is :%s", err.Error())
}
fmt.Sprintf("The redirect flow created has the description: %s ", redirectFlow.Description)

// Get Redirect Flow by ID
rFlow, err := client.RedirectFlow.GetRedirectFlow("RE123")
if err != nil {
    fmt.Sprintf("The error while getting a redirect flow is :%s", err.Error())
}
fmt.Sprintf("The redirecte flow retrieved with ID: %d is : %s", rFlow.ID, rFlow.Description)

rFlowCompleteReq := &RedirectFlowCompleteRequest{
    SessionToken: "SESS_wSs0uGYMISxzqOBq",
}
// complete a redirecflow
client.LoggingEnabled = true
redFlow, err := client.RedirectFlow.CompleteRedirectFlow("RE123", rFlowCompleteReq)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while completing a redirect flow pdf is :%s", err.Error())
}
fmt.Sprintf("The redirect flow completed has the description: %s ", redFlow)
```

* Sample Request Response
```json
 {"redirect_flows":{"id":"RE0000K6652W0ZPNPPRZ9MJKNHCPWCM7","description":"Wine vines","session_token":"SESS_wSs0uGYMISxzqOBq","scheme":null,"success_redirect_url":"https://wewee.ngrok.io/",
 "created_at":"2017-09-08T07:32:41.957Z","links":{"creditor":"CR00004YMS7RA5"},"redirect_url":"https://pay-sandbox.gocardless.com/flow/RE0000K6652W0ZPNPPRZ9MJKNHCPWCM7"}}
 RESPONSE https://api-sandbox.gocardless.com/redirect_flows/RE123
```

##### 11. Refunds

* Making the request

```go
refundCreateReq := &RefundCreateRequest{
    Amount: 100,
    TotalAmountConfirmation: "200",
    Links: map[string]string{
        "payment":"PM123",
    },

}
// create a refund
client.LoggingEnabled = true
refund, err := client.Refund.CreateRefund(refundCreateReq)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while creating a refund :%s", err.Error())
}
fmt.Sprintf("The refund created has the reference: %s ", refund.Reference)

// Get Refund by ID
refund, err = client.Refund.GetRefund("RF123")
if err != nil {
    fmt.Sprintf("The error while getting a refund is :%s", err.Error())
}
fmt.Sprintf("The refund retrieved with ID: %d is : %s", refund.ID, refund.Reference)


refundUpdateReq := &Refund{
    Reference: "New reference",
    Amount: 10,
}

// update a refund
client.LoggingEnabled = true
refundToUpdate, err := client.Refund.UpdateRefund(refundUpdateReq, make(map[string]string))
if err != nil {
    // do something with error
    fmt.Sprintf("The error while updating a refund is :%s", err.Error())
}
fmt.Sprintf("The refund updated is: %s ", refundToUpdate.Reference)

```

##### 12. Subscriptions

* Making the request

```go
subCreateReq := &SubscriptionCreateRequest{
    Amount: 100,
    Currency: "GBP",
    IntervalUnit: "monthly",


}
// create a subscription
client.LoggingEnabled = true
subscription, err := client.Subscription.CreateSubscription(subCreateReq)
if err != nil {
    // do something with error
    fmt.Sprintf("The error while creating a subscription :%s", err.Error())
}
fmt.Sprintf("The subscription created has the ID: %s ", subscription.ID)

// Get Subscription by ID
subscription, err = client.Subscription.GetSubscription("SB123")
if err != nil {
    fmt.Sprintf("The error while getting a subscription is :%s", err.Error())
}
fmt.Sprintf("The subscription retrieved with ID: %d is : %s", subscription.ID, subscription.Name)


subUpdateReq := &Subscription{
    Amount: 10,
}

// update a refund
client.LoggingEnabled = true
subToUpdate, err := client.Subscription.UpdateSubscription(subUpdateReq, "sample", "SB123", make(map[string]string))
if err != nil {
    // do something with error
    fmt.Sprintf("The error while updating a subscription is :%s", err.Error())
}
fmt.Sprintf("The subscription updated is: %s ", subToUpdate.Name)

// cancel a mandate
subCancelReq := &Subscription{
    Metadata: map[string]string{
        "order_no": "ABCD1234",
    },
}
client.LoggingEnabled = true
subToCancel, err := client.Subscription.CancelSubscription(subCancelReq, make(map[string]string))
if err != nil {
    // do something with error
    fmt.Sprintf("The error while canceling a subscription is :%s", err.Error())
}
fmt.Sprintf("The subscription canceled returned the response: %s ", subToCancel)
```