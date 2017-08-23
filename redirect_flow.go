package gocardless_pro_go


type RedirectFlowService service

type RedirectFlow struct {
	ID 							string              `json:"id,omitempty"`
	Description					string				`json:"description,omitempty"`
	CreatedAt					string				`json:"created_at,omitempty"`
	Scheme						string				`json:"scheme,omitempty"`
	RedirectUrl					string				`json:"redirect_url,omitempty"`
	SessionToken				string				`json:"session_token,omitempty"`
	SuccessRedirectUrl			string				`json:"success_redirect_url,omitempty"`
	Links						[]MandateLink		`json:"links,omitempty"`
}
