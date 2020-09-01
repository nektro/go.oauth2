package oauth2

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/nektro/go-util/util"
)

type SaveInfoFunc func(req http.ResponseWriter, res *http.Request, provider string, id string, name string, resp map[string]interface{})

func HandleOAuthLogin(isLoggedIn func(*http.Request) bool, doneURL string, idp Provider, appID, callbackPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isLoggedIn(r) {
			w.Header().Add("Location", doneURL)
		} else {
			urlR, _ := url.Parse(idp.AuthorizeURL)
			parameters := url.Values{}
			parameters.Add("client_id", appID)
			parameters.Add("redirect_uri", util.FullHost(r)+callbackPath)
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

func HandleOAuthCallback(idp Provider, appID, appSecret string, saveInfo SaveInfoFunc, doneURL, callbackPath string) http.HandlerFunc {
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
		parameters.Add("redirect_uri", util.FullHost(r)+callbackPath)
		parameters.Add("state", "none")

		req, _ := http.NewRequest("POST", idp.TokenURL, strings.NewReader(parameters.Encode()))
		req.Header.Set("User-Agent", "nektro/go-util")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(appID+":"+appSecret)))
		req.Header.Set("Accept", "application/json")

		body, err := util.DoHttpFetch(req)
		if err != nil {
			fmt.Fprintln(w, "error:", "POST:", idp.TokenURL)
			fmt.Fprintln(w, err.Error())
			return
		}
		var resp map[string]interface{}
		json.Unmarshal(body, &resp)
		at := resp["access_token"]
		if at == nil {
			fmt.Fprintln(w, "Identity Provider Login Error!")
			fmt.Fprintln(w, string(body))
			return
		}

		if len(idp.IDProp) == 0 {
			idp.IDProp = "id"
		}

		req2, _ := http.NewRequest("GET", idp.MeURL, strings.NewReader(""))
		req2.Header.Set("User-Agent", "nektro/go.auth2")
		req2.Header.Set("Authorization", "Bearer "+at.(string))
		req2.Header.Set("Accept", "application/json")

		body2, err := util.DoHttpFetch(req2)
		if err != nil {
			fmt.Fprintln(w, "error:", "GET:", idp.MeURL)
			fmt.Fprintln(w, err.Error())
			return
		}
		var respMe map[string]interface{}
		json.Unmarshal(body2, &respMe)
		_id := fixID(respMe[idp.IDProp])
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
	case int:
		return strconv.Itoa(id.(int))
	}
	return fmt.Sprintf("%v", id)
}

func HandleMultiOAuthLogin(isLoggedIn func(*http.Request) bool, doneURL string, clients []AppConf, callbackPath string) http.HandlerFunc {
	for _, item := range vcc {
		keys := strings.Split(item, "|")
		clients = append(clients, AppConf{keys[0], keys[1], keys[2], "", "", ""})
	}
	for i, item := range clients {
		cfr := strings.Split(item.For, ",")
		if len(cfr) == 1 {
			clients[i] = item
		}
		if len(cfr) == 2 {
			newprov, ok := ProviderIDMap["_"+cfr[0]]
			util.DieOnError(util.Assert(ok, "Unable to find customable provider '"+cfr[0]+"'"))
			util.DieOnError(util.Assert(newprov.Customable == true, "Attemped to use non-customizable provider '"+cfr[0]+"' with custom domain '"+cfr[1]+"'"))
			item.For = cfr[0] + "(" + cfr[1] + ")"
			newprov.ID = cfr[0] + "(" + cfr[1] + ")"
			newprov.AuthorizeURL = strings.ReplaceAll(newprov.AuthorizeURL, "{domain}", cfr[1])
			newprov.TokenURL = strings.ReplaceAll(newprov.TokenURL, "{domain}", cfr[1])
			newprov.MeURL = strings.ReplaceAll(newprov.MeURL, "{domain}", cfr[1])
			ProviderIDMap[item.For] = newprov
			clients[i] = item
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		with := r.URL.Query().Get("with")
		if len(with) == 0 {
			if len(clients) == 0 {
				http.NotFound(w, r)
				fmt.Fprintln(w, "No OAuth2 App configurations found.")
				return
			}
			if len(clients) == 1 {
				HandleOAuthLogin(isLoggedIn, doneURL, ProviderIDMap[clients[0].For], clients[0].ID, callbackPath)(w, r)
				return
			}
			if len(doa) > 0 {
				for _, item := range clients {
					if item.For == doa {
						HandleOAuthLogin(isLoggedIn, doneURL, ProviderIDMap[item.For], item.ID, callbackPath)(w, r)
						return
					}
				}
			}
			reader, err := mfs.Open("/selector.hbs")
			if err != nil {
				fmt.Fprintln(w, "error:", err)
				return
			}
			bytes, err := ioutil.ReadAll(reader)
			if err != nil {
				fmt.Fprintln(w, "error:", err)
				return
			}
			template := string(bytes)
			result, _ := raymond.Render(template, map[string]interface{}{
				"clients":   clients,
				"providers": ProviderIDMap,
			})
			fmt.Fprintln(w, result)

		} else {
			for _, item := range clients {
				if item.For == with {
					HandleOAuthLogin(isLoggedIn, doneURL, ProviderIDMap[item.For], item.ID, callbackPath)(w, r)
				}
			}
		}
	}
}

func HandleMultiOAuthCallback(doneURL string, clients []AppConf, saveInfo SaveInfoFunc, callbackPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idp := r.URL.Query().Get("state")
		for _, item := range clients {
			if item.For == idp {
				HandleOAuthCallback(ProviderIDMap[idp], item.ID, item.Secret, saveInfo, doneURL, callbackPath)(w, r)
				return
			}
		}
		fmt.Fprintln(w, "error: No handler found for provider: "+idp)
	}
}

func GetHandlers(isLoggedIn func(*http.Request) bool, doneURL, callbackPath string, clients []AppConf, saveInfo SaveInfoFunc) (http.HandlerFunc, http.HandlerFunc) {
	l := HandleMultiOAuthLogin(isLoggedIn, doneURL, clients, callbackPath)
	c := HandleMultiOAuthCallback(doneURL, clients, saveInfo, callbackPath)
	return l, c
}
