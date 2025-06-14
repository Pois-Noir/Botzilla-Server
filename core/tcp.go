package core

import (
	"bytes"
	"io"

	// "crypto/hmac"
	// "crypto/sha256"
	// "crypto/subtle"
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"

	global_configs "github.com/Pois-Noir/Botzilla-Utils/global_configs"
	utils_header "github.com/Pois-Noir/Botzilla-Utils/header"
	utils_hmac "github.com/Pois-Noir/Botzilla-Utils/hmac"
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

	// create a buffered reader
	// get the header
	// get the body bytes
	// get the hash bytes
	// check if the hash is okay or not

	// creating a buffered reader from the connection object
	bReader := bufio.NewReader(conn)

	// need to get the header
	header, err := utils_header.DecodeHeaderBuffered(bReader)

	// check if there were any errors while reading the header
	if err != nil {
		// do something
		// // call Amir 438 282 3324
	}
	// i dont know what this line does
	// call Amir 438 282 3324
	componentAddr := conn.RemoteAddr().String()

	messageLength := header.Length
	buffer := make([]byte, messageLength)

	// trying to read all of the message
	n, err := io.ReadFull(bReader, buffer[:])

	// if the bytes read
	// is less than the message Length
	// we know that there were transmission errors
	if n < int(messageLength) {
		// do something
		// // call Amir 438 282 3324
	}
	// if there were errors trying to read the error
	if err != nil {
		// we must send the sender appropriate message
		// to let them know what actually happened
	}

	hash := [global_configs.HASHLENGTH]byte{}
	_, err = conn.Read(hash[:])

	if err != nil {
		fmt.Printf("Error reading from connection: %v\n", err)
		return
	}

	if !utils_hmac.VerifyHMAC(buffer, []byte(secretKey), hash[:]) {
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

// func generateHMAC(data []byte, key []byte) []byte {
// 	mac := hmac.New(sha256.New, key) // 32 bytes
// 	mac.Write(data)
// 	return mac.Sum(nil)
// }

// func verifyHMAC(data []byte, key []byte, hash []byte) bool {
// 	// Generate HMAC for the provided data using the same key
// 	generatedHMAC := generateHMAC(data, key)

// 	// Use subtle.ConstantTimeCompare to securely compare the two HMACs
// 	return subtle.ConstantTimeCompare(generatedHMAC, hash) == 1
// }
