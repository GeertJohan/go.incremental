package incremental

import (
	"sync"
)

type Int32 struct {
	increment int32
	lock      sync.Mutex
}

// Next returns the following 
func (i *Int32) Next() int32 {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

func (i *Int32) Last() int32 {
	return i.increment
}
