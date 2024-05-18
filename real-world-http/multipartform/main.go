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
	w.WriteField("name", "John Smith") // ファイル以外のフォームフィールドの値はWriteFieldメソッドを使って登録する
	fileWriter, err := w.CreateFormFile("thumbnail", ".gitignore")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open(".gitignore")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	w.Close()

	resp, err := http.Post("http://localhost:18888", w.FormDataContentType(), &b)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println(resp.Status)
}
