package main

import (
	"fmt"
	"sync"
	"time"
)

//dictionary inherits the mutex lock
type dictionary struct {
	sync.Mutex
	rel map[string]string
}

//add a new key value pare if the given key is not present in the map
//returns an error if the given key is already present
func (d *dictionary) putIfAbsent(key string, value string) error {
	//Locks the object, add the key-value pare to the map, and then unlock the object
	d.Lock()
	defer d.Unlock()
	if _, isPresent := d.rel[key]; isPresent {
		return fmt.Errorf("key [%s] is already present in the dictionary", key)
	}
	d.rel[key] = value
	return nil
}

func newDictionary() dictionary {
	return dictionary{
		rel: make(map[string]string),
	}
}

func main() {
	d := newDictionary()
	keys := []string{"foo", "bar", "baz", "bar"}
	values := []string{"value1", "value2", "value3", "value4"}
	for i := 0; i < len(keys); i++ {
		// Starts a new goroutine
		// that that add to the dictionary as a key string representation of i % 7
		// and string representation of i as a value
		go func(index int) {
			err := d.putIfAbsent(keys[index], values[index])
			if err != nil {
				//print the error
				fmt.Println(err.Error())
			}
		}(i)
	}
	time.Sleep(time.Second * 2)
	//print the result
	fmt.Println(d.rel)
}
