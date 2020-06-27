package service

import (
	"encoding/json"
	"net/http"
	"school-hleper/spider"
	"sync"
)
var fetchTable = sync.Map{}

func responseJson(errorCode int, w http.ResponseWriter){
	w.Header().Set("Content-Type","application/json")
	res := response{
		errorCode: OK,
		errorMsg: allErrorMsg[errorCode],
	}
	content, _:=json.Marshal(res)
	_, _ = w.Write(content)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	userName := r.Form.Get("username")
	passWord := r.Form.Get("password")
	if userName == "" || passWord == ""{
		return
	}
	fetch := spider.CreateNewFetch()
	fetch.Prepare()
	fetch.Login(userName, passWord)
	session := [3]interface{}{fetch, userName, passWord}
	fetchTable.LoadOrStore(userName, session)

	responseJson(0,w)
}

func main() {
	http.HandleFunc("/api/login", loginHandler)
}