package main

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/ingojaeckel/golang-demo-protobuf"
)

func TestMarshal(t *testing.T) {
	param := int32(2)
	r := &messages.SomeRequest{Param: &param}

	marshalled, err := proto.Marshal(r)
	if err != nil {
		t.Fatalf("failed to marshal: %s\n", err.Error())
	}
	fmt.Printf("marshalled: %v\n", marshalled)

	var unmarshalled messages.SomeRequest
	if err := proto.Unmarshal(marshalled, &unmarshalled); err != nil {
		t.Fatalf("failed to unmarshal: %s\n", err.Error())
	}

	fmt.Printf("successfully unmarshalled to: '%s'\n", unmarshalled.String())
}
