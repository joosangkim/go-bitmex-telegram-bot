// functions
// 1. Init socket
// 2. Get market value
// 3. Get user status
// TBD. orderer

package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func main() {
	url := "wss://www.bitmex.com/realtime"
	var dialer = websocket.DefaultDialer
	c, res, err := dialer.Dial(url, nil)
	if nil != err {
		panic(err)
	}
	defer c.Close()

	fmt.Println(res.Body)
	err = c.WriteJSON("{\"op\": \"announcement\", \"args\": []}")
	if nil != err {
		panic(err)
	}
	t, msg, err := c.ReadMessage()
	if nil != err {
		panic(err)
	}
	fmt.Println(t)
	fmt.Println(string(msg))

}
