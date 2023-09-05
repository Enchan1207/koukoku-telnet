//
//
//

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// 接続
	connection, err := net.Dial("tcp", "koukoku.shadan.open.ad.jp:23")
	if err != nil {
		log.Fatal(err)
	}

	// 受信
	buffer := make([]byte, 1024)
	size, err := connection.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	for err == nil {
		if size > 0 {
			fmt.Print(string(buffer))
		}

		size, err = connection.Read(buffer)
	}
	if err != nil {
		log.Fatal(err)
	}

	connection.Close()
}
