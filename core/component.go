package core

type Component struct {
	Name        string
	Group       []string
	Address     string
	SendToken   []byte
	ListenToken []byte
}

func newComponent(name string, address string) *Component {
	return &Component{
		Name:        name,
		Group:       nil,
		Address:     address,
		SendToken:   []byte{1, 1, 11, 2, 14, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		ListenToken: []byte{1, 2, 3, 4},
	}
}

func GenerateListenToken() []byte {
	return []byte{1, 2, 3, 4}
}
