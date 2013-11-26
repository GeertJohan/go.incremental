package incremental

import (
	"sync"
)

type Int64 struct {
	increment int64
	lock      sync.Mutex
}

// Next returns the following 
func (i *Int64) Next() int64 {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

func (i *Int64) Last() int64 {
	return i.increment
}
