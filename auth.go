// functions
// 1. Init socket
// 2. Get market value
// 3. Get user status
// TBD. orderer

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"strconv"
	"time"
)

// auth _
func auth(requset http, key, verb, path, data string, secret []byte) { //(string key, secret) {
	// unix timestamp + key  + secret
	expire := time.Now().Add(time.Minute * 5).Unix()

	h := hmac.New(sha256.New, secret)
	h.Sum([]byte(verb + path + strconv.FormatInt(expire, 10) + data))

}

func getWallet() {

}
