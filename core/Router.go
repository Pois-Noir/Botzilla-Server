package core

import (
	"encoding/json"
	"errors"
	"net"
)

func router(message []byte, addr string) ([]byte, error) {

	// format
	// TODO add checking of status code in the tcp code
	// to make sure there was no transmission error
	// status code 1 byte
	// operation code 1 byte

	operationCode := message[0]
	body := message[1:]

	// Routing
	switch operationCode {
	case 0:
		return RegisterComponent(body, addr) // retuns byte slice and error
	case 2:
		return GetComponent(body)
	case 69:
		return GetComponents()
	}

	return nil, errors.New("Invalid Operation Code")

}

func RegisterComponent(body []byte, addr string) ([]byte, error) {

	// Decoding the message
	decodedBody := map[string]string{}
	err := json.Unmarshal(body, &decodedBody)
	if err != nil {
		return nil, err
	}
	name := decodedBody["name"]
	port := decodedBody["port"]

	// Generating listener addr of componenet
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	listenerAddress := net.JoinHostPort(host, port)

	// Adding Component to registery
	new_comp := newComponent(name, listenerAddress)
	registery := GetRegistery()
	err = registery.Add(name, new_comp) // if component exist returns error

	return []byte("registered"), err
}

func GetComponents() ([]byte, error) {

	var result []string

	registery := GetRegistery()
	registery.ForEach(func(s string, c *Component) {
		result = append(result, c.Name)
	})

	data, err := json.Marshal(result)
	return data, err
}

func GetComponent(body []byte) ([]byte, error) {

	name := string(body)

	registery := GetRegistery()
	comp := registery.Get(name)

	if comp == nil {
		return nil, errors.New("Component not found")
	}

	return []byte(comp.Address), nil
}
