package sae

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	// "reflect"
)

const (
	WB_API_URL = "https://api.weibo.com/2"
)

func (o *OAuth) UsersShow(uid string) (r interface{}) {
	param := url.Values{}
	param.Add("uid", uid)
	param.Add("access_token", o.AccessToken)
	str := param.Encode()

	res, err := http.Get(WB_API_URL + "/users/show.json?" + str)
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	json.Unmarshal(result, &r)

	return r
}

func (o *OAuth) FriendshipsFriends(uid, count, cursor string) (r interface{}) {
	param := url.Values{}
	param.Add("uid", uid)
	param.Add("count", count)
	param.Add("cursor", cursor)
	param.Add("access_token", o.AccessToken)
	str := param.Encode()
	res, err := http.Get(WB_API_URL + "/friendships/friends.json?" + str)
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	json.Unmarshal(result, &r)

	return r
}
func (o *OAuth) Upload(uid string) (r interface{}) {
	param := url.Values{}
	param.Add("uid", uid)

	str := param.Encode()

	body := bytes.NewBuffer([]byte(str))
	res, err := http.Post(WB_API_URL+"/users/show.json", "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	json.Unmarshal(result, &r)

	return r
}
