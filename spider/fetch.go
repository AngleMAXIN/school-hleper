package spider

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var ()

func init() {
	rand.Seed(time.Now().UnixNano())
}

type httpClient struct {
	client           *http.Client
	resp             *http.Response
	cookies  []*http.Cookie
	retry            int
	userAgentListLen int
}

func CreateHttpClient() *httpClient {
	return &httpClient{
		userAgentListLen: len(USER_AGENT_LIST),
		client:           &http.Client{},
	}
}


func (hc *httpClient) POST(url string, reqBody io.Reader) []byte {
	respContent, err := hc.doRequest("POST", url, reqBody)
	if err != nil {
		log.Printf("[POST] url: %s, reqBody: %s, error: %s\n", url, reqBody, err)
		return nil
	}
	return respContent
}

func (hc *httpClient) GET(url string, query io.Reader) []byte {
	respContent, err := hc.doRequest("GET", url, query)
	if err != nil {
		log.Printf("[GET] url: %s, error: %s\n", url, err)
		return nil
	}
	return respContent
}

// setHeaders 请求头
func (hc *httpClient) setHeaders(req *http.Request) {
	index := 0
	req.Header.Add("User-Agent", USER_AGENT_LIST[index])
}

func (hc *httpClient) setCookie(req *http.Request) {
	for _, cookie := range hc.cookies{
		req.AddCookie(cookie)
	}
}

func (hc *httpClient) storeCookie() {
	hc.cookies = hc.resp.Cookies()
}

func (hc *httpClient) build() {

}

// doRequest 发起请求
func (hc *httpClient) doRequest(method string, url string, reqBody io.Reader) ([]byte, error) {
	//sleepTime := rand.Intn(6666)
	//time.Sleep(time.Microsecond * time.Duration(sleepTime))

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		log.Printf("http.NewRequest error")
	}

	hc.setHeaders(req)
	retry :=  hc.retry

	for hc.resp, err = hc.client.Do(req); retry > 0; retry-- {
		if err != nil {
			log.Printf("client do is error")
			return nil, errors.New("client.Do error")
		} else {
			retry = 0
		}
	}
	if hc.resp == nil {
		log.Printf("resp is nil")
		return nil, err
	}

	defer hc.resp.Body.Close()
	if hc.resp.StatusCode != 200 {
		log.Println("status code is normal")
		return nil, errors.New(fmt.Sprintf("status code is: %d", hc.resp.StatusCode))
	}

	body, err := ioutil.ReadAll(hc.resp.Body)
	if err != nil {
		log.Println("read all body error")
		return nil, errors.New(fmt.Sprintf("read all body error: %s", err))

	}
	return body, nil
}
