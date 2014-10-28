package main

import (
	"fmt"
	"html/template"
	"lab/sae"
	"net/http"
)

func main() {

	// ch := make(chan int)
	// go func(ch chan int) {
	// ch <- 1
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/token", tokenHandler)
	http.ListenAndServe(":80", nil)
	// }(ch)
	// <-ch
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	s := sae.NewOAuth()

	redirect_uri := "http://www.52niuniu.net/token"
	response_type := "code"
	state := ""
	display := ""
	str := s.GetAuthorizeURL(redirect_uri, response_type, state, display)
	fmt.Println(str)
	t, err := template.New("foo").Parse(`{{define "T"}}<a href="{{.}}">访问微博</a>{{end}}`)
	err = t.ExecuteTemplate(w, "T", str)
	if err != nil {
		panic(err.Error())
	}
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {

}
