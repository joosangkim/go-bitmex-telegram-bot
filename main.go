// functions
// 1. Init ws
// 2. Get market value
// 3. Get user status
// TBD. orderer

package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	url := "https://testnet.bitmex.com/api/v1"
	resp, err := http.Get(url)
	if nil != err {
		panic(err.Error())
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	fmt.Printf(newStr)
}

func getQuote() {

}
