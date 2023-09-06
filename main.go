//
//
//

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {

	// 接続
	host := "koukoku.shadan.open.ad.jp:23"
	timeoutSec := 5
	connection, err := net.DialTimeout("tcp", host, time.Duration(timeoutSec)*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// 受信開始
	buffer := make([]byte, 32)
	size, err := connection.Read(buffer)
	for size != 0 && err == nil {
		if size > 0 {
			fmt.Print(convertSjisToUtf8(buffer[:size]))
		}
		size, err = connection.Read(buffer)
	}
	if err != nil {
		log.Fatal(err)
	}

	connection.Close()
}

func convertSjisToUtf8(data []byte) string {
	transformer := transform.NewReader(bytes.NewReader(data), japanese.ShiftJIS.NewDecoder())
	u8str, err := ioutil.ReadAll(transformer)
	if err != nil {
		log.Fatal(err)
	}
	return string(u8str)
}
