//
// お手製ガバガバTCPクライアント
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

	// 接続先とタイムアウトを指定して接続
	host := "koukoku.shadan.open.ad.jp:23"
	timeoutSec := 5
	connection, err := net.DialTimeout("tcp", host, time.Duration(timeoutSec)*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// 受信バッファを構成
	buffer := make([]byte, 32)

	// 受信開始
	size, err := connection.Read(buffer)
	for size != 0 && err == nil {
		fmt.Print(convertSjisToUtf8(buffer[:size]))
		size, err = connection.Read(buffer)
	}

	// エラー終了?
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
