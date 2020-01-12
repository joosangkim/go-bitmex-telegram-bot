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
	fmt.Println(u.String() + "/wallet")
	resp, err := http.Get(u.String() + "/user/wallet")
	if nil != err {
		panic(err.Error())
	}
	fmt.Println(u.String())
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	fmt.Printf(newStr)

}
