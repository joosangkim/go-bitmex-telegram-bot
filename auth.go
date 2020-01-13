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
func auth(req *http.Request, key, verb, path, data, secret string) { //(string key, secret) {
	// unix timestamp + key  + secret
	expire := time.Now().Add(time.Minute * 5).Unix()
	req.Header.Add("api-expires", strconv.FormatInt(expire, 10))
	req.Header.Add("api-key", key)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(verb + path + strconv.FormatInt(expire, 10) + data))
	req.Header.Add("api-signature", hex.EncodeToString(h.Sum(nil)))
	// DebugHttpRequest(req)
	return

}

func DebugHttpRequest(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}

func Signature(apiSecret, method, path, query, nonce, bodyStr string) string {
	str := ""
	if "" == query {
		str = strings.ToUpper(method) + path + nonce + bodyStr
	} else {
		str = strings.ToUpper(method) + path + "?" + query + nonce + bodyStr
	}
	return CalSignature(apiSecret, str)
}

func CalSignature(apiSecret, payload string) string {
	sig := hmac.New(sha256.New, []byte(apiSecret))
	sig.Write([]byte(payload))
	return hex.EncodeToString(sig.Sum(nil))
}
