package spider

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

type Fetch struct {
	client *httpClient
}

func CreateNewFetch() *Fetch {
	return &Fetch{
		client: CreateHttpClient(),
	}
}

func (ft *Fetch) Prepare()  {
	ft.client.GET(COOKIE_URL,nil,false)
	ft.client.storeCookie()
}

func (ft *Fetch) Login(userID, passWord string) {

	reqBody := url.Values{}
	reqBody.Add("USERNAME",userID)
	reqBody.Add("PASSWORD",passWord)

	content := ft.client.POST(LOGIN_URL, strings.NewReader(reqBody.Encode()),false)
	fmt.Println("-------------\n",string(content))

}

func (ft *Fetch) PullScore(date string) []*itemScore {

	reqBody := url.Values{}
	reqBody.Add("kksj",date)
	reqBody.Add("kcxz","")
	reqBody.Add("kcmc","")
	reqBody.Add("xsfs","all")
	content := ft.client.POST(SCORE_URL,strings.NewReader(reqBody.Encode()),true)
	if content == nil {
		log.Println("res content is nil")
		return nil
	}
	infoList, err := parseHtml(string(content))
	if err != nil {
		log.Printf("parse html error: %s",err)
		return nil
	}
	if infoList == nil || len(infoList) == 0{
		log.Println("info list is nil")
		return nil
	}
	return buildComplete(infoList)
}
