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
		w.WriteHeader(400)
		w.Header().Add("Content-Type", "application/octet-stream")

		v := err.Error()
		resp := messages.SomeResponse{Value: &v}
		respBytes, _ := proto.Marshal(&resp)
		w.Write(respBytes)
		return
	}
	fmt.Printf("Unmarshalled request to: %v\n", someRequest.GetParam())

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/octet-stream")
	v := "success"
	resp := messages.SomeResponse{Value: &v}
	respBytes, _ := proto.Marshal(&resp)
	fmt.Printf("Replying back bytes: %v\n", respBytes)
	w.Write(respBytes)
}
