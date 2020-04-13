package oauth2

type Provider struct {
	ID           string `json:"id"`
	AuthorizeURL string `json:"authorize_url"`
	TokenURL     string `json:"token_url"`
	MeURL        string `json:"me_url"`
	Scope        string `json:"scope"`
	NameProp     string `json:"name_prop"`
	NamePrefix   string `json:"name_prefix"`
	IDProp       string `json:"id_prop"`
	Logo         string `json:"logo"`
	Color        string `json:"color"`
	Customable   bool   `json:"customable"`
}
