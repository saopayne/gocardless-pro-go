[![Build Status](https://travis-ci.org/saopayne/gocardless-pro-go.svg?branch=master)](https://travis-ci.org/saopayne/gocardless-pro-go)

# GO Client for GoCardless Pro API

#### [__Work In Progress__]

Go library to access GoCardless API. This library provides a simple wrapper around the [GoCardless API](http://developer.gocardless.com/api-reference).

- [API Reference](https://developer.gocardless.com/api-reference/2015-07-06)


## Getting started

```
go get github.com/saopayne/gocardless-pro-go
```

## Initializing the client

## Usage

``` go

import "https://github.com/saopayne/gocardless-pro-go"

apiKey := "sandbox_o55p5OowBX59Rd8aDR7c_25LQdBTHRaACeVnqj0o"

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

// Creating a redirect flow
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

```

For a more descriptive usage [Click Here](https://github.com/saopayne/gocardless-pro-go/blob/master/Usage.md)

See the test files for more examples.

## TODO
- [x] Complete All Api calls
- [ ] Documentation
- [ ] Write Unit tests
- [ ] Better handling of API call errors


## CONTRIBUTING
- Fork the repository, make necessary changes and send the PR.
