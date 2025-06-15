package core

import (
	"encoding/json"
	"errors"
	"net"

	global_configs "github.com/Pois-Noir/Botzilla-Utils/global_configs"
	utils_message "github.com/Pois-Noir/Botzilla-Utils/message"
)

func router(message *utils_message.Message, addr string) ([]byte, error) {

	// format
	// TODO add checking of status code in the tcp code
	// to make sure there was no transmission error
	// status code 1 byte
	// operation code 1 byte

	operationCode := message.Header.OperationCode
	body := message.Payload

	// Routing
	switch operationCode {
	case global_configs.REGISTERCOMPONENTOPERATIONCODE:
		return RegisterComponent(body, addr) // retuns byte slice and error
	case global_configs.GETCOMPONENTOPERATIONCODE:
		return GetComponent(body)
	case global_configs.GETCOMPONENTSOPERATIONCODE:
		return GetComponents()
	}

	return nil, errors.New("Invalid Operation Code")

}

func RegisterComponent(body map[string]interface{}, addr string) ([]byte, error) {

	name := body["name"]
	port := body["port"]

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

func GetComponent(body map[string]interface{}) ([]byte, error) {

	name := string(body)

	registery := GetRegistery()
	comp := registery.Get(name)

	if comp == nil {
		return nil, errors.New("Component not found")
	}

	return []byte(comp.Address), nil
}
