package pkg

import (
    "github.com/golang/protobuf/proto"
    "google.golang.org/protobuf/encoding/protowire"
    "pbembedding/pkg/protodef"
)

func commonImplementation(messageBytes []byte, name string) (requestBytes []byte, err error) {
    // receive it from file or network, not important.
    var message protodef.HugeMessage
    _ = proto.Unmarshal(messageBytes, &message) // slow
    request := protodef.Request{
        Name:    name,
        Payload: &message,
    }
    return proto.Marshal(&request) // slow
}

func binaryEmbeddingImplementation(messageBytes []byte, name string) (requestBytes []byte, err error) {
    // 1. create a request with all ready except the payload. and marshal it.
    request := protodef.Request{
        Name: name,
    }
    requestBytes, err = proto.Marshal(&request)
    if err != nil {
        return nil, err
    }
    // 2. manually append the payload to the request, by protowire.
    requestBytes = protowire.AppendTag(requestBytes, 2, protowire.BytesType) //  embedded message is same as a bytes field, in wire view.
    requestBytes = protowire.AppendBytes(requestBytes, messageBytes)
    return requestBytes, nil
}
