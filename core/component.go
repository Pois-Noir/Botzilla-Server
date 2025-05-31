package core

type Component struct {
	Name    string
	Group   []string
	Address string
	Streams []string
}

func newComponent(name string, address string) *Component {
	return &Component{
		Name:    name,
		Group:   nil,
		Address: address,
	}
}
