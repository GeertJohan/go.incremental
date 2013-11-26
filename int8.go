package incremental

import (
	"sync"
)

type Int8 struct {
	increment int8
	lock      sync.Mutex
}

// Next returns the following 
func (i *Int8) Next() int8 {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

func (i *Int8) Last() int8 {
	return i.increment
}
