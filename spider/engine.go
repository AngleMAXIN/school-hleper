package spider

import (
	"net/url"
	"strings"
)

type Fetch struct {
	client *httpClient
}

func (ft *Fetch) login(userID, passWord  string) {

	reqBody := url.Values{}
	reqBody.Add("USERNAME",userID)
	reqBody.Add("PASSWORD",passWord)
	ft.client.POST(LOGIN_URL, strings.NewReader(reqBody.Encode()))
	ft.client.storeCookie()

}

func (ft *Fetch) pullScore(date string) {

	reqBody := url.Values{}
	reqBody.Add("kksj",date)
	reqBody.Add("xsfs","all")
	ft.client.GET(SCORE_URL, strings.NewReader(reqBody.Encode()))
}
