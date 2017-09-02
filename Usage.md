#### 1. Customers

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

#### 2. Creditors

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

#### 3. Customer Bank Account

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

#### 4. Creditor Bank Account

* Making request

