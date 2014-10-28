package sae

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	WB_AKEY            = "3971074890"
	WB_SKEY            = "73a1a6642e15fa99d1e75364d3691cdd"
	WB_CALLBACK_URL    = ""
	WB_OAUTH_API_URL   = "https://api.weibo.com/oauth2"
	WB_OAUTH_AUTHORIZE = "https://api.weibo.com/oauth2/authorize"
)

type OAuth struct {
	ClientId     string
	ClientSecret string
}

func NewOAuth() *OAuth {
	return &OAuth{ClientId: WB_AKEY, ClientSecret: WB_SKEY}
}

func (o *OAuth) AuthorizeURL() string {
	return "https://api.weibo.com/oauth2/authorize"
}

func (o *OAuth) AccessTokenURL() string {
	return "https://api.weibo.com/oauth2/access_token"
}

func (o *OAuth) GetAuthorizeURL(redirect_uri, response_type, state, display string) string {
	if response_type == "" {
		response_type = "code"
	}

	if display == "" {
		display = "default"
	}

	param := url.Values{}
	param.Add("client_id", o.ClientId)
	param.Add("redirect_uri", redirect_uri)
	param.Add("response_type", response_type)
	param.Add("state", state)
	param.Add("display", display)
	params := param.Encode()
	return o.AuthorizeURL() + "?" + params

}

func (o *OAuth) GetAccessToken(types string, keys map[string]string) {
	param := url.Values{}
	param.Add("client_id", o.ClientId)
	param.Add("client_secret", o.ClientSecret)

	if types == "token" {
		param.Add("grant_type", "refresh_token")
		param.Add("refresh_token", keys["refresh_token"])
	} else if types == "code" {
		param.Add("grant_type", "authorization_code")
		param.Add("code", keys["code"])
		param.Add("redirect_uri", keys["redirect_uri"])
	} else if types == "password" {
		param.Add("grant_type", "password")
		param.Add("username", keys["username"])
		param.Add("password", keys["password"])
	} else {
		panic("wrong auth type")
	}

	params := param.Encode()
	fmt.Println(params)
	res, err := http.NewRequest("post", o.AccessTokenURL(), nil)
	if err != nil {
		log.Fatal("post data failtrue")
	}

	fmt.Println(res)
}

func (s *OAuth) Get() {
	req, err := http.Get(WB_OAUTH_AUTHORIZE)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(req)

}
