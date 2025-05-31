package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func router(message []byte, addr string) ([]byte, error) {

	operationCode := message[0]
	body := message[1:]

	// Register Component
	if operationCode == 0 {
		return RegisterComponent(body, addr)
	}

	if operationCode == 69 {
		return GetComponents()
	} else if operationCode == 2 {
		return GetComponent(body)
	}

	return nil, nil

}

func RegisterComponent(body []byte, addr string) ([]byte, error) {

	decodedBody := map[string]string{}
	err := json.Unmarshal(body, &decodedBody)
	if err != nil {
		return nil, err
	}

	registery := GetRegistery()

	name := decodedBody["name"]
	port := decodedBody["port"]

	comp := registery.GetComponent(name)
	if comp != nil {
		return nil, errors.New("Component already exists")
	}

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	listenerAddress := net.JoinHostPort(host, port)

	fmt.Println(decodedBody)
	fmt.Println(addr)
	token, err := registery.AddComponent(name, listenerAddress)

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

	if comp == nil {
		return nil, errors.New("Component not found")
	}

	return []byte(comp.Address), nil
}
