package core

import (
	"encoding/json"
	"fmt"
	"net"
)

func router(message []byte, token [16]byte, addr string) ([]byte, error) {

	operationCode := message[0]
	body := message[1:]

	// Register Component
	if operationCode == 0 {
		return RegisterComponent(body, addr)
	}

	if !checkToken(addr, token) {
		return nil, nil
	}

	if operationCode == 69 {
		return GetComponents()
	} else if operationCode == 2 {
		return GetComponent(body)
	}

	return nil, nil

}

func checkToken(addr string, token [16]byte) bool {
	return true
}

func RegisterComponent(body []byte, addr string) ([]byte, error) {

	decodedBody := map[string]string{}
	err := json.Unmarshal(body, &decodedBody)
	if err != nil {
		return nil, err
	}

	registery := GetRegistery()

	port := decodedBody["port"]

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	listenerAddress := net.JoinHostPort(host, port)

	fmt.Println(decodedBody)
	fmt.Println(addr)
	token, err := registery.AddComponent(decodedBody["name"], listenerAddress)

	return token, err
}

func GetComponents() ([]byte, error) {
	registery := GetRegistery()

	var result []string

	for _, value := range registery.components {
		result = append(result, value.Name)
	}

	data, err := json.Marshal(result)

	return data, err
}

func GetComponent(body []byte) ([]byte, error) {

	name := string(body)

	registery := GetRegistery()

	comp := registery.GetComponent(name)

	// TODO, Define an error
	if comp == nil {
		return nil, nil
	}

	return []byte(comp.Address), nil
}
