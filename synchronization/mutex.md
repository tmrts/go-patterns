# Mutex/Lock pattern
A mutex is a synchronization pattern that allows exclusive access to shared data by locking and unlocking.

## Implementation

```go
package mutex

import (
	"errors"
	"sync"
	"fmt"
)

//dictionary inherits the mutex lock
type Dictionary struct {
	sync.Mutex
	rel map[string]string
}

func New() Dictionary {
	return Dictionary{
		rel: make(map[string]string),
	}
}

//add a new key value pare if the given key is not present in the map
//returns an error if the given key is already present
func (d *Dictionary) PutIfAbsent(key string, value string) error {
	d.Lock() //lock the object
	defer d.Unlock() //unlock the object when the work is done
	//return the error if key is already present in the map
	if _, isPresent := d.rel[key]; isPresent {
		return fmt.Errorf("key [%s] is already present in the dictionary", key)
	}
	d.rel[key] = value
	return nil
}
``` 

## Usage

```go
func main() {
    d := mutext.New()
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
```