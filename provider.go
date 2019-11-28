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
}

type AppConf struct {
	For    string `json:"for"`
	ID     string `json:"id"`
	Secret string `json:"secret"`
	Extra1 string `json:"extra_1"`
	Extra2 string `json:"extra_2"`
	Extra3 string `json:"extra_3"`
}
