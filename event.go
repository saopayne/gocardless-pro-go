package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type EventService service

type Event struct {
	Action       string            `json:"action,omitempty"`
	ID           string            `json:"id,omitempty"`
	Links        []EventLink       `json:"links,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	ResponseUrl  string            `json:"responseurl,omitempty"`
	Details      Details           `json:"details,omitempty"`
	ResourceType ResourceType      `json:"resource_type,omitempty"`
}

type Details struct {
	Cause       string `json:"cause,omitempty"`
	Description string `json:"description,omitempty"`
	Origin      Origin `json:"origin,omitempty"`
	ReasonCode  string `json:"reason_code,omitempty"`
	Scheme      Scheme `json:"scheme,omitempty"`
}

type Origin struct {
	Origin string `json:"origin,omitempty"`
}

type Scheme struct {
	Scheme string `json:"scheme,omitempty"`
}

type ResourceType struct {
	ResourceType string `json:"resource_type,omitempty"`
}

type EventListRequest struct {
	CreatedAt    CreatedAt    `json:"created_at,omitempty"`
	Limit        int          `json:"limit,omitempty"`
	Before       string       `json:"before,omitempty"`
	After        string       `json:"after,omitempty"`
	Include      string       `json:"include,omitempty"`
	Action       string       `json:"action,omitempty"`
	Mandate      string       `json:"mandate,omitempty"`
	ParentEvent  string       `json:"parent_event,omitempty"`
	Payment      string       `json:"payment,omitempty"`
	Refund       string       `json:"refund,omitempty"`
	Payout       string       `json:"payout,omitempty"`
	ResourceType ResourceType `json:"resource_type,omitempty"`
	Subscription Subscription `json:"subscription,omitempty"`
}

type EventList struct {
	Meta   ListMeta
	Values []Event `json:"data"`
}


// Returns a cursor-paginated list of your events.
// https://developer.gocardless.com/api-reference/#events-list-events
func (s *EventService) ListEvents(req *EventListRequest) (*EventList, error) {
	reqd, err := http.NewRequest("GET", "/events", nil)
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
	if req.CreatedAt.Lte != "" { params.Add("created_at[lte]", req.CreatedAt.Lte) }
	if req.Limit > 0 { params.Add("limit", string(req.Limit)) }
	if req.Action != "" { params.Add("action", req.Action) }
	if req.Include != "" { params.Add("include", req.Include) }
	if req.Mandate != "" { params.Add("mandate", req.Mandate) }
	if req.ParentEvent != "" { params.Add("parent_event", req.ParentEvent) }
	if req.Payout != "" { params.Add("payout", req.Payout) }
	if req.Payment != "" { params.Add("payment", req.Payment) }
	if req.Refund != "" { params.Add("refund", req.Refund) }
	if req.ResourceType.ResourceType != "" { params.Add("resource_type", req.ResourceType.ResourceType) }
	if req.Subscription.ID != "" { params.Add("subscription", req.Subscription.ID)}

	reqd.URL.RawQuery = params.Encode()

	path := reqd.URL.String()

	events := &EventList{}
	err = s.client.Call("GET", path, nil, events)

	return events, err
}

// Get returns the details of an event.
func (s *EventService) GetEvent(id string) (*Event, error) {
	u := fmt.Sprintf("/events/%s", id)
	event := &Event{}
	err := s.client.Call("GET", u, nil, event)

	return event, err
}
