package core

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func StartTCPServer(port int, secretKey string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	address := listener.Addr().String()
	fmt.Println("Listening on:", address)

	if err != nil {
		fmt.Println("There was an error starting the server: \n", err)
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Starting Server")

	go cleanupInactiveListeners()

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: \n", err)
			continue
		}

		go handler(conn, secretKey)
	}
}

func handler(conn net.Conn, secretKey string) {

	defer conn.Close()

	componentAddr := conn.RemoteAddr().String()

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

	hash := [32]byte{}
	_, err = conn.Read(hash[:])

	if err != nil {
		fmt.Printf("Error reading from connection: %v\n", err)
		return
	}

	if !verifyHMAC(buffer, []byte(secretKey), hash[:]) {
		fmt.Printf("Error verifying HMAC signature\n")
		return
	}

	response, err := router(buffer, componentAddr)
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

func cleanupInactiveListeners() {

	for {
		registry := GetRegistery()
		components := registry.Data()
		for name, component := range components {

			go func(_name string, _component *Component) {
				_, err := net.Dial("tcp", _component.Address)
				if err != nil {
					registery.Remove(_name)
				}
			}(name, component)
		}

		time.Sleep(10 * time.Second)
	}
}

func generateHMAC(data []byte, key []byte) []byte {
	mac := hmac.New(sha256.New, key) // 32 bytes
	mac.Write(data)
	return mac.Sum(nil)
}

func verifyHMAC(data []byte, key []byte, hash []byte) bool {
	// Generate HMAC for the provided data using the same key
	generatedHMAC := generateHMAC(data, key)

	// Use subtle.ConstantTimeCompare to securely compare the two HMACs
	return subtle.ConstantTimeCompare(generatedHMAC, hash) == 1
}
