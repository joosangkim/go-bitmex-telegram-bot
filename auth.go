// functions
// 1. Init socket
// 2. Get market value
// 3. Get user status
// TBD. orderer

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"
)

// auth _
func SetAuthHeader(req *http.Request, key, method, path, query, data, secret string) *http.Request { //(string key, secret) {
	// unix timestamp + key  + secret
	expire := time.Now().Add(time.Minute * 5).Unix()
	req.Header.Add("api-expires", strconv.FormatInt(expire, 10))
	req.Header.Add("api-key", key)
	sign := signature(key, method, path, query, strconv.FormatInt(expire, 10), data)
	req.Header.Add("api-signature", sign)
	// DebugHttpRequest(req)
	return req
}

// signature _
func signature(secret, method, path, query, nonce, bodyStr string) string {
	msg := ""
	if "" == query {
		msg = strings.ToUpper(method) + path + nonce + bodyStr
	} else {
		msg = strings.ToUpper(method) + path + "?" + query + nonce + bodyStr
	}
	return calcMsg(secret, msg)
}

// calcMsg _
func calcMsg(secret, msg string) string {
	sig := hmac.New(sha256.New, []byte(secret))
	sig.Write([]byte(msg))
	return hex.EncodeToString(sig.Sum(nil))
}

func DebugHttpRequest(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}
