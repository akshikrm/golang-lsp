package lsp

import (
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	request := Request{Method: "test"}
	_, err := Encode(request)
	if err != nil {
		t.Fatalf("Failed to encode %v due to %s", request, err)
	}
}

func TestDecode(t *testing.T) {
	request := "{\"method\":\"test\"}"
	reader := strings.NewReader(request)
	decoded := Request{}
	err := Decode(reader, &decoded)
	if err != nil {
		t.Fatalf("Failed to decode %s due to %s", request, err)
	}
}
