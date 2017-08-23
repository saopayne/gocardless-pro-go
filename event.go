package gocardless_pro_go

type EventService service

type Event struct {
	Action			string					`json:"action,omitempty"`
	ID        		string  				`json:"id,omitempty"`
	Links			[]EventLink				`json:"links,omitempty"`
	Metadata		map[string]string		`json:"metadata,omitempty"`
	ResponseUrl	  	string					`json:"responseurl,omitempty"`
	Details			Details					`json:"details,omitempty"`
	ResourceType    ResourceType  			`json:"resource_type,omitempty"`
}

type Details struct {
	Cause			string 		`json:"cause,omitempty"`
	Description		string		`json:"description,omitempty"`
	Origin			Origin		`json:"origin,omitempty"`
	ReasonCode		string		`json:"reason_code,omitempty"`
	Scheme			Scheme		`json:"scheme,omitempty"`
}

type Origin struct {
	Origin			string		`json:"origin,omitempty"`
}

type Scheme struct {
	Scheme			string		`json:"scheme,omitempty"`
}

type ResourceType struct {
	ResourceType 	string		`json:"resource_type,omitempty"`
}