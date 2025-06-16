package core

import (
	"sync"

	"github.com/Pois-Noir/Botzilla-Utils/safemap"
)

var (
	registery *safemap.SafeMap[string, Component] // file name changed to map instead of list
	once      sync.Once
)

func GetRegistery() *safemap.SafeMap[string, Component] {
	once.Do(func() {
		registery = safemap.NewSafeMap[string, Component]()
	})
	return registery
}
