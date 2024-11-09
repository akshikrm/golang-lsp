package main

import (
	"bufio"
	"bytes"
	"educationlsp/rpc"
	"log"
	"os"
	"strconv"
)

func main() {
	logger := getLogger("./test.log")
	logger.Println("Starting")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(Split)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			logger.Printf("Scan failed due to error %s", err)
		}
		handleMessage(
			logger,
			scanner.Text(),
		)
	}
}

func Split(data []byte, EOF bool) (int, []byte, error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}

	//Content-Length: <number>
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}
	if len(content) < contentLength {
		return 0, nil, nil
	}
	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"methods"`
	// Params...
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`
	// Result...
	// Error...
}

func handleMessage(logger *log.Logger, msg string) {
	method, content, err := rpc.DecodeMessage([]byte(msg))
	if err != nil {
		logger.Printf("failed to decode, %s", err)
	}
	logger.Println("method", method)
	logger.Println("content", string(content))
}

func getLogger(file string) *log.Logger {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic("Log file not found")
	}
	return log.New(f, "[educationlsp] ", log.Ldate|log.Ltime)
}
