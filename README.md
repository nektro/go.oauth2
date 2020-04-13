# go.oauth2
![loc](https://sloc.xyz/github/nektro/go.oauth2)
[![license](https://img.shields.io/github/license/nektro/go.oauth2.svg)](https://github.com/nektro/go.oauth2/blob/master/LICENSE)
[![discord](https://img.shields.io/discord/551971034593755159.svg)](https://discord.gg/P6Y4zQC)
[![paypal](https://img.shields.io/badge/donate-paypal-009cdf)](https://paypal.me/nektro)
[![goreportcard](https://goreportcard.com/badge/github.com/nektro/go.oauth2)](https://goreportcard.com/report/github.com/nektro/go.oauth2)
[![codefactor](https://www.codefactor.io/repository/github/nektro/go.oauth2/badge)](https://www.codefactor.io/repository/github/nektro/go.oauth2)

HTTP function handlers to easily add OAuth2 client support to your Go application

## `AppConf` Schema
```go
type AppConf struct {
	For    string `json:"for"`
	ID     string `json:"id"`
	Secret string `json:"secret"`
	Extra1 string `json:"extra_1"`
	Extra2 string `json:"extra_2"`
	Extra3 string `json:"extra_3"`
}
```
- `"for"` is the short-code this config refers to.
- `"id"` is your Client ID.
- `"secret"` is for your Client Secret.
- Extra 1, 2, and 3 are filler spots for misc. info your app may need, such as Discord's Bot Token for example.

## Creating Credentials
In order to use an app that uses this library, you will need to create an app on your Identity Provider of choice. Below, you will see a table of the supported Identity Providers and a link to the respective dashboards where you can go to create your [app](#appconf-schema) and obtain your App ID and App Secret.

| Identity Provider | Short Code | Developer Dashboard |
| --- | --- | --- |
| Amazon | `amazon` | https://developer.amazon.com/settings/console/securityprofile/overview.html |
| Battle.net | `battle.net` | https://develop.battle.net/access/clients |
| Discord | `discord` | https://discordapp.com/developers/applications/ |
| Facebook | `facebook` | https://developers.facebook.com/apps/ |
| GitHub | `github` | https://github.com/settings/developers |
| GitLab | `gitlab.com` | https://gitlab.com/profile/applications |
| Google | `google` | https://console.developers.google.com |
| Microsoft | `microsoft` | https://apps.dev.microsoft.com/ |
| Reddit | `reddit` | https://www.reddit.com/prefs/apps |

## Installing
```
$ go get -u github.com/nektro/go.oauth2
```

## `Provider` Schema
```go
type Provider struct {
	ID           string `json:"id"`
	AuthorizeURL string `json:"authorize_url"`
	TokenURL     string `json:"token_url"`
	MeURL        string `json:"me_url"`
	Scope        string `json:"scope"`
	NameProp     string `json:"name_prop"`
	NamePrefix   string `json:"name_prefix"`
	Logo         string `json:"logo"`
	Color        string `json:"color"`
	Customable   string `json:"customable"`
}
```
- `"id"` is the short-code this is creating.
- `"authorize_url"` is the OAuth2 authorization URL.
- `"token_url"` is the OAuth2 token URL.
- `"me_url"` is the service's URL to get the currently logged in user.
- `"scope"` is the OAuth2 scope required to be able to get the currently logged in user.
- `"name_prop"` is the JSON key of current user's real name in the response of fetching `"me_url"`.
- `"name_prefix"` is any prefix to put in front of all names, this is typically `@`, `u/`, blank, etc.

## `AppConf` Details for Self-Hosted Services

There are also a number of providers that allow you to specify a custom domain for that provider. They are accessed as such:

```json
...
"clients": [
	{
		"for": "{provider_id},{domain}",
		"id": "",
		"secret": ""
	}
],
...
```

So for example, if adding a login config for https://mastodon.social/, your `"for"` key would be `"mastodon,mastodon.social"`

The full list of customizable provider are as follows:

| Identity Provider | Short Code | Home Site |
| --- | --- | --- |
| Gitea | `gitea` | https://gitea.io/en-us/ |
| Gitlab | `gitlab` | https://about.gitlab.com/ |
| mastodon | `mastodon` | https://joinmastodon.org/ |
| pleroma | `pleroma` | https://pleroma.social/ |
## License
MIT
