package singleton

import (
	"sync"
)

type Object struct {
}

var once sync.Once
var instance *Object

func GetInstance() *Object {
	// Creates a singleton instance once.
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
