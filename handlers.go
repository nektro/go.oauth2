package oauth2

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/nektro/go-util/alias"
	"github.com/nektro/go-util/util"
)

func HandleOAuthLogin(isLoggedIn func(*http.Request) bool, doneURL string, idp Provider, appID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isLoggedIn(r) {
			w.Header().Add("Location", doneURL)
		} else {
			urlR, _ := url.Parse(idp.AuthorizeURL)
			parameters := url.Values{}
			parameters.Add("client_id", appID)
			parameters.Add("redirect_uri", util.FullHost(r)+"/callback")
			parameters.Add("response_type", "code")
			parameters.Add("scope", idp.Scope)
			parameters.Add("duration", "temporary")
			parameters.Add("state", idp.ID)
			urlR.RawQuery = parameters.Encode()
			w.Header().Add("Location", urlR.String())
		}
		w.WriteHeader(http.StatusFound)
	}
}

func HandleOAuthCallback(idp Provider, appID, appSecret string, saveInfo func(http.ResponseWriter, *http.Request, string, string, string, map[string]interface{}), doneURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if len(code) == 0 {
			return
		}

		parameters := url.Values{}
		parameters.Add("client_id", appID)
		parameters.Add("client_secret", appSecret)
		parameters.Add("grant_type", "authorization_code")
		parameters.Add("code", string(code))
		parameters.Add("redirect_uri", util.FullHost(r)+"/callback")
		parameters.Add("state", "none")

		urlR, _ := url.Parse(idp.TokenURL)
		req, _ := http.NewRequest("POST", urlR.String(), strings.NewReader(parameters.Encode()))
		req.Header.Set("User-Agent", "nektro/go-util")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(appID+":"+appSecret)))
		req.Header.Set("Accept", "application/json")

		body := util.DoHttpRequest(req)
		var resp map[string]interface{}
		json.Unmarshal(body, &resp)
		at := resp["access_token"]
		if at == nil {
			b, _ := json.Marshal(resp)
			fmt.Fprintln(w, "Identity Provider Login Error!")
			fmt.Fprintln(w, string(b))
			return
		}

		urlR2, _ := url.Parse(idp.MeURL)
		req2, _ := http.NewRequest("GET", urlR2.String(), strings.NewReader(""))
		req2.Header.Set("User-Agent", "nektro/andesite")
		req2.Header.Set("Authorization", "Bearer "+at.(string))

		body2 := util.DoHttpRequest(req2)
		var respMe map[string]interface{}
		json.Unmarshal(body2, &respMe)
		_id := fixID(respMe["id"])
		_name := respMe[idp.NameProp].(string)
		saveInfo(w, r, idp.ID, _id, _name, resp)

		w.Header().Add("Location", doneURL)
		w.WriteHeader(http.StatusFound)
	}
}

func fixID(id interface{}) string {
	switch id.(type) {
	case string:
		return id.(string)
	case float64:
		return strconv.FormatFloat(id.(float64), 'f', -1, 64)
	}
	return alias.F("%v", id)
}

func HandleMultiOAuthLogin(isLoggedIn func(*http.Request) bool, doneURL string, clients []AppConf) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		with := r.URL.Query().Get("with")
		if len(with) == 0 {
			HandleOAuthLogin(isLoggedIn, doneURL, ProviderIDMap[clients[0].For], clients[0].ID)(w, r)
		} else {
			for _, item := range clients {
				if item.For == with {
					HandleOAuthLogin(isLoggedIn, doneURL, ProviderIDMap[item.For], item.ID)(w, r)
				}
			}
		}
	}
}

func HandleMultiOAuthCallback(doneURL string, clients []AppConf, saveInfo func(http.ResponseWriter, *http.Request, string, string, string, map[string]interface{})) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idp := r.URL.Query().Get("state")
		for _, item := range clients {
			if item.For == idp {
				HandleOAuthCallback(ProviderIDMap[idp], item.ID, item.Secret, saveInfo, doneURL)(w, r)
			}
		}
	}
}
