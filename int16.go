package incremental

import (
	"sync"
)

type Int16 struct {
	increment int16
	lock      sync.Mutex
}

// Next returns the following 
func (i *Int16) Next() int16 {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

func (i *Int16) Last() int16 {
	return i.increment
}
