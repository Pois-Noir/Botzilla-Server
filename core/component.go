package core

type Component struct {
	Name    string
	Group   []string
	Address string
	token   string
}

func newComponent(name string, address string) *Component {
	return &Component{
		Name:    name,
		Group:   nil,
		Address: address,
		token:   "test123",
	}
}

func (c *Component) GetToken() []byte {
	return []byte{1, 1, 11, 2, 14, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

}
