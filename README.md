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

## Installing
```
$ go get -u github.com/nektro/go.oauth2
```

## License
MIT
