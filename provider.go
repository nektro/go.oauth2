package oauth2

type Provider struct {
	ID           string
	AuthorizeURL string
	TokenURL     string
	MeURL        string
	Scope        string
	NameProp     string
	NamePrefix   string
}

var (
)
