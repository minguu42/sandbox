package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// ファイル以外のフォームフィールドの値はWriteFieldメソッドを使って登録する
	w.WriteField("name", "John Smith")

	fileWriter, err := w.CreateFormFile("thumbnail", "example.txt")
	if err != nil {
		log.Fatal(err)
	}
	readFile, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	w.Close()

	resp, err := http.Post("http://localhost:18888", w.FormDataContentType(), &b)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.Status)
}
