package service

type response struct {
	errorCode int
	errorMsg string
	result interface{}
}


const (
	OK = 0
)

var allErrorMsg = map[int]string{
	OK: "",
}