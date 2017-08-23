package gocardless_pro_go

type CreditorService service

type SchemeIdentifier struct {
	Name string
}

type SchemeIdentifierList struct {
	list []SchemeIdentifier
}

type CreditorRequest struct {
	Id        			string  					`json:"id,omitempty"`
	Name      			string  					`json:"name,omitempty"`
	Region          	string 						`json:"region,omitempty"`
	PostalCode			string 						`json:"postal_code,omitempty"`
	LogoUrl				string 						`json:"logo_url,omitempty"`
	City				string 						`json:"city,omitempty"`
	AddressLine1		string						`json:"address_line_1,omitempty"`
	AddressLine2		string						`json:"address_line_2,omitempty"`
	AddressLine3		string						`json:"address_line_3,omitempty"`
	CountryCode			string						`json:"country_code,omitempty"`
	CreatedAt			string						`json:"created_at,omitempty"`
	VerificationStatus	[]string					`json:"verification_status,omitempty"`
	Links				[]string					`json:"links,omitempty"`
	SchemeIdentifiers	[]string					`json:"scheme_identifiers,omitempty"`
	ResponseUrl	  		string						`json:"responseurl,omitempty"`
	Metadata        	Metadata 					`json:"metadata,omitempty"`
}
