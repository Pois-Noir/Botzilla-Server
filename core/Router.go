package core

import (
	"encoding/json"
	"errors"
	"net"

	global_configs "github.com/Pois-Noir/Botzilla-Utils/global_configs"
)

// TODO
// Might get rid of the func since it is not doing shit
func router(encodedPayload []byte, operationCode uint8, addr string) ([]byte, error) {

	// format
	// TODO add checking of status code in the tcp code
	// to make sure there was no transmission error
	// status code 1 byte
	// operation code 1 byte

	// Routing
	switch operationCode {
	case global_configs.REGISTER_COMPONENT_OPERATION_CODE:
		return RegisterComponent(encodedPayload, addr) // retuns byte slice and error
	case global_configs.GET_COMPONENT_OPERATION_CODE:
		return GetComponent(encodedPayload)
	case global_configs.GET_COMPONENTS_OPERATION_CODE:
		return GetComponents()
	}

	return nil, errors.New("Invalid Operation Code")

}

func RegisterComponent(encodedPayload []byte, addr string) ([]byte, error) {

	var payload map[string]string

	json.Unmarshal(encodedPayload, &payload)

	name := payload["name"]
	port := payload["port"]

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

func GetComponent(encodedPayload []byte) ([]byte, error) {

	name := string(encodedPayload)

	registery := GetRegistery()
	comp := registery.Get(name)

	if comp == nil {
		return nil, errors.New("Component not found")
	}

	return []byte(comp.Address), nil
}
