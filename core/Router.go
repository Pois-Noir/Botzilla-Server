package core

import "encoding/json"

func router(message []byte, token [16]byte, addr string) ([]byte, error) {

	operationCode := message[0]
	content := string(message[1:])

	// Register Component
	if operationCode == 0 {
		return RegisterComponent(content, addr)
	}

	if !checkToken(addr, token) {
		return nil, nil
	}

	if operationCode == 69 {
		return GetComponents()
	} else if operationCode == 2 {
		return GetComponent(content)
	}

	return nil, nil

}

func checkToken(addr string, token [16]byte) bool {
	return true
}

func RegisterComponent(name string, addr string) ([]byte, error) {
	registery := GetRegistery()

	token, err := registery.AddComponent(name, addr)

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

func GetComponent(name string) ([]byte, error) {
	registery := GetRegistery()

	comp := registery.GetComponent(name)

	// TODO, Define an error
	if comp == nil {
		return nil, nil
	}

	return []byte(comp.Address), nil
}
