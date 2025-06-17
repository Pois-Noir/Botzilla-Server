package core

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	global_configs "github.com/Pois-Noir/Botzilla-Utils/global_configs"
	header_pkg "github.com/Pois-Noir/Botzilla-Utils/header"
	utils_hmac "github.com/Pois-Noir/Botzilla-Utils/hmac"
)

func StartTCPServer(port int, secretKey string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	address := listener.Addr().String()
	fmt.Println("Listening on:", address)

	// TODO
	if err != nil {
		fmt.Println("There was an error starting the server: \n", err)
		os.Exit(1)
	}

	defer listener.Close()

	// TODO
	fmt.Println("Starting Server")

	for {

		conn, err := listener.Accept()
		// TODO
		if err != nil {
			fmt.Println("Error accepting connection: \n", err)
			continue
		}

		go handler(conn, secretKey)
	}
}

func handler(conn net.Conn, secretKey string) {

	defer conn.Close()

	// getting incoming request addr
	componentAddr := conn.RemoteAddr().String()
	// creating a buffered reader from the connection object
	bReader := bufio.NewReader(conn)

	var RequestHeaderBuffer [global_configs.HEADER_LENGTH]byte
	n, err := io.ReadFull(bReader, RequestHeaderBuffer[:])

	// TODO
	if uint32(n) < global_configs.HEADER_LENGTH {

	}

	// TODO
	if err != nil {
		// do something
		// // call Amir 438 282 3324
	}

	RequestHeader, err := header_pkg.Decode(RequestHeaderBuffer[:])

	// TODO
	if err != nil {
		// do something
		// // call Amir 438 282 3324
	}

	PayloadBuffer := make([]byte, RequestHeader.PayloadLength)

	// trying to read all of the message
	n, err = io.ReadFull(bReader, PayloadBuffer)

	// TODO
	if uint32(n) < RequestHeader.PayloadLength {
		// do something
		// // call Amir 438 282 3324
	}

	// TODO
	if err != nil {
		// we must send the sender appropriate message
		// to let them know what actually happened
	}

	hash := [global_configs.HASH_LENGTH]byte{}
	_, err = io.ReadFull(bReader, hash[:])

	// TODO
	if err != nil {
		fmt.Printf("Error reading from connection: %v\n", err)
		return
	}

	// TODO
	if !utils_hmac.VerifyHMAC(PayloadBuffer, []byte(secretKey), hash[:]) {
		fmt.Printf("Error verifying HMAC signature\n")
		return
	}
	// after we have successfully determined that the message is not tampered with
	// we can decode it

	ResponsePayload, err := router(
		PayloadBuffer,
		RequestHeader.OperationCode,
		componentAddr,
	)

	// TODO
	if err != nil {
		fmt.Println("There was an error processing your request: \n", err)
		return
	}

	// Todo
	// Add server status code response
	ResponseHeader := header_pkg.NewHeader(
		global_configs.OK_STATUS,
		global_configs.USER_MESSAGE_OPERATION_CODE,
		uint32(len(ResponsePayload)),
	)

	Response := append(ResponseHeader.Encode(), ResponsePayload...)

	conn.Write(Response)

}
