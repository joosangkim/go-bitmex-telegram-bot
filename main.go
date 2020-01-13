// functions
// 1. Init socket
// 2. Get market value
// 3. Get user status
// TBD. orderer

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	u := url.URL{
		Scheme: "https",
		Host:   "testnet.bitmex.com",
		Path:   "api/v1",
	}
	client := http.Client{}
	req, err := http.NewRequest("GET", u.String()+"/user/wallet", nil)
	if nil != err {
		panic(err.Error())
	}
	key := ""
	verb := "GET"
	path := "/api/v1/user/wallet"
	data := ""
	secret := ""
	auth(req, key, verb, path, data, secret)
	fmt.Println(req.Header.Get("api-expires"))
	fmt.Println(req.Header.Get(("api-key")))
	fmt.Println(req.Header.Get("api-signature"))
	resp, err := client.Do(req)
	// resp, err := http.Get(u.String() + "/user/wallet")
	if nil != err {
		panic(err.Error())
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	fmt.Printf(newStr)

}
