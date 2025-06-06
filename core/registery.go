package core

import (
	"sync"

	"github.com/Pois-Noir/Ren"
)

var (
	registery *Ren.SafeList[string, Component]
	once      sync.Once
)

func GetRegistery() *Ren.SafeList[string, Component] {
	once.Do(func() {
		registery = Ren.NewSafeList[string, Component]()
	})
	return registery
}
