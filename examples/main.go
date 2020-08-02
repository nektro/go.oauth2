package main

import (
	"net/http"

	"github.com/nektro/go-util/vflag"
	"github.com/nektro/go.etc/htp"
	oauth2 "github.com/nektro/go.oauth2"
)

func main() {

	vflag.Parse()

	htp.Init()
	iLI := func(*http.Request) bool { return false }
	htp.Register("/login", http.MethodGet, oauth2.HandleMultiOAuthLogin(iLI, "./", []oauth2.AppConf{}))
	htp.StartServer(8000)
}
