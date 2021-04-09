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
	clients := &[]oauth2.AppConf{}
	lg, cb := oauth2.GetHandlers(iLI, "/", "/callback", clients, nil)
	htp.Register("/login", http.MethodGet, lg)
	htp.Register("/callback", http.MethodGet, cb)
	htp.StartServer(80)
}
