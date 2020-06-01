package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	body, err := ioutil.ReadFile("captcha.png")
	if err != nil {
		panic(err)
	}

	rsp, _ := http.Post("http://localhost:8080/file", "multipart/form-data", bytes.NewReader(body))
	ret := make([]byte, 4)
	_, _ = rsp.Body.Read(ret)
	fmt.Printf("rsp: code=%d, body=%s\n", rsp.StatusCode, string(ret))
}
