package spider

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var ()

func init() {
	rand.Seed(time.Now().UnixNano())
}

type httpClient struct {
	client           *http.Client
	cookies          []*http.Cookie
	resp *http.Response
	retry            int
	userAgentListLen int
}

func CreateHttpClient() *httpClient {
	return &httpClient{
		userAgentListLen: len(USER_AGENT_LIST),
		client:           &http.Client{},
		cookies: make([]*http.Cookie,0),
		retry: RETRY_COUNT,
	}
}

func (hc *httpClient) POST(url string, reqForm io.Reader, isOtherHeader bool) []byte {
	respContent, err := hc.doRequest("POST", url, reqForm, isOtherHeader)
	if err != nil {
		log.Printf("[POST] url: %s, reqBody: %s, error: %s\n", url, reqForm, err)
		return nil
	}
	return respContent
}

func (hc *httpClient) GET(url string, reqForm io.Reader, isOtherHeader bool) []byte {
	respContent, err := hc.doRequest("GET", url, reqForm, isOtherHeader)
	if err != nil {
		log.Printf("[GET] url: %s, error: %s\n", url, err)
		return nil
	}
	return respContent
}

// setHeaders 请求头
func (hc *httpClient) setHeaders(req *http.Request, other bool) {

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Content-Length", strconv.Itoa(int(req.ContentLength)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Proxy-Connection", "keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")

	fmt.Println("header:",req.Header)
}

func (hc *httpClient) setCookie(req *http.Request) {
	for _, cookie := range hc.cookies {
		req.AddCookie(cookie)
	}
}

func (hc *httpClient) storeCookie() {
	hc.cookies = hc.resp.Cookies()
}

// doRequest 发起请求
func (hc *httpClient) doRequest(method string, url string, reqForm io.Reader, other bool) ([]byte, error) {
	req, err := http.NewRequest(method, url, reqForm)
	if err != nil {
		log.Printf("http.NewRequest error")
		return nil, err
	}

	hc.setHeaders(req, other)
	hc.setCookie(req)

	hc.client = &http.Client{}
	retry := hc.retry
	var resp *http.Response
	for resp, err = hc.client.Do(req); retry > 0; retry-- {
		if err != nil {
			log.Printf("client do is error")
			return nil, err
		} else {
			retry = 0
		}
	}

	if resp == nil {
		log.Printf("resp is nil")
		return nil, errors.New("resp is nil")
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("status code is normal")
		return nil, errors.New(fmt.Sprintf("status code is: %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read all body error")
		return nil, err

	}
	hc.resp = resp
	return body, nil
}
