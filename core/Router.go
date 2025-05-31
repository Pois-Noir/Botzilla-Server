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

// TODO
// Have to generate token
func RegisterComponent(body []byte, addr string) ([]byte, error) {

	decodedBody := map[string]string{}
	err := json.Unmarshal(body, &decodedBody)
	if err != nil {
		return nil, err
	}

	name := decodedBody["name"]
	port := decodedBody["port"]

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	listenerAddress := net.JoinHostPort(host, port)

	comp := newComponent(name, listenerAddress)

	componentList := GetComponentList()

	err = componentList.Add(name, comp)
	if err != nil {
		return nil, err
	}

	token := []byte("amir12345")

	fmt.Println(decodedBody)
	fmt.Println(addr)

	return token, err
}

// TODO
// Current Code is old and it looks wrong
// future amir look into it
func GetComponents() ([]byte, error) {
	componentList := GetComponentList()

	var result []string

	for _, value := range componentList.data {
		result = append(result, value.Name)
	}

	data, err := json.Marshal(result)

	return data, err
}

func GetComponent(body []byte) ([]byte, error) {

	name := string(body)

	componentList := GetComponentList()

	comp := componentList.Get(name)

	if comp == nil {
		return nil, errors.New("Component not found")
	}

	return []byte(comp.Address), nil
}
