package rpc

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type Request struct {
	Method string `json:"method"`
}

func EncodeMessage(payload any) (*bytes.Buffer, error) {
	writer := bytes.NewBuffer([]byte(""))
	err := json.NewEncoder(writer).Encode(&payload)
	if err != nil {
		return nil, err
	}
	return writer, nil
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(data []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, nil
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	baseMessage := BaseMessage{}
	if err := json.NewDecoder(bytes.NewReader(content[:contentLength])).Decode(&baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil
}
