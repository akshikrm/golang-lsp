package lsp

import (
	"bytes"
	"encoding/json"
	"io"
)

type Request struct {
	Method string `json:"method"`
}

func Encode(payload any) (*bytes.Buffer, error) {
	writer := bytes.NewBuffer([]byte(""))
	err := json.NewEncoder(writer).Encode(&payload)
	if err != nil {
		return nil, err
	}
	return writer, nil
}

func Decode(reader io.Reader, decoded *Request) error {
	newDecoder := json.NewDecoder(reader)
	for {
		err := newDecoder.Decode(&decoded)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}
