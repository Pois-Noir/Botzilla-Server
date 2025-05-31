package core

import "sync"

var (
	tunnelList    *SafeList[Tunnel]
	componentList *SafeList[Component]
	once          sync.Once
)

func initialize() {
	tunnelList = &SafeList[Tunnel]{
		data: make(map[string]*Tunnel),
	}
	componentList = &SafeList[Component]{
		data: make(map[string]*Component),
	}
}

func GetComponentList() *SafeList[Component] {
	once.Do(initialize)

	return componentList
}

func GetTunnelList() *SafeList[Tunnel] {
	once.Do(initialize)

	return tunnelList
}
