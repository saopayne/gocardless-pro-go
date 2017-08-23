[![Build Status](https://travis-ci.org/rpip/gocardless-pro-go.svg?branch=master)](https://travis-ci.org/rpip/gocardless-pro-go)

# GO Client for GoCardless Pro API

This library provides a simple wrapper around the [GoCardless API](http://developer.gocardless.com/api-reference).

- ["Getting started" guide](https://developer.gocardless.com/getting-started/api/introduction/?lang=java) with copy and paste Java code samples
- [API Reference](https://developer.gocardless.com/api-reference/2015-07-06)
- [Example application](https://github.com/gocardless/gocardless-pro-java-example)

## Getting started

```
go get 'https://github.com/saopayne/gocardless-pro-go'
```

## Initializing the client

## Usage

``` go
import "https://github.com/saopayne/gocardless-pro-go"

client = gocardless.NewClient(apiKey)

cust := &Customer{
    FirstName: "User123",
    LastName:  "AdminUser",
    Email:     "user123@gmail.com",
    Phone:     "+4400000000000000",
}
// create the customer
customer, err := c.Customer.Create(cust)
if err != nil {
    // do something with error
}

// Get customer by ID
client.Customers.Get(customer.ID)

// retrieve list of plans
ch, err := client.Plan.List()
```

See the test files for more examples.

## TODO
- [ ] Documentation
- [ ] More test cases
- [ ] Better handling of API call errors
- [ ] Upload godocs

## CONTRIBUTING
