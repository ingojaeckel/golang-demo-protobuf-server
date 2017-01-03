package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ingojaeckel/golang-demo-protobuf"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handlePost(w, r)
		} else {
			w.WriteHeader(415)
		}
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("running")
	http.ListenAndServe(":8080", nil)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)
	var someRequest messages.SomeRequest

	if err := proto.Unmarshal(bytes, &someRequest); err != nil {
		v := err.Error()
		sendBytes(w, 400, messages.SomeResponse{Value: &v})
		return
	}
	fmt.Printf("Request: %v\n", someRequest.GetParam())

	v := "success"
	sendBytes(w, 200, messages.SomeResponse{Value: &v})
}

func sendBytes(w http.ResponseWriter, status int, resp messages.SomeResponse) {
	respBytes, _ := proto.Marshal(&resp)
	fmt.Printf("Response: %v\n", respBytes)

	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Write(respBytes)
}
