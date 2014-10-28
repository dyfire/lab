package main

import (
	"fmt"
	"net/http"
	"weibo/sae"
)

func main() {
	s := sae.NewOAuth()

	redirect_uri := "http://www.52niuniu.net"
	response_type := "code"
	state := ""
	display := ""
	str := s.GetAuthorizeURL(redirect_uri, response_type, state, display)
	fmt.Println(str)


	http.ListenAndServe("/token", tokenHandler)

}

func tokenHandler() {
	sae. 
}