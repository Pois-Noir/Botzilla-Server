package core

import (
	"encoding/json"
	"errors"
	"net"
)

func router(message []byte, addr string) ([]byte, error) {

	operationCode := message[0]
	body := message[1:]

	switch operationCode {
	case 0:
		return RegisterComponent(body, addr)
	case 2:
		return GetComponent(body)
	case 69:
		return GetComponents()
	}

	return nil, errors.New("invalid operation code")

}

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

	response := []byte("component registered")

	return response, err
}

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
