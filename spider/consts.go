package spider

const (
	//重试次数
	RETRY_COUNT = 3
	JSFILE = "js/md5.js"
	LOGIN_VIEW_PAGE = "http://218.59.189.236:8080/sso/login?service=http://cas.slcupc.edu.cn/user/simpleSSOLogin"

	LOGIN_URL = "http://218.59.189.226/jsxsd/xk/LoginToXk"

	SCORE_URL = "http://218.59.189.226/jsxsd/kscj/cjcx_list"
)

var USER_AGENT_LIST = []string{

}


//_ := `
//Accept:text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
//Accept-Encoding:gzip, deflate
//Accept-Language:zh-CN,zh;q=0.9
//Cache-Control:max-age=0
//Content-Length:39
//Content-Type:application/x-www-form-urlencoded
//Host:218.59.189.226
//Origin:http://218.59.189.226
//Proxy-Connection:keep-alive
//Referer:http://218.59.189.226/jsxsd/
//Upgrade-Insecure-Requests:1
//User-Agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36
//`