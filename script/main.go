package main

import (
    "fmt"
    "github.com/golang/protobuf/proto"
    "pbembedding/pkg/protodef"
)

func main() {
    message := protodef.HugeMessage{
        Data: []byte{0xAA},
    }
    request := protodef.Request{
        Payload: &message,
    }
    cs, _ := proto.Marshal(&message)
    fmt.Printf("message: %X\n", cs)
    bs, _ := proto.Marshal(&request)
    fmt.Printf("request: %X\n", bs)
}
