package sae

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	WB_AKEY                 = "3971074890"
	WB_SKEY                 = "73a1a6642e15fa99d1e75364d3691cdd"
	WB_CALLBACK_URL         = "http://www.52niuniu.net"
	WB_OAUTH_AuthorizeURL   = "https://api.weibo.com/oauth2/authorize"
	WB_OAUTH_AccessTokenURL = "https://api.weibo.com/oauth2/access_token"
)

type OAuth struct {
	ClientId     string
	ClientSecret string
	AccessToken  string
	RefreshToken string
}

func NewOAuth() *OAuth {
	return &OAuth{ClientId: WB_AKEY, ClientSecret: WB_SKEY}
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
	return WB_OAUTH_AuthorizeURL + "?" + params

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
	res, err := http.NewRequest("post", WB_OAUTH_AccessTokenURL, nil)
	if err != nil {
		log.Fatal("post data failtrue")
	}

	fmt.Println(res)
}
