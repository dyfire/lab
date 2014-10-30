package main

import (
	"fmt"
	"html/template"
	"lab/sae"
	"net/http"
)

var to *sae.Token
var oauth *sae.OAuth

func main() {

	// ch := make(chan int)
	// go func(ch chan int) {
	// ch <- 1
	oauth = sae.NewOAuth()
	to = &sae.Token{}
	fmt.Println("server is start")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/token", tokenHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/friends", friendHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", nil)
	// }(ch)
	// <-ch
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[auth] ", &oauth)
	redirect_uri := sae.WB_CALLBACK_URL
	response_type := "code"
	state := ""
	display := ""
	str := oauth.GetAuthorizeURL(redirect_uri, response_type, state, display)

	t, err := template.New("foo").Parse(`{{define "T"}}<a href="{{.}}">访问微博</a>{{end}}`)
	err = t.ExecuteTemplate(w, "T", str)
	if err != nil {
		panic(err.Error())
	}
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[auth] ", &oauth)
	r.ParseForm()
	types := "code"
	keys := make(map[string]string)
	keys["code"] = r.Form.Get("code")
	keys["redirect_uri"] = sae.WB_CALLBACK_URL
	token := oauth.GetAccessToken(types, keys)

	to = token

	t, err := template.New("foo").Parse(`{{define "T"}}<a href="/user">用户信息</a> <a href="/friends?cursor=0">关系</a><a href="/upload">发布</a>{{end}}`)
	err = t.ExecuteTemplate(w, "T", "hello")
	if err != nil {
		panic(err.Error())
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[auth] ", &oauth)
	t := oauth.UsersShow(to.Uid)
	f := t.(map[string]interface{})
	fmt.Println(f["location"])
	fmt.Println(t)
}

func friendHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	count := "20"
	cursor := r.Form.Get("cursor")
	t := oauth.FriendshipsFriends(to.Uid, count, cursor)
	fmt.Println(t.(map[string]interface{}))
	fmt.Println(t)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	content := "just golang test"
	t := oauth.Upload(content)
	fmt.Println(t.(map[string]interface{}))
	fmt.Println(t)
}
