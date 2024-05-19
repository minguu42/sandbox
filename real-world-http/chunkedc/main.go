package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// TCPソケットオープン
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	conn, err := dialer.Dial("tcp", "localhost:18888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// リクエスト送信
	request, err := http.NewRequest("GET", "http://localhost:18888/chunked", nil)
	if err != nil {
		log.Fatal(err)
	}
	if err := request.Write(conn); err != nil {
		log.Fatal(err)
	}

	// 読み込み
	reader := bufio.NewReader(conn)
	// フィールドを読み込む
	resp, err := http.ReadResponse(reader, request)
	if err != nil {
		log.Fatal(err)
	}
	if resp.TransferEncoding[0] != "chunked" {
		log.Fatal("wrong transfer encoding")
	}
	for {
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(sizeStr[0]))
		// 16進数のサイズをパース。サイズがゼロならクローズ
		// -2にしているのは改行文字が\r\nで2バイトあるから
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		if size == 0 {
			break
		}
		// サイズ数分バッファを確保して読み込み
		line := make([]byte, int(size))
		reader.Read(line)
		reader.Discard(2)
		fmt.Printf("  %s\n", strings.TrimSpace(string(line)))
	}
}
