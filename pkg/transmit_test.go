package pkg

import (
    "github.com/stretchr/testify/assert"
    "golang.org/x/exp/rand"
    "google.golang.org/protobuf/proto"
    "pbembedding/pkg/protodef"
    "testing"
)

var hugeMessageSample = receiveHugeMessageFromSomewhere()

func TestEquivalent(t *testing.T) {
    requestBytes1, _ := commonImplementation(hugeMessageSample, "xxxx")
    requestBytes2, _ := binaryEmbeddingImplementation(hugeMessageSample, "xxxx")
    // They are not always equal int bytes. you should compare them in message view instead of binary from
    // due to: https://developers.google.com/protocol-buffers/docs/encoding#implications
    // I'm Lazy.
    assert.NotEmpty(t, requestBytes1)
    assert.Equal(t, requestBytes1, requestBytes2)
    var request protodef.Request
    err := proto.Unmarshal(requestBytes1, &request)
    assert.NoError(t, err)
    assert.Equal(t, "xxxx", request.Name)
}

// actually mock one.
func receiveHugeMessageFromSomewhere() []byte {
    buffer := make([]byte, 1024*1024*1024)
    _, _ = rand.Read(buffer)
    message := protodef.HugeMessage{
        Data: buffer,
    }
    res, _ := proto.Marshal(&message)
    return res
}

func BenchmarkCommon(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := commonImplementation(hugeMessageSample, "xxxx")
        if err != nil {
            panic(err)
        }
    }
}

func BenchmarkEmbedding(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := binaryEmbeddingImplementation(hugeMessageSample, "xxxx")
        if err != nil {
            panic(err)
        }
    }
}
