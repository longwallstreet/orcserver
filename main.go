package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/otiai10/gosseract"
)

func handleFile(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	client := gosseract.NewClient()
	defer client.Close()

	client.SetImageFromBytes(body)
	text, err := client.Text()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	ret := strings.Replace(text, " ", "", -1)
	fmt.Printf("recognize=%s\n", ret)
	w.Write([]byte(ret))
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "http listen port")

	flag.Parse()
	fmt.Printf("http listening %d\n", port)

	http.HandleFunc("/file", handleFile)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
