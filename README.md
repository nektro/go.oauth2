# go.oauth2
![loc](https://tokei.rs/b1/github/nektro/go.oauth2)
[![license](https://img.shields.io/github/license/nektro/go.oauth2.svg)](https://github.com/nektro/go.oauth2/blob/master/LICENSE)
[![discord](https://img.shields.io/discord/551971034593755159.svg)](https://discord.gg/P6Y4zQC)
[![sourcegraph](https://sourcegraph.com/github.com/nektro/go.oauth2/-/badge.svg)](https://sourcegraph.com/github.com/gorilla/sessions?badge)

HTTP function handlers to easily add OAuth2 client support to your Go application

[![buymeacoffee](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/nektro)

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
| Discord | `discord` | https://discordapp.com/developers/applications/ |
| Reddit | `reddit` | https://www.reddit.com/prefs/apps |
| GitHub | `github` | https://github.com/settings/developers |
| Google | `google` | https://console.developers.google.com |
| Facebook | `facebook` | https://developers.facebook.com/apps/ |
| Microsoft | `microsoft` | https://apps.dev.microsoft.com/ |

## Installing
```
$ go get -u github.com/nektro/go.oauth2
```

## License
MIT
