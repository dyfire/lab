package sae

import (
	// "bytes"
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

func (o *OAuth) Upload(content string) (r interface{}) {
	// visible := 1
	param := url.Values{}
	param.Add("status", content)
	// param.Add("visible", visible)
	param.Add("access_token", o.AccessToken)

	res, err := http.PostForm(WB_API_URL+"/statuses/update.json", param)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	s, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("read failtrue")
	}
	err = json.Unmarshal(s, &r)
	if err != nil {
		log.Fatal("er")
	}

	// str := param.Encode()

	// body := bytes.NewBuffer([]byte(str))
	// res, err := http.Post(WB_API_URL+"/statuses/update.json", "", body)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// result, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()

	// json.Unmarshal(result, &r)

	// client := &http.Client{}
	// req, err := http.NewRequest("POST", WB_API_URL+"/statuses/update.json", param)
	// log.Println(WB_API_URL + "/statuses/update.json")
	// if err != nil {
	// 	panic("errr")
	// 	log.Fatal("post failtrue")
	// }

	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	// req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	// req.Header.Add("Accept-Encoding", "gzip, deflate")
	// req.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	// req.Header.Add("Connection", "keep-alive")
	// req.Header.Add("Host", "login.sina.com.cn")
	// req.Header.Add("Referer", "http://weibo.com/")
	// req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	// resp, err := client.Do(req)
	// defer resp.Body.Close()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// log.Println(resp)
	// result, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()

	// json.Unmarshal(result, &r)
	return r
}
