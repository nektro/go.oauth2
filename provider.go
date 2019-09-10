package oauth2

type Provider struct {
	ID           string `json:"id"`
	AuthorizeURL string `json:"authorize_url"`
	TokenURL     string `json:"token_url"`
	MeURL        string `json:"me_url"`
	Scope        string `json:"scope"`
	NameProp     string `json:"name_prop"`
	NamePrefix   string `json:"name_prefix"`
}

type AppConf struct {
	For    string `json:"for"`
	ID     string `json:"id"`
	Secret string `json:"secret"`
	Extra1 string `json:"extra_1"`
	Extra2 string `json:"extra_2"`
	Extra3 string `json:"extra_3"`
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
	ProviderFacebook = Provider{
		"facebook",
		"https://graph.facebook.com/oauth/authorize",
		"https://graph.facebook.com/oauth/access_token",
		"https://graph.facebook.com/me",
		"",
		"name",
		"",
	}
	ProviderGitHub = Provider{
		"github",
		"https://github.com/login/oauth/authorize",
		"https://github.com/login/oauth/access_token",
		"https://api.github.com/user",
		"read:user",
		"login",
		"@",
	}
	ProviderGoogle = Provider{
		"google",
		"https://accounts.google.com/o/oauth2/v2/auth",
		"https://www.googleapis.com/oauth2/v4/token",
		"https://www.googleapis.com/oauth2/v1/userinfo?alt=json",
		"profile",
		"name",
		"",
	}
	ProviderMicrosoft = Provider{
		"microsoft",
		"https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
		"https://login.microsoftonline.com/common/oauth2/v2.0/token",
		"https://graph.microsoft.com/v1.0/me/",
		"https://graph.microsoft.com/user.read",
		"displayName",
		"",
	}
	ProviderReddit = Provider{
		"reddit",
		"https://old.reddit.com/api/v1/authorize",
		"https://old.reddit.com/api/v1/access_token",
		"https://oauth.reddit.com/api/v1/me",
		"identity",
		"name",
		"u/",
	}
)

var (
	ProviderIDMap = map[string]Provider{
		"discord":   ProviderDiscord,
		"facebook":  ProviderFacebook,
		"github":    ProviderGitHub,
		"google":    ProviderGoogle,
		"microsoft": ProviderMicrosoft,
		"reddit":    ProviderReddit,
	}
)
