package main

import (
	"net/http"
)

// NewClient _
func NewClient() *http.Client {
	return &http.Client{}
}

// // NewRequest _
// func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
// 	req, err := http.NewRequest(method, url, body)
// 	if nil != err {
// 		return nil, err
// 	}
// 	return http.NewRequest(method, url, body)
// }
