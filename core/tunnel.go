package core

type Tunnel struct {
	Name    string
	Address string
}

func newTunnel(name string, address string) *Tunnel {

	return &Tunnel{
		Name:    name,
		Address: address,
	}
}
