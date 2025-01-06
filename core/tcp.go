package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func StartTCPServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("There was an error starting the server: \n", err)
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Commanding server has started on port 8080")

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: \n", err)
			continue
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {

	defer conn.Close()

	componentAddr := conn.RemoteAddr().String()

	var token [16]byte
	conn.Read(token[:])

	var requestHeader [4]byte
	conn.Read(requestHeader[:])

	// Convert Response Header to int32
	requestSize := int32(requestHeader[0]) |
		int32(requestHeader[1])<<8 |
		int32(requestHeader[2])<<16 |
		int32(requestHeader[3])<<24

	buffer := make([]byte, requestSize)

	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading from connection: %v\n", err)
		return
	}

	response, err := router(buffer, token, componentAddr)
	if err != nil {
		fmt.Println("There was an error processing your request: \n", err)
		return
	}

	// Generate Header for server
	messageLength := int32(len(response))
	responseHeaderBuf := new(bytes.Buffer)
	err = binary.Write(responseHeaderBuf, binary.LittleEndian, messageLength) // LittleEndian like umar
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return
	}

	conn.Write(responseHeaderBuf.Bytes())
	conn.Write(response)

}

func callComponent(addr string) {

}
