package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/ingojaeckel/golang-demo-protobuf"
)

func main() {
	param := int32(2)
	r := &messages.SomeRequest{Param: &param}

	marshalled, err := proto.Marshal(r)
	if err != nil {
		fmt.Printf("failed to marshal: %s\n", err.Error())
		return
	}
	fmt.Printf("marshalled: %v\n", marshalled)

	var unmarshalled messages.SomeRequest
	if err := proto.Unmarshal(marshalled, &unmarshalled); err != nil {
		fmt.Printf("failed to unmarshal: %s\n", err.Error())
		return
	}
	fmt.Printf("successfully unmarshalled to: '%s'\n", unmarshalled.String())
}
