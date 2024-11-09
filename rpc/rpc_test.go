package rpc

import (
	"testing"
)

func TestEncode(t *testing.T) {
	request := Request{Method: "test"}
	_, err := EncodeMessage(request)
	if err != nil {
		t.Fatalf("Failed to encode %v due to %s", request, err)
	}
}

func TestDecode(t *testing.T) {
	request := "Content-Length: 23\r\n\r\n{\"method\":\"initialize\"}"
	method, content, err := DecodeMessage([]byte(request))
	if err != nil {
		t.Fatalf("Failed to decode %s due to %s", request, err)
	}

	contentLength := len(content)
	if method != "initialize" {
		t.Fatalf("Expected method is initialize but got %s", method)
	}

	if contentLength != 23 {
		t.Fatalf("Expected length is 23 but got %d", contentLength)
	}

}
