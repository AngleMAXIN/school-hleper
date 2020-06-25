package spider

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
	"regexp"
)

var regPattern *regexp.Regexp

func init() {
	reg := `JSESSIONID=(.*); Path=/sso; HttpOnly`
	regPattern = regexp.MustCompile(reg)
}

func Decrypt(input string) (string, error) {
	bytes, err := ioutil.ReadFile(JSFILE)
	if err != nil {
		return "", err
	}
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	if err != nil {
		return "", err
	}
	enc, err := vm.Call("hex_md5", nil, input)
	if err != nil {
		return "", err
	}
	return enc.String(), nil
}

func getValueFromCookie(cookie string) string {

	// 匹配
	result := regPattern.FindAllStringSubmatch(cookie, -1)

	if len(result) < 1 || len(result[0]) < 1 {
		log.Panicln("get values from source cookie failed")
		return ""
	}
	return result[0][1]

}
