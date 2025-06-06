package core

import (
	"sync"

	"github.com/Pois-Noir/Ren"
)

var (
	registery *Ren.SafeMap[string, Component] // file name changed to map instead of list
	once      sync.Once
)

func GetRegistery() *Ren.SafeMap[string, Component] {
	once.Do(func() {
		registery = Ren.NewSafeMap[string, Component]()
	})
	return registery
}
