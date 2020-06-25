package spider

import (
	"log"
	"net/http"
)

func getCookiesFromLoginView() (string, error) {
	resp, err := http.Get(LOGIN_VIEW_PAGE)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// 获取cookie
	sourceCookie := resp.Header.Get("Set-Cookie")
	return getValueFromCookie(sourceCookie), nil
}

func decryptPassword(password string) string {
	password, err := Decrypt(password)
	if err != nil {
		log.Printf("Decrypt password failed,password: %s\n", password)
		return ""
	}
	return password
}
