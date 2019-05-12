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
	ProviderDiscord = Provider{
		"discord",
		"https://discordapp.com/api/oauth2/authorize",
		"https://discordapp.com/api/oauth2/token",
		"https://discordapp.com/api/users/@me",
		"identify",
		"username",
		"@",
	}
)
