package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Handling connection from", conn.RemoteAddr())
	if err := handleConnection(conn); err != nil {
		fmt.Println("Error handling connection: ", err.Error())
		os.Exit(1)
	}
}

func handleConnection(conn net.Conn) error {
	// var buffer [4096]byte // arbitary size

	// n, err := conn.Read(buffer[:])
	// fmt.Printf("Read %d bytes, message: %s\n", n, string(buffer[:n]))
	// if err != nil {
	// 	return err
	// }

	req, err := readClientRequest(conn)
	if err != nil {
		return err
	}
	fmt.Printf("Read client request, size %d, correlation id %d", req.MessageSize, req.Header.CorrelationId())

	resp := handleClientRequest(req)

	n, err := conn.Write(resp.Encode())
	fmt.Printf("Wrote %d bytes, message: %x\n", n, resp)

	return err
}

func readClientRequest(conn net.Conn) (*Request, error) {
	var messageSizeBuf [4]byte
	n, err := conn.Read(messageSizeBuf[:])
	if err != nil {
		return nil, err
	}
	if n != 4 {
		return nil, fmt.Errorf("expected 4 bytes, got %d", n)
	}

	// Get size of the request from first 4 bytes
	messageSize := int32(binary.BigEndian.Uint32(messageSizeBuf[:]))
	if messageSize < 0 {
		return nil, fmt.Errorf("invalid message size: %d", messageSize)
	}

	requestBytes := make([]byte, messageSize)
	n, err = conn.Read(requestBytes)
	if err != nil {
		return nil, err
	}
	if n != int(messageSize) {
		return nil, fmt.Errorf("expected %d bytes, got %d", messageSize, n)
	}

	return NewRequest(messageSize, requestBytes), nil
}

func handleClientRequest(req *Request) *Response {
	return NewResponse(0, req.Header.CorrelationId())
}
